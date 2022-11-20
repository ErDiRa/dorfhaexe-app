<template>
	<div :class="$style.dateElement">
		<div :class="$style.textFields">
			<div :class="$style.event">{{ event }}</div>
			<div :class="$style.meta">{{ date }}</div>
			<div v-if="time && time !== '-'" :class="$style.meta">{{ time }} Uhr</div>
			<div v-if="outfit" :class="$style.meta">{{ outfit }}</div>
			<div v-if="meetingPoint" :class="$style.meta">{{ meetingPoint }}</div>
		</div>
		<a
			ref="downloadTag"
			@click="download"
			:class="$style.icsLink"
			type="text/calendar"
		>
			<div :class="$style.icon">
				<calendar></calendar>
			</div>
		</a>
	</div>
</template>

<script setup>
	import { onMounted, ref } from 'vue';
	import Calendar from '../../assets/calendar.svg';
	const props = defineProps({
		event: String,
		date: String,
		time: String,
		outfit: String,
		icsFile: String,
		icsFileName: String,
		meetingPoint: String
	});

	const downloadTag = ref(null);

	onMounted(() => {
		if (downloadTag) {
			downloadTag.value.setAttribute(
				'href',
				'data:text/calendar;charset=utf-8,' + encodeURIComponent(props.icsFile)
			);
			downloadTag.value.setAttribute('download', props.icsFileName);
		}
	});

	//TODO: implement download like this: https://stackoverflow.com/questions/53772331/vue-html-js-how-to-download-a-file-to-browser-using-the-download-tag
	// so that a dialog window can be shown before downloading
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
				min-width: 34rem;
				white-space: pre-wrap;
			}

			@media (max-width: 650px) {
				.event {
					max-width: 15rem;
					min-width: 0;
				}
			}

			.meta {
				color: #605b5b;
				font-weight: 500;
			}

			div + div {
				margin-top: 0.5rem;
			}
		}

		.icsLink {
			color: inherit;
			display: flex;
			margin-left: auto;
			margin-right: 1rem;
			.icon {
				margin-left: auto;
				color: #555555;
			}
		}
	}
</style>
