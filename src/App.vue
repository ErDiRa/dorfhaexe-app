<template>
	<div>
		<header-view />
		<router-view :class="$style.main"></router-view>
		<footer-view />
	</div>
</template>

<script setup>
	import { RouterView } from 'vue-router';
	import FooterView from './components/footer/footer.vue';
	import HeaderView from './components/header/header.vue';
</script>
<style lang="scss" module>
	.main {
		display: flex;
		flex-direction: column;
	}
</style>
