<template>
	<TransitionRoot
		:show="props.isOpen"
		as="template"
		:enter="$style.enter"
		:enter-from="$style.enterFrom"
		:enter-to="$style.enterTo"
		:leave="$style.leave"
		:leave-from="$style.leaveFrom"
		:leave-to="$style.leaveTo"
		:entered="$style.entered"
	>
		<Dialog @close="emitClose" :class="$style.modal">
			<!-- The backdrop, rendered as a fixed sibling to the panel container -->
			<div :class="$style.backdrop" aria-hidden="true"></div>

			<!-- Full-screen container to center the panel -->
			<div :class="$style.content">
				<close-icon :class="$style.closeIcon" @click="emitClose"></close-icon>
				<!-- The actual dialog panel -->
				<DialogPanel :class="$style.panel">
					<DialogTitle><slot name="title"></slot></DialogTitle>
					<slot name="content"></slot>
				</DialogPanel>
			</div>
		</Dialog>
	</TransitionRoot>
</template>

<script setup>
	import {
		Dialog,
		DialogPanel,
		DialogTitle,
		TransitionRoot
	} from '@headlessui/vue';

	import CloseIcon from '../../assets/x.svg';

	const emits = defineEmits(['close']);

	const props = defineProps({
		isOpen: Boolean
	});

	const emitClose = () => {
		emits('close');
	};
</script>

<style lang="scss" module>
	.enter {
		transition: opacity 0.2s ease-out;
	}

	.enterFrom,
	.leaveTo {
		opacity: 0 !important;
	}

	.enterTo,
	.leaveFrom {
		opacity: 1 !important;
	}

	.leave {
		transition: opacity 0.1s linear;
	}

	.entered {
		opacity: 1 !important;
	}
	.modal {
		position: fixed;
		z-index: 50;
		top: 0;
		right: 0;
		bottom: 0;
		left: 0;
		.backdrop {
			position: fixed;
			top: 0;
			right: 0;
			bottom: 0;
			left: 0;
			background-color: rgba(0, 0, 0, 0.5);
		}

		.content {
			position: relative;
			left: 50%;
			top: 50%;
			transform: translate(-50%, -50%);
			border-radius: 0.25rem;
			width: 95%;
			line-height: 1.5rem;

			.closeIcon {
				position: absolute;
				top: 20px;
				right: 20px;
				z-index: 500;
			}

			.panel {
				width: 100%;
				position: relative;
				background-color: #ffffff;
				border-radius: 0.25rem;
				padding: 2rem 3rem;
				background-color: #fef7f7;
				max-height: 95vh;
				overflow: scroll;
			}
		}
	}
</style>
