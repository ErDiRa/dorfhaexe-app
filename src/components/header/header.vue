<template>
	<div :class="$style.container">
		<div :class="$style.logo">
			<haexe-logo></haexe-logo>
		</div>
		<div :class="$style.menu">
			<menu-view @navigate-to="(route) => $emit('navigateTo', route)"></menu-view>
		</div>
	</div>
	<hr :class="$style.separator"/>
</template>

<script setup>
import HaexeLogo from '../../assets/dorfhaexe-emblem.svg';
import MenuView from './menu.vue';
const emit = defineEmits(['navigateTo'])
</script>

<style lang="scss" module>
.container{
	display: flex;
	.logo{
		z-index: 500;
	}
	.menu{
		margin-left: auto;
	}
}
.separator {
	border: 0;
	border-top: 2px solid #D9D9D9;
	z-index: 501;
}
</style>