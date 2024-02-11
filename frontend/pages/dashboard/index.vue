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
    <pre>
        {
            "authStore.token": "{{ authStore.token }}",
            "authStore.uid": "{{ authStore.uid }}",
            "authStore.userType": "{{ authStore.userType }}"
        }
    </pre>
    <button @click="logout">Logout</button>
</template>