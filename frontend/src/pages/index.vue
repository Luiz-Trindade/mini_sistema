<template>
    <q-layout view="lHh Lpr lFf">
        <!-- Oculta o header na página de login -->
        <q-header v-if="!isLoginPage" elevated>
            <q-toolbar>
                <q-btn flat dense round icon="menu" aria-label="Menu" @click="toggleLeftDrawer" />

                <q-toolbar-title> Mini Sistema </q-toolbar-title>

                <div>Quasar v{{ $q.version }}</div>
            </q-toolbar>
        </q-header>

        <!-- Oculta o menu lateral (drawer) na página de login -->
        <q-drawer v-if="!isLoginPage" v-model="leftDrawerOpen" show-if-above bordered>
            <q-list class="fit column justify-between">
                <!-- Seção Superior: Links principais -->
                <div>
                    <q-item-label header> Links Essenciais </q-item-label>
                    <EssentialLink v-for="link in linksList" :key="link.label" v-bind="link" />
                </div>

                <!-- Seção Inferior: Botão de Logout -->
                <div class="q-pa-sm">
                    <q-separator class="q-mb-sm" />
                    <q-item clickable v-ripple class="text-negative rounded-borders" @click="handleLogout">
                        <q-item-section avatar>
                            <q-icon name="logout" color="negative" />
                        </q-item-section>
                        <q-item-section class="text-weight-bold">
                            Sair / Logout
                        </q-item-section>
                    </q-item>
                </div>
            </q-list>
        </q-drawer>

        <q-page-container>
            <router-view />
        </q-page-container>
    </q-layout>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useQuasar } from 'quasar'
import { api } from '@/boot/axios'
import EssentialLink from '@/components/EssentialLink.vue'

const $q = useQuasar()
const route = useRoute()
const router = useRouter()

// Reatividade para verificar se está na rota de login
const isLoginPage = computed(() => route.path === '/login')

const leftDrawerOpen = ref(false)

function toggleLeftDrawer() {
    leftDrawerOpen.value = !leftDrawerOpen.value
}

// Função de Logout com diálogo de confirmação
function handleLogout() {
    $q.dialog({
        title: 'Confirmar saída',
        message: 'Tem a certeza de que deseja encerrar a sua sessão?',
        cancel: {
            label: 'Cancelar',
            flat: true,
            color: 'grey-8'
        },
        ok: {
            label: 'Sair',
            color: 'negative'
        },
        persistent: true
    }).onOk(async () => {
        try {
            // Envia requisição para invalidar o token no Go
            await api.post('/auth/logout')
        } catch (error) {
            console.error('Erro ao registrar logout no servidor:', error)
        } finally {
            // Remove o token do navegador independente de falhas de rede
            localStorage.removeItem('token')

            $q.notify({
                type: 'info',
                message: 'Sessão encerrada com sucesso',
                position: 'top'
            })

            // Redireciona para o login
            router.push('/login')
        }
    })
}

const linksList = [
    {
        label: "Docs",
        caption: "quasar.dev",
        icon: "school",
        link: "https://quasar.dev"
    },
    {
        label: "GitHub",
        caption: "github.com/quasarframework",
        icon: "code",
        link: "https://github.com/quasarframework"
    },
    {
        label: "Discord Chat Channel",
        caption: "chat.quasar.dev",
        icon: "chat",
        link: "https://chat.quasar.dev"
    },
    {
        label: "Forum",
        caption: "forum.quasar.dev",
        icon: "record_voice_over",
        link: "https://forum.quasar.dev"
    },
    {
        label: "Twitter",
        caption: "@quasarframework",
        icon: "rss_feed",
        link: "https://twitter.quasar.dev"
    },
    {
        label: "Facebook",
        caption: "@QuasarFramework",
        icon: "public",
        link: "https://facebook.quasar.dev"
    },
    {
        label: "Quasar Awesome",
        caption: "Community Quasar projects",
        icon: "favorite",
        link: "https://awesome.quasar.dev"
    }
]
</script>

<style scoped></style>