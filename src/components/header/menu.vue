<template>
<nav class="menu" id="main-menu">
	<button class="menu-toggle" id="toggle-menu">
	toggle menu
	</button>
	<div class="menu-dropdown">
		<ul class="nav-menu">
			<li><a @click="navigateTo(navigation.home.ROUTE)">{{navigation.home.NAME}}</a></li>
			<li><a @click="navigateTo(navigation.dates.ROUTE)">{{navigation.dates.NAME}}</a></li>
			<li><a @click="navigateTo(navigation.about.ROUTE)">{{navigation.about.NAME}}</a></li>
			<li><a @click="navigateTo(navigation.contact.ROUTE)">{{navigation.contact.NAME}}</a></li>
		</ul>
	</div>
</nav>
</template>

<script setup>
import { navigation } from '../../const/strings';

const navigateTo = (route) => {
  switch (route){
    case navigation.home.ROUTE:
      window.history.pushState("Home", navigation.home.NAME, navigation.home.ROUTE);
      break;
    case navigation.about.ROUTE:
      console.log("is About")
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
}

const emit = defineEmits(['navigateTo'])

</script>

<menu>

</menu>