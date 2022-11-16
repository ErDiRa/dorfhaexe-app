<template>
	<main :class="$style.main">
		<h1>Termine 2022/2023</h1>
		<collapsable
			:title="'Kampagne'"
			:data="kampagneData"
			:class="$style.collapsable"
		>
		</collapsable>
		<collapsable
			:title="'Narrenfahrplan'"
			:data="kampagneData"
			:class="$style.collapsable"
		>
		</collapsable>
		<collapsable
			:title="'Auf-/Abbauplan'"
			:data="kampagneData"
			:class="$style.collapsable"
		>
		</collapsable>
	</main>
</template>

<script setup>
	import { ref } from 'vue';
	import Collapsable from '../elements/collapsable.vue';

	import kampagneData from '../../data/kampagne.json';

	const kampagne = ref(kampagneData);
</script>

<style lang="scss" module>
	.main {
		min-height: 100vh;
		display: flex;
		flex-direction: column;

		.collapsable {
			margin-left: auto;
			margin-right: auto;
			max-width: 400px;
		}

		.collapsable + .collapsable {
			margin-top: 0.5rem;
		}
	}
	.tabList {
		box-sizing: border-box;
		border: 0 solid #e5e7eb;
		display: flex;
		padding: 0.25rem;
		border-radius: 0.5rem;
		background-color: rgba(189, 189, 189, 0.251);
		.tab {
			background-color: transparent;
			background-image: none;
			border-style: none;
			border-radius: 0.5rem;
			padding-top: 0.75rem;
			padding-bottom: 0.75rem;
			font-size: 1.5rem;
			line-height: 1.25rem;
			font-weight: 600;
			line-height: 1.25rem;
			width: 100%;
		}

		.tab + .tab {
			margin-left: 0.25rem;
		}
		.selected {
			background-color: #ffffff;
			box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.1),
				0 1px 2px 0 rgba(0, 0, 0, 0.06);
		}

		.unselected {
			color: #9a9c9f;
			:hover {
				color: #ffffff;
			}
		}
	}

	.card {
		position: relative;
		display: flex;
		align-content: center;
		justify-content: center;
		border-radius: 5px;
		max-height: 40rem;
		padding: 1rem;
		background-color: #ffffea;
		box-shadow: 0.5px 0.5px 1.6px rgba(0, 0, 0, 0.022),
			1.1px 1.1px 2.4px rgba(0, 0, 0, 0.031),
			2.3px 2.4px 4.5px rgba(0, 0, 0, 0.039),
			5.7px 5.9px 21.3px rgba(0, 0, 0, 0.048),
			36px 37px 80px rgba(0, 0, 0, 0.07);
	}

	.card + .card {
		margin-top: 0.75rem;
	}
</style>
