
<template>
  <header-view @navigate-to="navigate"/>
  <main>
    <component :is="mainComponent"></component>
  </main>
  <footer-view/>
</template>

<script setup>
import { computed } from 'vue';
import FooterView from './components/footer/footer.vue';
import HeaderView from './components/header/header.vue';
import About from './components/pages/about.vue';
import Contact from './components/pages/contact.vue';
import Dates from './components/pages/dates.vue';
import Home from './components/pages/home.vue';
import { navigation } from './const/strings';
import { useRouter } from './router';

const router = useRouter()

let mainComponent  = computed(() => {
  switch (router.currentView){
    case navigation.home.ROUTE:
      return Home;
    case navigation.about.ROUTE:
      console.log("is About")
      return About;
    case navigation.dates.ROUTE:
      return Dates;
    case navigation.contact.ROUTE:
      return Contact;
    default:
      return null;
  }
})

// set the location and load the correct component
// use pushState to prevent a page reload as vue.js is loading the component
// which looks like a navigation
const navigate = (route) => {
}


</script>
<style lang="scss">
</style>
