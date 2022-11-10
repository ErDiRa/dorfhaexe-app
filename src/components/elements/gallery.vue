<template>
	<div :class="$style.card">
		<button :class="$style.chevronLeft" @click="previousImage()">
			<ChevronLeft></ChevronLeft>
		</button>
		<div v-if="enlarged" :class="$style.backdrop"></div>
		<Motion
			:key="current"
			ref="imageContainer"
			:initial="{ opacity: 0 }"
			:animate="{
				opacity: 1
			}"
			:exit="{ opacity: 0 }"
			:class="
				enlarged
					? [$style.slide, $style.scaled]
					: [$style.slide, $style.unscaled]
			"
			@click="enlarge"
		>
			<img
				:src="imageOneMedium"
				:srcset="`${imageOneSmall} 1024w,
									${imageOneMedium} 1440w,
									${imageOneLarge} 2400w`"
				v-show="current === 0"
				:class="$style.image"
			/>
			<img
				:src="imageTwoMedium"
				:srcset="`${imageTwoSmall} 1024w,
									${imageTwoMedium} 1440w,
									${imageTwoLarge} 2400w`"
				v-show="current === 1"
				:class="$style.image"
			/>
			<img
				:src="imageThreeMedium"
				:srcset="`${imageThreeSmall} 1024w,
									${imageThreeMedium} 1440w,
									${imageThreeLarge} 2400w`"
				v-show="current === 2"
				:class="$style.image"
			/>
			<img
				:src="imageFourMedium"
				:srcset="`${imageFourSmall} 1024w,
									${imageFourMedium} 1440w,
									${imageFourLarge} 2400w`"
				v-show="current === 3"
				:class="$style.image"
			/>
		</Motion>
		<button @click="nextImage()" :class="$style.chevronRight">
			<ChevronRight></ChevronRight>
		</button>
	</div>
</template>

<script setup>
	import { Motion } from 'motion/vue';
	import { ref } from 'vue';
	import ChevronLeft from '../../assets/chevron-left.svg';
	import ChevronRight from '../../assets/chevron-right.svg';
	import imageOneSmall from '../../assets/group-1-1024w.JPG';
	import imageOneMedium from '../../assets/group-1-1440w.JPG';
	import imageOneLarge from '../../assets/group-1-2400w.JPG';
	import imageTwoSmall from '../../assets/group-2-1024w.JPG';
	import imageTwoMedium from '../../assets/group-2-1440w.JPG';
	import imageTwoLarge from '../../assets/group-2-2400w.JPG';
	import imageThreeSmall from '../../assets/group-3-1024w.JPG';
	import imageThreeMedium from '../../assets/group-3-1440w.JPG';
	import imageThreeLarge from '../../assets/group-3-2400w.JPG';
	import imageFourSmall from '../../assets/group-4-1024w.JPG';
	import imageFourMedium from '../../assets/group-4-1440w.JPG';
	import imageFourLarge from '../../assets/group-4-2400w.JPG';

	let current = ref(0);

	let imageContainer = ref(null);

	let enlarged = ref(false);

	const enlarge = () => {
		enlarged.value = !enlarged.value;
		return;
	};

	const nextImage = () => {
		current.value++;
		if (current.value > 3) {
			current.value = 0;
		}
	};

	const previousImage = () => {
		current.value--;
		if (current.value < 0) {
			current.value = 3;
		}
	};
</script>

<style lang="scss" module>
	.backdrop {
		position: fixed;
		top: 0;
		right: 0;
		bottom: 0;
		left: 0;
		background-color: rgba(0, 0, 0, 0.5);
		transition: all 0.25s ease;
	}

	.scaled {
		transform: scale(1.5);
		transition: 0.25s ease;
	}

	.unscaled {
		transform: scale(1);
		transition: 0.25s ease;
	}
	.card {
		position: relative;
		display: flex;
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
				margin-right: 0.25rem;
				svg {
					height: 1.5rem;
					width: auto;
				}
			}

			.chevronLeft {
				margin-left: 0.25rem;
				svg {
					height: 1.5rem;
					width: auto;
				}
			}
		}
	}

	@media (max-width: 600px) {
		.card {
			padding: 1rem 0;
		}
	}
</style>
