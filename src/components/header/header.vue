<template>
	<div :class="$style.container">
		<div :class="$style.logo">
			<haexe-logo @click="router.push(navigation.home.ROUTE)"></haexe-logo>
		</div>
		<menu-view
			@navigate-to="(route) => $emit('navigateTo', route)"
			:class="$style.menu"
		></menu-view>
	</div>
	<hr :class="$style.separator" />
</template>

<script setup>
	import { useRouter } from 'vue-router';
	import HaexeLogo from '../../assets/dorfhaexe-emblem.svg';
	import { navigation } from '../../const/strings';
	import MenuView from './menu.vue';

	const router = useRouter();
</script>

<style lang="scss" module>
	.container {
		display: flex;
		.logo {
			z-index: 500;
		}
		.menu {
			margin-left: auto;
			margin-top: auto;
			margin-bottom: auto;
		}
	}
	.separator {
		border: 0;
		border-top: 2px solid #d9d9d9;
		z-index: 501;
	}
</style>
