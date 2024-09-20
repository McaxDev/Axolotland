import {ref} from "vue"
import {fetchNginxFile} from "./nginx"
import {LogError} from "./logger"

// 获取Wiki元数据
export const wikiMeta = ref<Record<string, string>>({})
fetchNginxFile<Record<string, string>>('/wiki/metadata.json').then(response => {
	wikiMeta.value = response
}).catch(LogError)
