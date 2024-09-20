import {staticServerAddr} from "@/config";
import axios from "axios";

interface NginxDirectory {
	name: string,
	type: string,
	mtime: string,
	size: number,
}

export async function listNginxDir(path: string): Promise<string[]> {
	const url = staticServerAddr.value + path
	const response = await axios.get<NginxDirectory[]>(url)
	const json = response.data
	return json.map(item => item.name)
}

export async function fetchNginxFile<T>(path: string): Promise<T> {
	const url = staticServerAddr.value + path
	const response = await axios.get<T>(url)
	return response.data
}
