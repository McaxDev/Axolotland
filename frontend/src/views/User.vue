<script setup lang="ts">
import axios from 'axios';
import { type WikiMeta } from '@/App.vue';
import { ref } from 'vue';

async function getWikiContent(filename: string): Promise<string> {
	try {
		const url = `${import.meta.env.VITE_STATIC}/wiki/${filename}`
		return (await axios.get<string>(url)).data
	} catch (error) {
		console.log(error)
		return ''
	}
}

const test = ref<string>('');
(async () => { test.value = await getWikiContent('index.md') })()

const props = defineProps<{
	wikis: WikiMeta;
}>()
</script>
<template>
	<b-container style="background-color: cyan; padding-top: 56px; display: flex;" class="vh-100">
		<b-row>
			<b-col cols="3">
				<b-nav vertical>
					<b-nav-item href="aaa.html">a</b-nav-item>
					<b-nav-item>b</b-nav-item>
					<b-nav-item>c</b-nav-item>
					<b-nav-item>d</b-nav-item>
					a {{ props.wikis }} a
				</b-nav>
			</b-col>
			<b-col class="overflow-y-auto mh-100">
			</b-col>
		</b-row>
	</b-container>
</template>
