<template>
	<main :class="$style.main">
		<h1 :class="$style.title">Termine 2022/2023</h1>
		<p :class="$style.info">
			Hier findet ihr unsere aktuellen Termine. <br />Vielleicht sieht man sich
			ja hin und wieder auf der ein oder anderen Verstanstaltung. Wir würden uns
			freuen 🤟.
		</p>
		<header-collapsable :title="'Kampagne'" :class="$style.collapsable">
			<div v-for="val in kampagneData" :class="$style.card">
				<date-item
					:event="val.event"
					:date="val.date"
					:time="val.time"
					:outfit="val.outfit"
					:meeting-point="val.meeting_point"
					:ics-file="val.ics_file"
					:ics-file-name="val.ics_file_name"
				></date-item>
			</div>
		</header-collapsable>
		<header-collapsable :title="'Narrenfahrplan'" :class="$style.collapsable">
			<div v-for="val in narrenfahrplanData" :class="$style.card">
				<date-item
					:event="val.event"
					:date="val.date"
					:time="val.time"
					:outfit="val.outfit"
					:meeting-point="val.meeting_point"
					:ics-file="val.ics_file"
					:ics-file-name="val.ics_file_name"
				></date-item>
			</div>
		</header-collapsable>
		<header-collapsable :title="'Auf-/Abbauplan'" :class="$style.collapsable">
			<div v-for="val in aufabbauData" :class="$style.card">
				<date-item
					:event="val.event"
					:date="val.date"
					:time="val.time"
					:outfit="val.outfit"
					:meeting-point="val.meeting_point"
					:ics-file="val.ics_file"
					:ics-file-name="val.ics_file_name"
				></date-item>
			</div>
		</header-collapsable>
	</main>
</template>

<script setup>
	import { ref } from 'vue';
	import DateItem from '../elements/date-item.vue';
	import HeaderCollapsable from '../elements/header-collapsable.vue';

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

		.title {
			font-size: 2.5rem;
			text-align: left;
		}

		@media (max-width: 600px) {
			.title {
				font-size: 1.5rem;
			}
		}

		.info {
			white-space: inherit;
			margin-top: 0;
			margin-bottom: 2rem;
			font-size: 1.5rem;
			line-height: 1.9rem;
			text-align: left;
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
			.card {
				position: relative;
				display: flex;
				cursor: pointer;
				flex-direction: column;
				align-content: center;
				justify-content: center;
				border-bottom: 1px solid #e2e0e0;
				max-height: 40rem;
				padding: 1rem;
				background: #ffffea;
				max-width: 40rem;
				margin-left: auto;
				margin-right: auto;
			}

			.card:first-child {
				margin-top: 16px;
				border-top-left-radius: 8px;
				border-top-right-radius: 8px;
			}

			.card:last-child {
				border-bottom-left-radius: 8px;
				border-bottom-right-radius: 8px;
			}
			.card:hover {
				background-color: #eaead585;
			}
		}

		.collapsable + .collapsable {
			margin-top: 0.5rem;
		}
	}
</style>
