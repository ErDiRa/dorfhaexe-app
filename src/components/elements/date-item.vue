<template>
	<div :class="$style.dateElement">
		<div :class="$style.textFields">
			<div :class="$style.event">
				<span>{{ event }}</span>
			</div>
			<div :class="$style.meta">
				<calendar :class="$style.metaIcon"></calendar>
				<span>{{ date }}</span>
			</div>
			<div v-if="time && time !== '-'" :class="$style.meta">
				<clock :class="$style.metaIcon"></clock>
				<span>{{ time }} Uhr</span>
			</div>
			<div v-if="outfit" :class="$style.meta">
				<user :class="$style.metaIcon"></user>
				<span>{{ outfit }}</span>
			</div>
			<div v-if="meetingPoint" :class="$style.meta">
				<map-pin :class="$style.metaIcon"></map-pin>
				<span>{{ meetingPoint }}</span>
			</div>
		</div>
		<a ref="downloadTag" :class="$style.icsLink" type="text/calendar">
			<div :class="$style.icon">
				<download></download>
			</div>
		</a>
	</div>
</template>

<script setup>
	import { onMounted, ref } from 'vue';
	import Calendar from '../../assets/calendar.svg';
	import Clock from '../../assets/clock.svg';
	import Download from '../../assets/download.svg';
	import MapPin from '../../assets/map.svg';
	import user from '../../assets/user.svg';
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
				margin-bottom: 0.5rem;
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

				.metaIcon {
					width: 12px;
					height: 12px;
					margin-right: 0.5rem;
				}
			}

			div + div {
				margin-top: 0.5rem;
			}
		}

		.icsLink {
			color: inherit;
			display: flex;
			margin-left: auto;
			.icon {
				margin-left: auto;
				color: #555555;
			}
		}
	}
</style>
