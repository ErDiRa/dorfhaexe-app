import { onMounted, onUnmounted, ref, shallowRef, watch } from 'vue';
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
	const currentPath = ref(window.location.pathname);

	const setCurrentPath = () => {
		currentPath.value = window.location.pathname;
	};

	const currentView = shallowRef(null);
	watch(
		() => currentPath.value,
		(newVal) => {
			currentView.value = routes[newVal] || null;
		}
	);

	onMounted(() => {
		window.addEventListener('hashchange', setCurrentPath);
	});

	onUnmounted(() => {
		window.removeEventListener('hashchange', setCurrentPath);
	});

	return currentView;
}
