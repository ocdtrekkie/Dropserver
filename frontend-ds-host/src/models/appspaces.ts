import { reactive } from 'vue';

import {get, post} from '../controllers/userapi';

import { AppVersion } from './app_versions';

// these are owner's appspaces, not remotes.

// hierarchical data:
// - appspaces (listing)
// - appspace (node of above list)
// - appspace status (derived live data on readiness of appspace)
// - appspace upgrade available 

// relations:
// - appspace -> appversion
// - appspace =>* contacts

// For relations (contacts here), we need:
// - related contact ids
// - data for these ids.
//   ..here should be a preview

// from Go:
// type AppspaceMeta struct {
// 	AppspaceID int            `json:"appspace_id"`
// 	AppID      int            `json:"app_id"`
// 	AppVersion domain.Version `json:"app_version"`
// 	Subdomain  string         `json:"subdomain"`
// 	Created    time.Time      `json:"created_dt"`
// 	Paused     bool           `json:"paused"`
// }

export class Appspace {
	loaded = false;

	id = 0;
	subdomain = "";
	created_dt = new Date();
	paused = false;
	app_id = 0;
	app_version = '';
	upgrade :AppVersion|undefined;

	async fetch(id: number) {
		const resp_data = await get('/appspace/'+id);
		this.setFromRaw(resp_data);
	}
	async refresh() {
		await this.fetch(this.id);
	}
	setFromRaw(raw :any) {
		this.id = Number(raw.appspace_id);
		this.subdomain = raw.subdomain+'';
		this.created_dt = new Date(raw.created_dt);
		this.paused = !!raw.paused;
		this.app_id = Number(raw.app_id);
		this.app_version = raw.app_version+'';

		if( raw.upgrade ) {
			this.upgrade = new AppVersion;
			this.upgrade.setFromRaw(raw.upgrade)
		}

		this.loaded = true;
	}
	
	// actions:
	async setPause(pause :boolean) {
		const data = await post('/appspace/'+this.id+'/pause', {pause});
		this.paused = pause;
	}
}

export class Appspaces {
	loaded = false;

	as : Map<number,Appspace> = new Map();

	async fetchForOwner() {
		const resp_data = await get('/appspace');
		resp_data.appspaces.forEach( (raw:any) => {
			const appspace = new Appspace;
			appspace.setFromRaw(raw);
			this.as.set(appspace.id, appspace);
		});
		this.loaded = true;
	}

	get asArray() : Appspace[] {
		// maybe this should return an empty array if all_loaded === false
		// Otherwise, some views might load some appspaces, then the appspace view will render a partial list.
		return Array.from(this.as.values());
	}
}

export function ReactiveAppspaces() {
	return reactive(new Appspaces);
}

export async function createAppspace(app_id:number, app_version:string) :Promise<number> {
	const resp_data = await post('/appspace', {app_id, app_version});
	return Number(resp_data.appspace_id);
}