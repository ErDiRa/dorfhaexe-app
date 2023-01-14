<template>
	<div :class="$style.container">
		<button
			@click="open"
			:class="isOpen ? [$style.button, $style.active] : [$style.button]"
		>
			<span>{{ title }}</span>
			<ChevronDown v-if="!isOpen" :class="$style.chevron"></ChevronDown>
			<ChevronUp v-else :class="$style.chevron"></ChevronUp>
		</button>
		<div :class="$style.content" ref="content">
			<slot></slot>
		</div>
	</div>
</template>

<script setup>
	import { ref } from 'vue';
	import ChevronDown from '../../assets/chevron-down.svg';
	import ChevronUp from '../../assets/chevron-up.svg';

	const content = ref(null);

	const props = defineProps({
		title: String
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
</script>

<style lang="scss" module>
	.container {
		.button {
			border-style: none;
			background-color: #ffffea;
			font-size: 1rem;
			width: 100%;
			height: 3rem;
			font-weight: 600;
			display: flex;
			justify-content: flex-start;
			align-items: center;
			padding: 1rem;
			color: inherit;
			cursor: pointer;
			box-shadow: 0.5px 0.5px 1.6px rgba(0, 0, 0, 0.022),
				1.1px 1.1px 2.4px rgba(0, 0, 0, 0.031),
				2.3px 2.4px 4.5px rgba(0, 0, 0, 0.039),
				5.7px 5.9px 21.3px rgba(0, 0, 0, 0.048),
				36px 37px 80px rgba(0, 0, 0, 0.07);

			.chevron {
				margin-left: auto;
				margin-right: 0.5rem;
			}
		}

		.active {
			color: #75a366;
		}
		.content {
			max-height: 0;
			overflow: hidden;
			transition: max-height 0.2s ease-out;
			background-color: #ffffea;
			border-radius: 5px;
			border-top-left-radius: 0;
			border-top-right-radius: 0;
			box-shadow: 0.5px 0.5px 1.6px rgba(0, 0, 0, 0.022),
				1.1px 1.1px 2.4px rgba(0, 0, 0, 0.031),
				2.3px 2.4px 4.5px rgba(0, 0, 0, 0.039),
				5.7px 5.9px 21.3px rgba(0, 0, 0, 0.048),
				36px 37px 80px rgba(0, 0, 0, 0.07);

			:last-child {
				border: none;
			}
		}
	}
</style>
