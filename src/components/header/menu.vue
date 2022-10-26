<template>
<nav :class="$style.menu" id="main-menu">
	<button :class="$style.menuToggle" id="toggle-menu">
		<menu-icon v-if="!isOpen" @click="open"/>
		<close-icon v-else @click="close" />
	</button>
	<div v-if="isOpen" :class="$style.menuDropdown">
		<div  ref="animatedDiv" v-if="!animationFinished" :class="$style.animatedCircle"></div>
		<ul v-else :class="$style.navMenu">
			<div :class="showMenuContent ?  [$style.navContent, $style.show] : [$style.navContent]">
				<li><a @click="navigateTo(navigation.home.ROUTE)">{{navigation.home.NAME}}</a></li>
				<hr/>
				<li><a @click="navigateTo(navigation.dates.ROUTE)">{{navigation.dates.NAME}}</a></li>
				<hr/>
				<li><a @click="navigateTo(navigation.about.ROUTE)">{{navigation.about.NAME}}</a></li>
				<hr/>
				<li><a @click="navigateTo(navigation.contact.ROUTE)">{{navigation.contact.NAME}}</a></li>
			</div>
			</ul>
	</div>
</nav>
</template>

<script setup>
import { ref } from 'vue';
import MenuIcon from '../../assets/menu.svg';
import CloseIcon from '../../assets/x.svg';
import { navigation } from '../../const/strings';

let isOpen = ref(false);

let animationFinished = ref(false);

let showMenuContent = ref (false);

let animatedDiv = ref(null)

const close = () => {
	isOpen.value = false;
	animationFinished.value = false;
	showMenuContent.value = false;
}

const open = () => {
	isOpen.value = true;
	if (isOpen.value){
		setTimeout( () => {
			animationFinished.value = true;
			setTimeout(() => {
				showMenuContent.value = true;
			}, 50);
		},150)
	}
}

const navigateTo = (route) => {
  switch (route){
    case navigation.home.ROUTE:
      window.history.pushState("Home", navigation.home.NAME, navigation.home.ROUTE);
      break;
    case navigation.about.ROUTE:
      window.history.pushState("About", navigation.about.NAME, navigation.about.ROUTE);
      break;
    case navigation.dates.ROUTE:
      window.history.pushState("Dates", navigation.dates.NAME, navigation.dates.ROUTE);
      break;
    case navigation.contact.ROUTE:
      window.history.pushState("Contact", navigation.contact.NAME, navigation.contact.ROUTE);
      break;
    default:
			break;
  }
	// fire the event hashchange to trigger a route change in App.vue as pushState
	// does not trigger this event
	window.dispatchEvent(new HashChangeEvent('hashchange'));
	close();
}

</script>

<style lang="scss" module>
.menu{
	position: relative;
	margin-right: 2rem;
	.menuToggle{
		position: absolute;
		border: 0;
		background-color: transparent;
		z-index: 500;
	}

	.menuDropdown{
		position: fixed;
		left: 0;
		right: 0;
		top: 0;
		bottom: 0;


		.animatedCircle{
			position: fixed;
			z-index:500;
			background: #F6D7CD;

			animation: fadeIn 150ms 1 ease-out;
			animation-fill-mode: forwards;
			transform: translate(50%,-50%);
			border-radius: 50%;
			height: 3rem;
			width: 3rem;
			overflow: hidden;

			@keyframes fadeIn {

				from{
					transform: translate(50%,-50%);
					border-radius: 50%;
					top: 0;
					right: 0;
					border: 3rem solid #F6D7CD;
					height:	3rem;
					width: 3rem;
					border-radius: 50%;
				}

				to {
					transform: translate(50%,-50%) !important;
					border-radius: 50%;
					top: 0;
					right: 0;
					border: 72rem solid #F6D7CD;
					height: 72rem;
					width:	72rem;
					border-radius: 50%;
				}
		}
		}

		.navMenu{
			display: flex;
			flex-direction: column;
			padding: 6rem 2rem 2rem 2rem;
			margin-top: 0;
			margin-bottom: 0;
			position: fixed;
			z-index:500;
			background: #F6D7CD;
			top: 0;
			right: 0;
			border: 1px solid #F6D7CD;
			width: 100vw;
			height: 100vh;

			.navContent{
				opacity: 0;
				transition: opacity 150ms ease-out;
				hr{
					border: 0;
					border-top: 2px solid #bcb9b9;
					margin: 0 auto;
					width: 80%;
				}

				li{
					list-style-type: none;
					margin-top: 1.5rem;
					margin-bottom: 1.5rem;
					font-size: 1.5rem;
					font-weight: 600;
				}

				a:link {
					color: #1476b8;
					font-weight: bold;
					text-decoration: none;
				}
				a:visited {
					color: #1430b8;
				}
				a:hover {
					text-decoration: underline;
				}
				a:active {
					color: #b81414;
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
}

</style>