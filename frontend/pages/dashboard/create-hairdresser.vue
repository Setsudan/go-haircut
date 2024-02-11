<template>
    <form @submit.prevent="submitForm">
        <label for="firstName">First Name:</label>
        <input type="text" id="firstName" v-model="firstName" required>

        <label for="speciality">Speciality:</label>
        <input type="text" id="speciality" v-model="speciality" required>

        <button type="submit">Submit</button>
    </form>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';

const saloonID = computed(() => {
    return useAuthStore().uid;
});
const firstName = ref('');
const speciality = ref('');

const submitForm = () => {
    createHairdresser(
        saloonID.value,
        firstName.value,
        speciality.value
    ).then((res) => {
        console.log(res);
        handleResponse(res);
    });
};

const handleResponse = (res: ApiResponse) => {
    console.log('Handling response...', res);
    if (res.code === 200) {
        console.log('Hairdresser created successfully');
        navigateTo('/dashboard');
    }
    else if (res.code === 400) {
        console.log('Hairdresser creation failed');
    }
};

onMounted(() => {
    const userType = useAuthStore().userType;
    if (userType !== 'saloon') {
        navigateTo('/dashboard');
    }
});
</script>
