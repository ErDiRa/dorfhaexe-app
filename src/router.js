import { computed, onMounted, onUnmounted, ref } from 'vue';
import About from './components/pages/about.vue';
import Contact from './components/pages/contact.vue';
import Dates from './components/pages/dates.vue';
import Home from './components/pages/home.vue';

import { navigation } from './const/strings';

// Define some routes
const routes = {
	'/': {
		path: navigation.home.ROUTE,
		component: Home
	},
	'/termine': {
		path: navigation.dates.ROUTE,
		component: Dates
	},
	'/ueber_uns': {
		path: navigation.about.ROUTE,
		component: About
	},
	'/kontakt': {
		path: navigation.contact.ROUTE,
		component: Contact
	}
};

export function useRouter() {
	const currentPath = ref(window.location.hash);

	const setCurrentPath = () => {
		currentPath.value = window.location.hash;
	};

	onMounted(() => {
		window.addEventListener('hashchange', setCurrentPath);
	});

	onUnmounted(() => {
		window.removeEventListener('hashchange', setCurrentPath);
	});

	const currentView = computed(() => {
		return routes[currentPath.value.slice(1) || '/'] || NotFound;
	});

	return {
		currentView
	};
}
