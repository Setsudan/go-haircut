import { useAuthStore } from "~/stores/auth.store"

export default defineNuxtRouteMiddleware((to, from) => {
    const authStore = useAuthStore()
    if (authStore.isAuthenticated) {
        return
    }
    return navigateTo('/auth/login')
})