import { defineRouter } from '#q-app'
import { routes, handleHotUpdate } from 'vue-router/auto-routes'
import {
	createMemoryHistory,
	createRouter,
	createWebHashHistory,
	createWebHistory,
} from 'vue-router'

export default defineRouter((/* { store, ssrContext } */) => {
	const createHistory = import.meta.env.QUASAR_SERVER
		? createMemoryHistory
		: import.meta.env.QUASAR_VUE_ROUTER_MODE === 'history'
			? createWebHistory
			: createWebHashHistory

	const Router = createRouter({
		scrollBehavior: () => ({ left: 0, top: 0 }),
		routes,

		// Leave this as is and make changes in quasar.conf.js instead!
		// quasar.conf.js -> build -> vueRouterMode
		// quasar.conf.js -> build -> publicPath
		history: createHistory(import.meta.env.QUASAR_VUE_ROUTER_BASE),
	})

	// Guard de Navegação Global para proteção de rotas
	Router.beforeEach((to, from, next) => {
		const token = localStorage.getItem('token')

		// 1. Se tentar acessar qualquer rota diferente de '/login' sem estar autenticado
		if (to.path !== '/login' && !token) {
			return next('/login')
		}

		// 2. Se já tiver token e tentar acessar a tela de '/login' novamente
		if (to.path === '/login' && token) {
			return next('/')
		}

		// 3. Permite a navegação normal
		next()
	})

	// enable HMR for it
	if (import.meta.hot) {
		handleHotUpdate(Router)
	}

	return Router
})
