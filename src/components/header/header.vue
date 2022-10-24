<template>
	<div>
		<menu-view @navigate-to="(route) => $emit('navigateTo', route)"></menu-view>
	</div>
</template>

<script setup>
import MenuView from './menu.vue';
const emit = defineEmits(['navigateTo'])
</script>