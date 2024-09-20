<script setup lang="ts">
import {listNginxDir} from '@/utils/nginx';
import { ref } from 'vue';

const albums = ref<string[]>([]);

(async () => {
	albums.value = await listNginxDir('/gallery/')
})()

</script>
<template>
	<div style="height: 100vh; padding-top: 56px;">
		<div style="overflow: hidden; height: 100%; position: relative;">
			<img src="https://static.axtl.cn:520/carousel/图片一.jpg" style="object-fit: cover; object-position: center; width: 100%; height: 100%;" alt="">
		</div>
	</div>
</template>
