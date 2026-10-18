<template>
    <q-page class="bg-grey-2 flex flex-center padding">
        <q-card style="width: 100%; max-width: 400px" class="q-pa-md shadow-3">
            <q-card-section class="text-center">
                <div class="text-h5 text-bold text-primary">Mini Sistema</div>
                <div class="text-subtitle2 text-grey-7">Informe suas credenciais para entrar</div>
            </q-card-section>

            <q-card-section>
                <q-form @submit="onSubmit" class="q-gutter-md">
                    <q-input v-model="email" type="email" label="E-mail" outlined dense lazy-rules :rules="[
                        val => (val && val.length > 0) || 'Por favor, digite seu e-mail',
                        val => /.+@.+\..+/.test(val) || 'Por favor, digite um e-mail válido'
                    ]">
                        <template #prepend>
                            <q-icon name="mail" />
                        </template>
                    </q-input>

                    <q-input v-model="password" :type="isPasswordVisible ? 'text' : 'password'" label="Senha" outlined
                        dense lazy-rules :rules="[val => (val && val.length > 0) || 'Por favor, digite sua senha']">
                        <template #prepend>
                            <q-icon name="lock" />
                        </template>
                        <template #append>
                            <q-icon :name="isPasswordVisible ? 'visibility_off' : 'visibility'" class="cursor-pointer"
                                @click="isPasswordVisible = !isPasswordVisible" />
                        </template>
                    </q-input>

                    <div>
                        <q-btn icon="login" label="Entrar" type="submit" color="primary" class="full-width" size="lg"
                            :loading="loading" />
                    </div>
                </q-form>
            </q-card-section>
        </q-card>
    </q-page>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useQuasar } from 'quasar'
import { api } from '@/boot/axios'

const $q = useQuasar()
const router = useRouter()

const email = ref('')
const password = ref('')
const isPasswordVisible = ref(false)
const loading = ref(false)

async function onSubmit() {
    loading.value = true

    try {
        const response = await api.post('/auth/login', {
            email: email.value,
            senha: password.value
        })

        if (response.data.token) {
            localStorage.setItem('token', response.data.token)
        }

        $q.notify({
            type: 'positive',
            message: 'Login realizado com sucesso!',
            position: 'top'
        })

        router.push('/')
    } catch (error) {
        const errorMsg = error.response?.data?.error || 'Erro ao realizar login'

        $q.notify({
            type: 'negative',
            message: errorMsg,
            position: 'top'
        })
    } finally {
        loading.value = false
    }
}
</script>

<style scoped></style>