<template>
	<nav :class="$style.menu" id="main-menu">
		<button :class="$style.menuToggle" id="toggle-menu">
			<menu-icon v-if="!isOpen" @click="open" />
			<close-icon v-else @click="close" />
		</button>
		<Transition>
			<div v-if="isOpen" :class="$style.menuDropdown">
				<ul :class="$style.navMenu">
					<div :class="$style.navContent">
						<li>
							<a
								@click="navigateTo(navigation.home.ROUTE)"
								:class="
									currentView && currentView === navigation.home.ROUTE
										? $style.active
										: ''
								"
								>{{ navigation.home.NAME }}</a
							>
						</li>
						<hr />
						<li>
							<a
								@click="navigateTo(navigation.romo.ROUTE)"
								:class="
									currentView && currentView === navigation.romo.ROUTE
										? $style.active
										: ''
								"
								>{{ navigation.romo.NAME }}</a
							>
						</li>
						<hr />
						<li>
							<a
								@click="navigateTo(navigation.dates.ROUTE)"
								:class="
									currentView && currentView === navigation.dates.ROUTE
										? $style.active
										: ''
								"
								>{{ navigation.dates.NAME }}</a
							>
						</li>
						<hr />
						<li>
							<a
								@click="navigateTo(navigation.news.ROUTE)"
								:class="
									currentView && currentView === navigation.news.ROUTE
										? $style.active
										: ''
								"
								>{{ navigation.news.NAME }}</a
							>
						</li>
						<hr />
						<li>
							<a
								@click="navigateTo(navigation.about.ROUTE)"
								:class="
									currentView && currentView === navigation.about.ROUTE
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
									currentView && currentView === navigation.contact.ROUTE
										? $style.active
										: ''
								"
								>{{ navigation.contact.NAME }}</a
							>
						</li>
						<hr />
						<li>
							<a
								@click="navigateTo(navigation.history.ROUTE)"
								:class="
									currentView && currentView === navigation.history.ROUTE
										? $style.active
										: ''
								"
								>{{ navigation.history.NAME }}</a
							>
						</li>
					</div>
				</ul>
			</div>
		</Transition>
	</nav>
</template>

<script setup>
	import { computed } from '@vue/reactivity';
	import { ref } from 'vue';
	import { useRoute, useRouter } from 'vue-router';
	import MenuIcon from '../../assets/menu.svg';
	import CloseIcon from '../../assets/x.svg';
	import { navigation } from '../../const/strings';

	const route = useRoute();

	const router = useRouter();

	const currentView = computed(() => {
		return route.path;
	});

	const close = () => {
		isOpen.value = false;
		document.documentElement.style.overflow = 'auto';
	};

	let isOpen = ref(false);

	const open = () => {
		isOpen.value = true;
		document.documentElement.style.overflow = 'hidden';
	};

	const navigateTo = (path) => {
		router.push(path);
		close();
	};
</script>

<style>
	.v-enter-active {
		transition: opacity 0.3s ease;
	}

	.v-enter-from,
	.v-leave-to {
		opacity: 0;
	}
</style>

<style lang="scss" module>
	.menu {
		.menuToggle {
			position: relative;
			border: 0;
			background-color: transparent;
			color: inherit;
			z-index: 502;
		}

		.menuDropdown {
			position: fixed;
			z-index: 500;
			left: 0;
			right: 0;
			top: 0;
			bottom: 0;

			.navMenu {
				display: flex;
				flex-direction: column;
				padding: 5rem 2rem 2rem 2rem;
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
				overflow: scroll;

				.navContent {
					opacity: 1;
					transition: opacity 150ms ease-out;
					hr {
						border: 0;
						border-top: 2px solid #bcb9b9;
						margin: 0 auto;
						width: 75%;
						max-width: 900px;
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
			}
		}
	}

	.hide {
		display: none;
	}

	.active {
		color: #75a366;
	}
</style>
