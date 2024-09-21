<script setup lang="ts">
import {staticServerAddr} from '@/config';
import {LogError} from '@/utils/logger';
import {fetchNginxFile, listNginxDir} from '@/utils/nginx';
import {watch} from 'vue';
import { computed, ref } from 'vue';
import {useRoute, useRouter} from 'vue-router';

const route = useRoute()
const router = useRouter()
const currentAlbum = computed(() => route.path.split('/')[2])

// 定义相册描述和封面
interface AlbumMeta {
	path: string;
	cover: string;
	title: string;
	description: string;
}

// 获取相册列表和描述
const albumMeta = ref<AlbumMeta[]>([])
fetchNginxFile<AlbumMeta[]>('/gallery/metadata.json').then(response => {
	albumMeta.value = response
}).catch(LogError)

// 加载相册里的相片
const album = ref<string[]>([])
async function LoadAlbum() {
	listNginxDir(`/gallery/${currentAlbum.value}/`).then(response => {
		album.value = response
	}).catch(LogError)
}

// 检测路由变化并重新加载相片
watch(currentAlbum, () => {
	LoadAlbum()
})

</script>
<template>
	<div class="margin-nav py-3">
		<b-container class="p-4 rounded shadow">
			<b-row v-if="currentAlbum">
				<b-col v-for="item in album" :key="item"
					cols="6" md="4" lg="3"
				>
				<b-card
					:img-src="`${staticServerAddr}/gallery/${currentAlbum}/${item}`"
					:img-alt="item"
				>
					<b-card-text>{{item}}</b-card-text>
				</b-card>
				</b-col>
			</b-row>
			<b-row v-else>
				<b-col v-for="item in albumMeta" :key="item.path"
					cols="6" md="4" lg="3"
				>
					<b-card
						:overlay="false"
						:img-src="`${staticServerAddr}/gallery/cover/${item.cover}`"
						:img-alt="item.title"
						:title="item.title"
						text-variant="dark"
						@click="router.push('/gallery/'+item.path)"
					>
						<b-card-text>{{ item.description }}</b-card-text>
					</b-card>
				</b-col>
			</b-row>
		</b-container>
	</div>
</template>
