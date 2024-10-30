import axios from "axios";

export const apiClient = axios.create({
	baseURL: 'https://api.axtl.cn:520',
})

export interface Resp<T> {
	message: string;
	data: T;
}
