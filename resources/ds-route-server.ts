import * as path from "https://deno.land/std/path/mod.ts";
import { Server } from "https://deno.land/std/http/server.ts";
import type {ServerRequest} from "https://deno.land/std/http/server.ts";

import Metadata from "./ds-metadata.ts";
import DsServices from './ds-services.ts';

const serverService = 11;
const serverReadyCommand = 11;

export class DsRouteServer {
	private sock_file:string;
	private server :Server|undefined;

	private mod_cache :Map<string,any> = new Map;

	constructor() {
		this.sock_file = path.join(Metadata.sock_path, 'server.sock');
	}

	async startServer() {
		const listener = await Deno.listen({ path: this.sock_file, transport: "unix" });
		this.server = new Server(listener);

		this.listen();	// does this mean errors are uncaught?

		const twine = DsServices.getTwine();
		const reply = await twine.sendBlock(serverService, serverReadyCommand, undefined)
		if(!reply.ok) {
			throw reply.error;
		}
	}
	private async listen() {
		if( this.server === undefined ) return;
		for await (const request of this.server) {
			this.handleRequest(request);
		}
	}

	async handleRequest(request :ServerRequest) {
		const t0 = performance.now();

		const headers = request.headers;
		const mod_file = headers.get('appspace-module');

		const fn = headers.get('appspace-function');

		if( mod_file === null ) {
			this.replyError(request, "appspace-module header is null");
			return;
		}

		let mod : any;
		try {
			mod = await this.loadModule(mod_file);
		}
		catch(e) {
			this.replyError(request, "Failed to import module "+mod_file+" Error: "+e);
			return;
		}

		let fnc = mod;
		if( fn ) {
			fnc = mod[fn];
		}

		try {
			fnc(request);
		}
		catch(e) {
			this.replyError(request, e);
			return;
		}

		const t1 = performance.now();
		console.log(`request took ${t1 - t0} milliseconds.`);
	}

	async loadModule(mod_file:string) :Promise<any> {
		if( this.mod_cache.has(mod_file) ) return this.mod_cache.get(mod_file);

		const mod = await import(mod_file);
		this.mod_cache.set(mod_file, mod);

		return mod;
	}
	
	async replyError(req :ServerRequest, message :string) {
		console.error(message);
		req.respond({status: 500, body: message})
	}

}

const sym = Symbol.for("DropServer DsRouteServer class singleton");
const w = <{[sym]?:DsRouteServer}>window;
if(w[sym] === undefined) w[sym] = new DsRouteServer;

export default w[sym] as DsRouteServer;