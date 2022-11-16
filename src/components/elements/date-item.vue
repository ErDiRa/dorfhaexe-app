<template>
	<div :class="$style.dateElement">
		<div :class="$style.textFields">
			<div :class="$style.event">{{ event }}</div>
			<div :class="$style.meta">{{ date }}</div>
			<div v-if="time" :class="$style.meta">{{ time }} Uhr</div>
			<div v-if="outfit" :class="$style.meta">{{ outfit }}</div>
		</div>
		<div :class="$style.icon">
			<calendar></calendar>
		</div>
	</div>
</template>

<script setup>
	import Calendar from '../../assets/calendar.svg';
	const props = defineProps({
		event: String,
		date: String,
		time: String,
		outfit: String
	});

	//TODO: use https://github.com/nwcell/ics.js to provide calender file
</script>

<style lang="scss" module>
	.dateElement {
		display: flex;

		.textFields {
			display: flex;
			flex-direction: column;
			align-items: flex-start;
			.event {
				font-weight: 600;
				text-align: left;
			}

			.meta {
				color: #605b5b;
				font-weight: 500;
			}

			div + div {
				margin-top: 0.5rem;
			}
		}

		.icon {
			margin-left: auto;
		}
	}
</style>
