<template>
	<div :class="$style.card">
		<div
			:class="$style.profile"
			:style="{ 'background-image': `url(${getImageUrl(person.profileImg)})` }"
			@click="openEnlargedImage()"
		></div>
		<div :class="enlarged ? [$style.modal, $style.show] : $style.modal">
			<div :class="$style.backdrop" @click="openEnlargedImage()"></div>
			<div
				:class="$style.enlargedImage"
				:style="{
					'background-image': `url(${getImageUrl(person.profileImg)})`
				}"
				@click="openEnlargedImage()"
			></div>
		</div>
		<div :class="$style.name">{{ person.name }}</div>
		<div :class="$style.position">{{ person.position }}</div>
	</div>
</template>

<script setup>
	import { ref } from 'vue';
	const props = defineProps({
		person: Object
	});

	const getImageUrl = (url) => {
		return new URL(url, import.meta.url).href;
	};

	const enlarged = ref(false);

	const openEnlargedImage = () => {
		enlarged.value = !enlarged.value;
		//https://css-tricks.com/prevent-page-scrolling-when-a-modal-is-open/
		if (enlarged.value) {
			document.body.style.position = 'fixed';
			document.body.style.top = `-${window.scrollY}px`;
		} else {
			const scrollY = document.body.style.top;
			document.body.style.position = '';
			document.body.style.top = '';
			window.scrollTo(0, parseInt(scrollY || '0') * -1);
		}
	};
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
			height: 115px;
			width: 115px;
			margin-left: auto;
			margin-right: auto;
			margin-bottom: 1rem;
			background-repeat: no-repeat;
			background-size: 170px auto;
			background-position: center;
		}

		.name {
			font-weight: 600;
			margin-bottom: 0.5rem;
		}

		.modal {
			position: fixed;
			top: 0;
			right: 0;
			bottom: 0;
			left: 0;
			display: none;
			opacity: 0;
			transition: opacity 0.2s ease-out;
			z-index: 500;

			.backdrop {
				width: 100%;
				position: fixed;
				top: 0;
				right: 0;
				bottom: 0;
				left: 0;
				background-color: rgba(0, 0, 0, 0.5);
			}

			.enlargedImage {
				position: relative;
				left: 50%;
				top: 50%;
				transform: translate(-50%, -50%);
				border-radius: 0.25rem;
				height: 420px;
				max-width: 90%;
				background-repeat: no-repeat;
				background-size: auto 90%;
				background-position: center;
			}
		}
		.show {
			display: block;
			opacity: 1;
		}
	}
</style>
