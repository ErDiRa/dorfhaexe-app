import { createRouter, createWebHashHistory } from 'vue-router';
import About from './components/pages/about.vue';
import Contact from './components/pages/contact.vue';
import Dates from './components/pages/dates.vue';
import History from './components/pages/history.vue';
import Home from './components/pages/home.vue';
import News from './components/pages/news.vue';

import { navigation } from './const/strings';

// Define some routes
const routes = [
	{
		path: navigation.home.ROUTE,
		component: Home
	},
	{
		path: navigation.dates.ROUTE,
		component: Dates
	},
	{
		path: navigation.news.ROUTE,
		component: News
	},
	{
		path: navigation.about.ROUTE,
		component: About
	},
	{
		path: navigation.contact.ROUTE,
		component: Contact
	},
	{
		path: navigation.history.ROUTE,
		component: History
	}
];

const router = createRouter({
	// 4. Provide the history implementation to use. We are using the hash history for simplicity here.
	history: createWebHashHistory(),
	routes
});

export default router;
