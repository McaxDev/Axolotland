<script setup lang="ts">
import AxLink from './components/AxLink.vue'
import NavBar from './components/NavBar.vue';

</script>

<template>

	
	<NavBar />

	<main class="ax-light-color min-vh-100">
		<router-view />
	</main>

  <footer class="p-4 ax-deep-color text-center">
    <AxLink href="#" class="text-center" target="_blank"> 备案号：豫ICP备2023025148号 </AxLink>
  </footer>
</template>

<style scoped>
.navbar {
	height: 56px;
}
</style>
