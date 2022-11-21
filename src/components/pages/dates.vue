<template>
	<main :class="$style.main">
		<h1>Termine 2022/2023</h1>
		<p :class="$style.info">
			Hier findet ihr unsere aktuellen Termine. <br />Vielleicht sieht man sich
			ja hin und wieder auf der ein oder anderen Verstanstaltung. Wir würden uns
			freuen 🤟.
		</p>
		<collapsable
			:title="'Kampagne'"
			:data="kampagneData"
			:class="$style.collapsable"
		>
		</collapsable>
		<collapsable
			:title="'Narrenfahrplan'"
			:data="narrenfahrplanData"
			:class="$style.collapsable"
		>
		</collapsable>
		<collapsable
			:title="'Auf-/Abbauplan'"
			:data="aufabbauData"
			:class="$style.collapsable"
		>
		</collapsable>
	</main>
</template>

<script setup>
	import { ref } from 'vue';
	import Collapsable from '../elements/collapsable.vue';

	import aufabbauData from '../../data/auf-abbau.json';
	import kampagneData from '../../data/kampagne.json';
	import narrenfahrplanData from '../../data/narrenfahrplan.json';

	const kampagne = ref(kampagneData);
</script>

<style lang="scss" module>
	.main {
		min-height: 100vh;
		display: flex;
		flex-direction: column;

		.info {
			white-space: pre-wrap;
			min-width: 35rem;
			max-width: 35rem;
			margin-left: auto;
			margin-right: auto;
			margin-top: 0;
			margin-bottom: 2rem;
			font-size: 1.5rem;
			line-height: 1.9rem;
		}

		@media (max-width: 650px) {
			.info {
				font-size: 1rem;
				line-height: 1.4rem;
				max-width: 21rem;
				min-width: 0;
			}
		}
		.collapsable {
			margin-left: auto;
			margin-right: auto;
			max-width: 600px;
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
