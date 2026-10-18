import { boot } from 'quasar/wrappers'
import axios from 'axios'

const api = axios.create({
	baseURL: '',
	headers: {
		'Content-Type': 'application/json',
	},
})

// Interceptor de Requisição: Adiciona o Token JWT no cabeçalho se existir no localStorage
api.interceptors.request.use(
	(config) => {
		const token = localStorage.getItem('token')
		if (token) {
			config.headers.Authorization = `Bearer ${token}`
		}
		return config
	},
	(error) => {
		return Promise.reject(error)
	},
)

export default boot(({ app }) => {
	app.config.globalProperties.$axios = axios
	app.config.globalProperties.$api = api
})

export { api }
