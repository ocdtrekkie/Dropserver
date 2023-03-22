import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';
import Home from '../views/Home.vue';
import User from '../views/User.vue';

import Appspaces from '../views/Appspaces.vue';
import ManageAppspace from '../views/ManageAppspace.vue';
import MigrateAppspace from '../views/MigrateAppspace.vue';
import RestoreAppspace from '../views/RestoreAppspace.vue';
import ManageAppspaceUser from '../views/ManageAppspaceUser.vue';
import NewAppspace from '../views/NewAppspace.vue';

import NewRemoteAppspace from '../views/NewRemoteAppspace.vue';
import ManageRemoteAppspace from '../views/ManageRemoteAppspace.vue';

import Apps from '../views/Apps.vue';
import ManageApp from '../views/ManageApp.vue';
import NewAppVersion from '../views/NewAppVersion.vue';
import NewAppVersionInProcess from '../views/NewAppVersionInProcess.vue';
import NewApp from '../views/NewApp.vue';
import NewAppInProcess from '../views/NewAppInProcess.vue';

import Contacts from '../views/Contacts.vue';
import ManageContact from '../views/ManageContact.vue';
import NewContact from '../views/NewContact.vue';

import NewDropID from '../views/NewDropID.vue';
import ManageDropID from '../views/ManageDropID.vue';

import AdminHome from '../views/admin/AdminHome.vue';
import Users from '../views/admin/Users.vue';
import AdminSettings from '../views/admin/AdminSettings.vue';

const routes: Array<RouteRecordRaw> = [
	{
		path: '/',
		name: 'home',
		component: Home
	},{
		path: '/user',
		name: 'user',
		component: User
	},{
		path: '/appspace',
		name: 'appspaces',
		component: Appspaces
	},{
		path: '/appspace/:id',
		name: 'manage-appspace',
		component: ManageAppspace
	},{
		path: '/appspace/:id/migrate',
		name: 'migrate-appspace',
		component: MigrateAppspace
	},{
		path: '/appspace/:appspace_id/restore',
		name: 'restore-appspace',
		component: RestoreAppspace,
		props: true
	},{
		path: '/appspace/:id/new-user',
		name: 'appspace-new-user',
		component: ManageAppspaceUser
	},{
		path: '/appspace/:id/user/:proxy_id',
		name: 'appspace-user',
		component: ManageAppspaceUser
	},{
		path: '/new-appspace/',
		name: 'new-appspace',
		component: NewAppspace
	},{
		path: '/remote-appspace/:domain',
		name: 'manage-remote-appspace',
		component: ManageRemoteAppspace,
		props: true
	},{
		path: '/new-remote-appspace/',
		name: 'new-remote-appspace',
		component: NewRemoteAppspace
	},{
		path: '/app',
		name: 'apps',
		component: Apps
	},{
		path: '/app/:id',
		name: 'manage-app',
		component: ManageApp,
		props: true
	},{
		path: '/app/:id/new-version',
		name: 'new-app-version',
		component: NewApp,
		props: route => {
			const p = route.params.id;
			if( Array.isArray(p) ) throw new Error("id can not be an array");
			return {
				app_id: parseInt(route.params.id as string)
			}
		}
	},{
		path: '/app/:id/new-version/:app_get_key',
		name: 'new-app-version-in-process',
		component: NewAppInProcess,
		props: route => {
			const p = route.params.id;
			if( Array.isArray(p) ) throw new Error("id can not be an array");
			return {
				app_id: parseInt(route.params.id as string),
				app_get_key: route.params.app_get_key
			}
		}
	},{
		path: '/new-app',
		name: 'new-app',
		component: NewApp
	},{
		path: '/new-app/:app_get_key',
		name: 'new-app-in-process',
		component: NewAppInProcess,
		props: true
	},{
		path: '/contact',
		name: 'contacts',
		component: Contacts
	},{
		path: '/contact/:contact_id',
		name: 'manage-contact',
		component: ManageContact
	},{
		path: '/contact-new',
		name: 'new-contact',
		component: NewContact
	},{
		path: '/dropid-new',
		name: 'new-dropid',
		component: NewDropID
	},{
		path: '/dropid',
		name: 'dropid',
		component: ManageDropID
	},{
		path: '/admin',
		name: 'admin',
		component: AdminHome
	},{
		path: '/admin/users',
		name: 'admin-users',
		component: Users
	},{
		path: '/admin/settings',
		name: 'admin-settings',
		component: AdminSettings
	}
];

const router = createRouter({
	history: createWebHistory(import.meta.env.BASE_URL),
	routes
});

export default router;
