<template>
    <main>
        <h1>Register</h1>
        <form @submit.prevent="register">
            <label for="email">Email:</label>
            <input type="email" id="email" v-model="email" required>
            <label for="password">Password:</label>
            <input type="password" id="password" v-model="password" required>
            <button type="submit">Register</button>
            <span v-if="error">{{ error }}</span>
        </form>
        <div class="auth_options">
            <nuxt-link to="/auth/register/saloon">
                <span>
                    I'm a hairdressing salon. Register
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

const email = ref('');
const password = ref('');
const error = ref('');
const register = () => {
    const client = {
        email: email.value,
        password: password.value
    };
    signup(client, "client").then((res) => {
        console.log(res);
        handleResponse(res);
    });
};

const handleResponse = (res: ApiResponse) => {
    console.log('Handling response...', res);
    if (res.code === 201) {
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