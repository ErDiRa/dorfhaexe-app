<template>
	<div :class="$style.card">
		<button :class="$style.chevronLeft" @click="previousImage()">
			<ChevronLeft></ChevronLeft>
		</button>
		<Motion
			:key="current"
			:initial="{ opacity: 0 }"
			:animate="{
				opacity: 1
			}"
			:exit="{ opacity: 0 }"
			:class="$style.slide"
		>
			<img :src="imageTwo" v-show="current === 0" :class="$style.image" />
			<img :src="imageTwo" v-show="current === 1" :class="$style.image" />
			<img :src="imageThree" v-show="current === 2" :class="$style.image" />
			<img :src="imageFour" v-show="current === 3" :class="$style.image" />
		</Motion>
		<button @click="nextImage()" :class="$style.chevronRight">
			<ChevronRight></ChevronRight>
		</button>
	</div>
</template>

<script setup>
	import { Motion } from 'motion/vue';
	import { ref, watch } from 'vue';
	import ChevronLeft from '../../assets/chevron-left.svg';
	import ChevronRight from '../../assets/chevron-right.svg';
	import imageTwo from '../../assets/group-2.JPG';
	import imageThree from '../../assets/group-3.JPG';
	import imageFour from '../../assets/group-4.JPG';
	import { useSwipe } from '../../composables/swipe';

	//TODO: resize images using: ImageMagick https://www.smashingmagazine.com/2015/06/efficient-image-resizing-with-imagemagick/

	let current = ref(0);

	const nextImage = () => {
		current.value++;
		if (current.value > 3) {
			current.value = 0;
		}
	};

	const { xDiff, yDiff } = useSwipe();

	// listen to touch events
	watch([xDiff, yDiff], ([newXDiff], [newYDiff]) => {
		if (Math.abs(newXDiff) > Math.abs(newYDiff)) {
			/*most significant*/
			if (xDiff > 0) {
				nextImage();
			} else {
				previousImage();
			}
		}
	});

	const previousImage = () => {
		current.value--;
		if (current.value < 0) {
			current.value = 3;
		}
	};
</script>

<style lang="scss" module>
	.card {
		position: relative;
		display: flex;
		align-content: center;
		justify-content: center;
		border-radius: 5px;
		max-height: 40rem;
		margin-top: 2.25rem;
		margin-bottom: 3.75rem;
		padding: 1rem;
		background-color: #ffffea;
		box-shadow: 0.5px 0.5px 1.6px rgba(0, 0, 0, 0.022),
			1.1px 1.1px 2.4px rgba(0, 0, 0, 0.031),
			2.3px 2.4px 4.5px rgba(0, 0, 0, 0.039),
			5.7px 5.9px 21.3px rgba(0, 0, 0, 0.048),
			36px 37px 80px rgba(0, 0, 0, 0.07);

		.slide {
			border-radius: 5px;
			color: var(--white);
			display: flex;
			justify-content: center;
			align-items: center;
			position: relative;
			overflow: hidden;
			max-width: 70%;
			.image {
				max-width: 100%;
				max-height: 100%;
			}
		}

		.chevronLeft {
			margin-right: auto;
			margin-left: 1rem;
			margin-top: 1rem;
			margin-bottom: 1rem;
			border-style: none;
			background-color: transparent;
			cursor: pointer;
			svg {
				height: 48px;
				width: 48px;
				stroke-width: 3;
				padding: 0.2rem;
				color: white;
				background-color: rgba(0, 0, 0, 0.4);
				border-radius: 5px;
			}

			svg:hover {
				background-color: rgba(0, 0, 0, 0.8);
			}
		}

		.chevronRight {
			margin-right: 1rem;
			margin-left: auto;
			margin-top: 1rem;
			margin-bottom: 1rem;
			border-style: none;
			background-color: transparent;
			cursor: pointer;

			svg {
				height: 48px;
				width: 48px;
				stroke-width: 3;
				padding: 0.2rem;
				color: white;
				background-color: rgba(0, 0, 0, 0.4);
				border-radius: 5px;
			}

			svg:hover {
				background-color: rgba(0, 0, 0, 0.8);
			}
		}

		@media (max-width: 600px) {
			.chevronRight {
				display: none;
			}

			.chevronLeft {
				display: none;
			}
		}
	}
</style>
