
<template>
  <header-view @navigate-to="navigate"/>
  <main>
    <component :is="mainComponent"/>
  </main>
  <footer-view/>
</template>

<script setup>
import { computed, markRaw } from 'vue';
import FooterView from './components/footer/footer.vue';
import HeaderView from './components/header/header.vue';
import { useRouter } from './router';

const router = useRouter()

let mainComponent  = computed(() => {
  if (! router.value) return;
  return markRaw(router.value.component);
})

// set the location and load the correct component
// use pushState to prevent a page reload as vue.js is loading the component
// which looks like a navigation
const navigate = (route) => {
}


</script>
<style lang="scss">
</style>
