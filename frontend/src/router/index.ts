import { 
	createRouter, 
	createWebHistory, 
	type RouteRecordRaw 
} from 'vue-router'

import MainPage from '@/views/MainPage.vue'
import Gallery from '@/views/Gallery.vue'
import Wiki from '@/views/Wiki.vue'
import Dash from '@/views/Dash.vue'
import User from '@/views/User.vue'
import Login from '@/views/Login.vue'

export const routes: readonly RouteRecordRaw[] = [
		{
			path: '/',
			name: 'mainpage',
			component: MainPage
		},
		{
			path: '/gallery/:page?',
			name: 'gallery',
			component: Gallery
		},
		{
			path: '/wiki/:page?',
			name: 'wiki',
			component: Wiki
		},
		{
			path: '/dash',
			name: 'dash',
			component: Dash
		},
		{
			path: '/user',
			name: 'user',
			component: User,
			meta: { requiresAuth: true }
		},
		{
			path: '/login',
			name: '/login',
			component: Login
		}
	]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: routes
})

router.beforeEach((to, _, next) => {
	if (to.matched.some(record => record.meta.requiresAuth)) {
		if (!localStorage.getItem('authToken')) {
			next({ path: '/login', query: {redirect: to.fullPath}})
		} else {
			next()
		}
	} else {
		next()
	}
})

export default router
