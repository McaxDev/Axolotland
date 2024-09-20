<script setup lang="ts">
import { ref } from 'vue';
import {listNginxDir} from '@/utils/nginx';
import {LogError} from '@/utils/logger';
import {staticServerAddr} from '@/config';

// 获取轮播图列表
const CarouselItems = ref<string[]>([])
listNginxDir('/carousel/').then(response => {
	CarouselItems.value = response
}).catch(LogError)

</script>

<template>
	<b-carousel fade indicators>
		<b-carousel-slide v-for="item in CarouselItems" :key="item">

			<template #img>
				<div class="carousel-container">
					<img :src="`${staticServerAddr}/carousel/${item}`" alt="" class="carousel-img">
				</div>
			</template>

			<template #caption>
				<b-container>
					<b-row class="d-flex justify-content-between" style="height: 200px;">
						<b-col cols="12" sm="8" class="fs-1">
							Axolotland Gaming Club
						</b-col>
						<b-col cols="12" sm="4" class="d-flex flex-column justify-content-end gap-2 align-items-center">
							<b-button variant="info" to="/wiki">
								{{ $t('mainpage.carousel.btn-docs') }}
							</b-button>
							<b-button href="#">
								{{$t('mainpage.carousel.btn-join') }}
							</b-button>
						</b-col>
					</b-row>
				</b-container>
			</template>

			<template #text>
				<div class="carousel-text">{{item.split('.')[0]}}</div>
			</template>
		</b-carousel-slide>
	</b-carousel>
</template>

<style scoped>
.carousel-text {
	text-align: left;
}

.carousel-container {
	overflow: hidden;
	position: relative;
	height: 100vh;
}

.carousel-img {
	width: 100%;
	height: 100%;
	object-fit: cover;
	object-position: center;
	filter: brightness(70%);
}
</style>
