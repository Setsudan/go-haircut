<script setup lang="ts">
const authStore = useAuthStore();

const logout = () => {
    authStore.token = '';
    authStore.uid = '';
    authStore.userType = '';
    navigateTo('/auth/login');
};
definePageMeta({
    middleware: 'default'
})
</script>
<template>
    <main id="saloon" v-if="authStore.userType === 'saloon'">
        <h1>Dashboard</h1>
        <p>Welcome to your dashboard</p>
        <nuxt-link to="/dashboard/create-hairdresser">
            <button>Create Hairdresser</button>
        </nuxt-link>
    </main>
    <button @click="logout">Logout</button>
</template>