<template>
    <main>
        <h1>Create Hairdressing Saloon</h1>
        <form @submit.prevent="registerSaloon">
            <label for="name">Name:</label>
            <input type="text" id="name" v-model="saloon.name" required>

            <label for="address">Address:</label>
            <input type="text" id="address" v-model="saloon.address" required>

            <label for="email">Email:</label>
            <input type="email" id="email" v-model="saloon.email" required>
            <span v-if="!isEmailValid && saloon.email.length > 3" style="color: red;">Invalid email format</span>

            <label for="phone">Phone:</label>
            <input type="text" id="phone" v-model="saloon.phone" required>
            <span v-if="!isPhoneValid && saloon.phone.length > 3" style="color: red;">Invalid phone number format</span>

            <label for="openingTime">Opening Time:</label>
            <input type="text" id="openingTime" v-model="saloon.openingTime" required placeholder="eg: 09:00">
            <span v-if="!isTimeValid(saloon.openingTime) && saloon.openingTime.length > 3" style="color: red;">Invalid time
                format</span>

            <label for="closingTime">Closing Time:</label>
            <input type="text" id="closingTime" v-model="saloon.closingTime" required placeholder="eg: 18:00">
            <span v-if="!isTimeValid(saloon.closingTime) && saloon.closingTime.length > 3" style="color: red;">Invalid time
                format</span>
            <span v-if="isTimeConflict" style="color: red;">Closing time must be greater than opening time</span>

            <label for="password">Password:</label>
            <input type="password" id="password" v-model="saloon.password" required>
            <span v-if="!isPasswordValid && saloon.password.length > 3" style="color: red;">Invalid password format</span>


            <span v-if="error" style="color: red;">{{ error }}</span>
            <button type="submit">Register</button>
        </form>
        <div class="auth_options">
            <nuxt-link to="/auth/register/client">
                <span>
                    I'm a client salon. Register
                </span>
            </nuxt-link>
            <nuxt-link to="/auth/login">
                <span>
                    I have an account. Login
                </span>
            </nuxt-link>
        </div>
    </main>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const passwordRegex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d]{8,}$/;
const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
const phoneRegex = /^\d{10}$/;
const timeRegex = /^([01]\d|2[0-3]):([0-5]\d)$/;

const saloon = ref({
    name: '',
    address: '',
    email: '',
    phone: '',
    openingTime: '',
    closingTime: '',
    password: '',
});

const isEmailValid = ref(true);
const isPhoneValid = ref(true);
const isTimeConflict = ref(false);
const isPasswordValid = ref(true);
const error = ref('');

const reformatTime = (time: string) => {
    const [hours, minutes] = time.split(':');
    return `${hours.padStart(2, '0')}:${minutes.padStart(2, '0')}:00`;
};

const isTimeValid = (time: string) => {
    return timeRegex.test(time);
};

const registerSaloon = () => {
    isEmailValid.value = emailRegex.test(saloon.value.email);
    isPhoneValid.value = phoneRegex.test(saloon.value.phone);
    isPasswordValid.value = passwordRegex.test(saloon.value.password);

    if (isTimeValid(saloon.value.openingTime) && isTimeValid(saloon.value.closingTime)) {
        const openingTime = reformatTime(saloon.value.openingTime);
        const closingTime = reformatTime(saloon.value.closingTime);

        if (openingTime > closingTime) {
            isTimeConflict.value = true;
            return;
        }

        saloon.value.openingTime = openingTime;
        saloon.value.closingTime = closingTime;
    } else {
        isTimeConflict.value = false;
        return;
    }

    console.log(saloon.value)

    signup(saloon.value, "saloon").then((res) => {
        console.log(res);
        handleResponse(res);
    });
};

const handleResponse = (res: ApiResponse) => {
    console.log('Handling response...', res);
    if (res.code === 200) {
        console.log('Registration successful');
        navigateTo('/auth/login');
    }
    else if (res.code === 400) {
        // wrong credentials
        console.log('Registration failed');
        error.value = res.message;
    }
};
</script>

<style scoped lang="scss">
main {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 1rem;
    height: 100vh;
    justify-content: center;
    text-align: center;
}

h1 {
    color: #333;
    font-size: 24px;
    margin-bottom: 1rem;
}

form {
    display: flex;
    flex-direction: column;
    gap: 1rem;
}

input {
    padding: 0.5rem;
    border: 1px solid #ccc;
    border-radius: 4px;
}

button {
    padding: 0.5rem;
    background-color: #333;
    color: #fff;
    border: none;
    cursor: pointer;
    border-radius: 4px;
}

button:hover {
    background-color: #222;
}

a {
    color: #333;
    text-decoration: none;
}

a:hover {
    color: #222;
}

.auth_options {
    display: flex;
    flex-direction: column;
    gap: 1rem;
    margin-top: 1rem;
}
</style>