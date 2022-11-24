<template>
	<div :class="$style.card">
		<div :class="$style.profile"></div>
		<div :class="$style.name">{{ person.name }}</div>
		<div :class="$style.position">{{ person.position }}</div>
	</div>
</template>

<script setup>
	import { defineProps } from 'vue';

	const props = defineProps({
		person: Object
	});
</script>

<style lang="scss" module>
	.card {
		position: relative;
		display: flex;
		flex-direction: column;
		align-content: center;
		justify-content: center;
		border-radius: 5px;
		max-height: 40rem;
		margin-top: 1rem;
		margin-bottom: 2rem;
		padding: 1rem;
		background-color: #ffffea;
		box-shadow: 0.5px 0.5px 1.6px rgba(0, 0, 0, 0.022),
			1.1px 1.1px 2.4px rgba(0, 0, 0, 0.031),
			2.3px 2.4px 4.5px rgba(0, 0, 0, 0.039),
			5.7px 5.9px 21.3px rgba(0, 0, 0, 0.048);

		.profile {
			background-color: #c9c6c6;
			border-radius: 50%;
			height: 100px;
			width: 100px;
			margin-left: auto;
			margin-right: auto;
			margin-bottom: 1rem;
		}

		.name {
			font-weight: 600;
			margin-bottom: 0.5rem;
		}
	}
</style>
