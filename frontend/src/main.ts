import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import { createBootstrap } from 'bootstrap-vue-next'

import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap/dist/js/bootstrap.bundle.js'
import 'bootstrap-vue-next/dist/bootstrap-vue-next.css'
import 'bootstrap-icons/font/bootstrap-icons.css'
import './styles/global.css'
import './styles/themes.css'

import {createI18n} from 'vue-i18n'
import en from './i18n/en.json'
import zh from './i18n/zh.json'

createApp(App)
	.use(createPinia())
	.use(router)
	.use(createBootstrap())
	.use(createI18n({
		locale: navigator.language.split('-')[0],
		fallbackLocale: 'en',
		messages: {
			en, zh,
		}}))
	.mount('#app')
