<script setup lang="ts">
import {useRoute} from 'vue-router';
import {computed, ref, watch} from 'vue';
import {fetchNginxFile} from '@/utils/nginx';
import {marked} from 'marked';
import {LogError} from '@/utils/logger';
import {wikiMeta} from '@/utils/wiki';

const route = useRoute()

const currentWiki = computed(() => route.path.split('/')[2].split('.')[0])

// 获取当前页的wiki内容
const wikiContent = ref<string>('')
loadWikiContent(currentWiki.value)

// 获取wiki内容的函数
async function loadWikiContent(name: string) {
	fetchNginxFile<string>(`/wiki/${name}.md`).then(async response => {
		wikiContent.value = await marked(response)
	}).catch(LogError)
}

// 监听当前页面的变化并重新加载
watch(
	currentWiki, (newName) => {
		loadWikiContent(newName)
})

</script>

<template>
	<b-container>
		<b-row class="wiki-row">

			<b-col cols="12" md="4" lg="3" xl="2" class="wiki-col-nav mt-5">
				<b-nav vertical class="wiki-nav p-3 rounded-4">
					<b-nav-item v-for="(value, key) in wikiMeta" :key="key" 
						:to="`${key}.html`" class="test"
					><div class="test">{{ value }}</div>
					</b-nav-item>
				</b-nav>
			</b-col>

			<b-col cols="12" md="8" lg="9" xl="10" class="wiki-col-content p-5">
				<div v-html="wikiContent" />
			</b-col>

		</b-row>
	</b-container>
</template>

<style scoped lang="scss">

.wiki-row {
	padding-top: 56px;
	height: 100vh;
	overflow-y: auto;

	.wiki-col-nav {
		max-height: 100%;
		overflow-y: auto;

		.wiki-nav {
			background-color: var(--subtle);
		}
	}
	
	.wiki-col-content {
		max-height: 100%;
		overflow-y: auto;
	}
}

.test {
	color: var(--deep);
}
</style>
