import {staticServerAddr} from "@/config";

export function joinImgSrc(path: string): URL {
	return new URL(staticServerAddr.value + path)
}
