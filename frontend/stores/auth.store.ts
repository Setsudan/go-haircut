import { defineStore } from 'pinia'

export const useAuthStore = defineStore('auth', {
    state: () => ({
        token: '',
        uid: '',
        userType: '',
    }),

    getters: {
        isAuthenticated(): boolean {
            return this.token !== ''
        },
    },
})