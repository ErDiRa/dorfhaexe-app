<template>
	<nav :class="$style.menu" id="main-menu">
		<button :class="$style.menuToggle" id="toggle-menu">
			<menu-icon v-if="!isOpen" @click="open" />
			<close-icon v-else @click="close" />
		</button>
		<div v-if="isOpen" :class="$style.menuDropdown">
			<div
				ref="animatedDiv"
				v-if="!animationFinished"
				:class="$style.animatedCircle"
			></div>
			<ul v-else :class="$style.navMenu">
				<div
					:class="
						showMenuContent
							? [$style.navContent, $style.show]
							: [$style.navContent]
					"
				>
					<li>
						<a
							@click="navigateTo(navigation.home.ROUTE)"
							:class="
								currentView && currentView.path === navigation.home.ROUTE
									? $style.active
									: ''
							"
							>{{ navigation.home.NAME }}</a
						>
					</li>
					<hr />
					<li>
						<a
							@click="navigateTo(navigation.dates.ROUTE)"
							:class="
								currentView && currentView.path === navigation.dates.ROUTE
									? $style.active
									: ''
							"
							>{{ navigation.dates.NAME }}</a
						>
					</li>
					<hr />
					<li>
						<a
							@click="navigateTo(navigation.about.ROUTE)"
							:class="
								currentView && currentView.path === navigation.about.ROUTE
									? $style.active
									: ''
							"
							>{{ navigation.about.NAME }}</a
						>
					</li>
					<hr />
					<li>
						<a
							@click="navigateTo(navigation.contact.ROUTE)"
							:class="
								currentView && currentView.path === navigation.contact.ROUTE
									? $style.active
									: ''
							"
							>{{ navigation.contact.NAME }}</a
						>
					</li>
				</div>
			</ul>
		</div>
	</nav>
</template>

<script setup>
	import { inject, ref, watch } from 'vue';
	import MenuIcon from '../../assets/menu.svg';
	import CloseIcon from '../../assets/x.svg';
	import { navigation } from '../../const/strings';

	let currentView = inject('currentView');

	watch(
		() => currentView.value,
		(newVal, oldVal) => {
			console.log(newVal, oldVal);
		}
	);

	let isOpen = ref(false);

	let animationFinished = ref(false);

	let showMenuContent = ref(false);

	let animatedDiv = ref(null);

	const close = () => {
		isOpen.value = false;
		animationFinished.value = false;
		showMenuContent.value = false;
		document.documentElement.style.overflow = 'auto';
	};

	const open = () => {
		isOpen.value = true;
		document.documentElement.style.overflow = 'hidden';
		if (isOpen.value) {
			setTimeout(() => {
				// stop the animation to display full size div (menu container)
				animationFinished.value = true;
				// trigger displaying of menu content
				setTimeout(() => {
					showMenuContent.value = true;
				}, 30);
			}, 150);
		}
	};

	const navigateTo = (route) => {
		switch (route) {
			case navigation.home.ROUTE:
				window.history.pushState(
					'Home',
					navigation.home.NAME,
					navigation.home.ROUTE
				);
				break;
			case navigation.about.ROUTE:
				window.history.pushState(
					'About',
					navigation.about.NAME,
					navigation.about.ROUTE
				);
				break;
			case navigation.dates.ROUTE:
				window.history.pushState(
					'Dates',
					navigation.dates.NAME,
					navigation.dates.ROUTE
				);
				break;
			case navigation.contact.ROUTE:
				window.history.pushState(
					'Contact',
					navigation.contact.NAME,
					navigation.contact.ROUTE
				);
				break;
			default:
				break;
		}
		// fire the event hashchange to trigger a route change in App.vue as pushState
		// does not trigger this event
		window.dispatchEvent(new HashChangeEvent('hashchange'));
		close();
	};
</script>

<style lang="scss" module>
	.menu {
		.menuToggle {
			position: relative;
			border: 0;
			background-color: transparent;
			z-index: 502;
		}

		.menuDropdown {
			position: fixed;
			z-index: 500;
			left: 0;
			right: 0;
			top: 0;
			bottom: 0;

			.animatedCircle {
				position: fixed;
				z-index: 500;
				background: #f6d7cd;

				animation: fadeIn 150ms 1 ease-in;
				animation-fill-mode: forwards;
				transform: translate(50%, -50%);
				border-radius: 50%;
				height: 3rem;
				width: 3rem;
				overflow: hidden;

				@keyframes fadeIn {
					from {
						transform: translate(50%, -50%);
						border-radius: 50%;
						top: 0;
						right: 0;
						height: 3rem;
						width: 3rem;
						border-radius: 50%;
					}

					to {
						transform: translate(25%, -25%);
						border-radius: 50%;
						top: 0;
						right: 0;
						height: 400vh;
						width: 300vh;
						border-radius: 50%;
					}
				}
			}

			.navMenu {
				display: flex;
				flex-direction: column;
				padding: 6rem 2rem 2rem 2rem;
				margin-top: 0;
				margin-bottom: 0;
				position: fixed;
				z-index: 500;
				background: #f6d7cd;
				top: 0;
				right: 0;
				border: 1px solid #f6d7cd;
				width: 100vw;
				height: 100vh;

				.navContent {
					opacity: 0;
					transition: opacity 150ms ease-out;
					hr {
						border: 0;
						border-top: 2px solid #bcb9b9;
						margin: 0 auto;
						width: 45%;
					}

					li {
						list-style-type: none;
						margin-top: 1.5rem;
						margin-bottom: 1.5rem;
						font-size: 1.5rem;
						font-weight: 600;
						cursor: pointer;
					}

					a:link {
						color: #1476b8;
						font-weight: bold;
						text-decoration: none;
					}
					a:hover {
						color: #a7c69d;
					}
				}
				.show {
					opacity: 1;
				}
			}
		}

		.hide {
			display: none;
		}

		.active {
			color: #75a366;
		}
	}
</style>
