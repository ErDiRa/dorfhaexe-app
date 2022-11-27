<template>
	<div :class="$style.container">
		<div
			@click="open"
			:class="isOpen ? [$style.button, $style.active] : [$style.button]"
		>
			<span>{{ title }}</span>
			<ChevronDown v-if="!isOpen" :class="$style.chevron"></ChevronDown>
			<ChevronUp v-else :class="$style.chevron"></ChevronUp>
		</div>
		<div :class="$style.content" ref="content">
			<slot></slot>
		</div>
	</div>
</template>

<script setup>
	import { onMounted, ref } from 'vue';
	import ChevronDown from '../../assets/chevron-down.svg';
	import ChevronUp from '../../assets/chevron-up.svg';

	const content = ref(null);

	const props = defineProps({
		title: String,
		initIsOpen: Boolean
	});

	let isOpen = ref(false);

	const open = () => {
		isOpen.value = !isOpen.value;
		if (isOpen.value && content.value) {
			content.value.style.maxHeight = content.value.scrollHeight + 'px';
		} else {
			content.value.style.maxHeight = 0;
		}
	};

	onMounted(() => {
		if (props.initIsOpen) {
			open();
		}
	});
</script>

<style lang="scss" module>
	.container {
		.button {
			border-style: none;
			background-color: #fef7f7;
			font-size: 1rem;
			width: 100%;
			height: auto;
			font-weight: 600;
			display: flex;
			justify-content: flex-start;
			align-items: center;
			color: inherit;
			padding: 1rem 0;
			border-bottom: 1px solid #dedcdc;
			cursor: pointer;

			span {
				font-size: 1.5rem;
			}

			@media (max-width: 600px) {
				span {
					font-size: 1rem;
				}
			}

			.chevron {
				margin-left: auto;
			}
		}
		.active {
			color: #75a366;
		}
		.content {
			max-height: 0;
			overflow: hidden;
			transition: max-height 0.2s ease-out;
			border-radius: 5px;
			border-top-left-radius: 0;
			border-top-right-radius: 0;
		}
	}
</style>
