import * as path from "https://deno.land/std@0.106.0/path/mod.ts";
import libSupportIface from 'https://deno.land/x/dropserver_lib_support@v0.1.0/mod.ts';

import Metadata from './metadata.ts';
import DsServices from './ds-services.ts';
import Migrations from './migrations.ts';
import MigrationService from './ds-migrate-service.ts';
import AppRoutes from './app-router.ts';
import DsAppService from './ds-app-service.ts';
import DsRouteServer from './ds-route-server.ts';
import LibSupport from './libsupport.ts';

const metadata = new Metadata;
const services = new DsServices;

const w = <{["DROPSERVER"]?:libSupportIface}>window;
const libSupport = new LibSupport(metadata, services);
w["DROPSERVER"] = libSupport;

const migrations = new Migrations;
libSupport.setMigrations(migrations);

const migrationService = new MigrationService(migrations);
services.setMigrationService(migrationService);

const appRoutes = new AppRoutes;
libSupport.setAppRoutes(appRoutes);

const appService = new DsAppService(appRoutes);
services.setAppService(appService);

const server = new DsRouteServer(services, libSupport.appRoutes);
services.setServer(server);

const appMod = path.join(metadata.app_path, 'app.ts');
console.log("app module:", appMod);
await import(appMod);

appRoutes.loadRoutes();

services.initTwine(metadata.rev_sock_path);	// this results in host getting "Ready"
server.startServer(metadata.sock_path);	// only start it if we need to, but deal with later.
