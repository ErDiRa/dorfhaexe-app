<template>
	<div>
		<header-view />
		<main :class="$style.main">
			<component :is="mainComponent" />
		</main>
		<footer-view />
	</div>
</template>

<script setup>
	import { computed, provide } from 'vue';

	import FooterView from './components/footer/footer.vue';
	import HeaderView from './components/header/header.vue';
	import Home from './components/pages/home.vue';
	import { useRouter } from './router';

	import { navigation } from './const/strings';

	const router = useRouter();

	provide('currentView', router);

	let mainComponent = computed(() => {
		if (!router.value) {
			router.value = {
				path: navigation.home.ROUTE,
				component: Home
			};
			return router.value.component;
		}
		return router.value.component;
	});
</script>
<style lang="scss" module>
	.main {
		display: flex;
		flex-direction: column;
	}
</style>
