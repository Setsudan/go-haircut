import { defineStore } from 'pinia'

export const useAuthStore = defineStore('auth', {
    state: () => ({
        jwt: null as string | null
    }),
    getters: {
        isAuthenticated: (state) => !!state.jwt
    },
    actions: {
        setJWT(jwt: string) {
            this.jwt = jwt
        },
        clearJWT() {
            this.jwt = null
        }
    }
})