import * as path from "https://deno.land/std@0.159.0/path/mod.ts";

import {Context, RouteType} from 'https://deno.land/x/dropserver_lib_support@v0.2.0/mod.ts';

import DsServices from './services.ts';
import type AppRoutes from '../approutes.ts';

// Some of this is heavily inspired by Deno std lib's server.ts:
// https://github.com/denoland/deno_std/blob/main/http/server.ts

export default class DsRouteServer {
	private listener :Deno.Listener|undefined;

	private httpConnections: Set<Deno.HttpConn> = new Set();

	private stop_resolve :undefined | ((value?: unknown) => void);

	constructor(private services:DsServices, private appRoutes: AppRoutes) {}

	async startServer(sockPath: string) {
		const sockFile = path.join(sockPath, 'server.sock');

		this.listener = await Deno.listen({ path: sockFile, transport: "unix" });

		this.listen();

		this.services.serverReady();

	}
	private async listen() {
		if( this.listener === undefined ) throw new Error("no listener");
		for await (const conn of this.listener) {
			this.serveHttp(conn);
		}
		if( this.stop_resolve != undefined ) {
			this.stop_resolve();
		}
	}
	stopServer() {
		return new Promise((resolve, reject) => {
			if( this.stop_resolve != undefined ) {
				reject("stop already called");
				return;
			}
			this.listener?.close();
			for (const httpConn of this.httpConnections) {
				this.closeHttpConn(httpConn);
			}
			this.httpConnections.clear();
			this.stop_resolve = resolve;
		});
	}

	async serveHttp(conn: Deno.Conn) {
		const httpConn = Deno.serveHttp(conn);
		this.httpConnections.add(httpConn);
		for await (const requestEvent of httpConn ) {
			this.handleRequest(requestEvent);
		}
		this.closeHttpConn(httpConn);
	}
	closeHttpConn(httpConn:Deno.HttpConn) {
		this.httpConnections.delete(httpConn);
		try {
			httpConn.close();
		} catch{
			// already closed
		}
	}

	async handleRequest(reqEvent: Deno.RequestEvent) {
		const t0 = performance.now();

		if( !this.appRoutes ) {	// should no longer happen
			this.replyError(reqEvent, "app ruouter not loaded");
			return;
		}

		const headers = reqEvent.request.headers;
		const matched_route = headers.get("X-Dropserver-Route-ID");
		if( matched_route === null ) {
			this.replyError(reqEvent, "X-Dropserver-Route-ID header is null");	
			// How would this happen? No sure, but it's not an app error, 
			// it's DS that didn't registerroutes correctly, or some other snafu.
			return;
		}
		const route = this.appRoutes.getRouteWithMatch(matched_route);
		if( route === undefined) {
			this.replyError(reqEvent, "route id not found");
			// similar to above
			return;
		}
		if( route.type !== RouteType.function ) {
			this.replyError(reqEvent, "matched route is not of type function");
			// same again
			return;
		}
		if( !route.handler ) {
			this.replyError(reqEvent, "no handler attached to route");
			// same.
			return;
		}
		if( route.match === undefined ) {
			// same again
			throw new Error("route returned without match function");
		}

		const req_url_str = headers.get("X-Dropserver-Request-URL");
		if( !req_url_str ) {
			this.replyError(reqEvent, "no request url found in headers");
			// That's a DS screw up
			return;
		}
		const req_url = new URL(req_url_str, "https://appspace/");
		const route_match = route.match(req_url.pathname);
		if( !route_match ) {
			this.replyError(reqEvent, "route failed to match in sandbox");
			// Weird incompatibility between ds host side and sandbox side route matching
			return;
		}

		const proxyId = headers.get("X-Dropserver-User-ProxyID");

		const ctx :Context = {
			req: reqEvent,
			url: req_url,	// request.url is readonly, so we can't set it, so we pass url in context. We could wrap request in Proxy and intercept get url
			params: <Record<string, unknown>>route_match.params,
			proxyId: proxyId
		};

		try {
			await route.handler(ctx);
		}
		catch(e) {
			this.replyError(reqEvent, e);
			// this error comes from app code, but it may also be from dropserver_app lib
			// But that's technically app code. Maybe dropserver_lib_suport can define an error
			// That then gets used to indicate a problem at the library level?
			return;
		}

		const t1 = performance.now();
		console.log(`request took ${t1 - t0} milliseconds.`);
	}

	replyError(reqEvent :Deno.RequestEvent, message :string) {
		console.error(message);
		reqEvent.respondWith( new Response(message,{status: 500}));
	}
}
