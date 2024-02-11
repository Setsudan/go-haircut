<template>
    <main>
        <h1>Login</h1>
        <form @submit.prevent="tryLogin">
            <label for="email">Email:</label>
            <input type="email" id="email" v-model="email" required>
            <label for="password">Password:</label>
            <input type="password" id="password" v-model="password" required>
            <div>
                <h3>
                    Login as a:
                </h3>
                <label for="client">Client</label>
                <input type="radio" id="client" value="client" v-model="userType">
                <label for="salon">Saloon</label>
                <input type="radio" id="salon" value="saloon" v-model="userType">
            </div>
            <span v-if="error">{{ error }}</span>
            <button type="submit">Login</button>
        </form>
        <div>
            <nuxt-link to="/auth/register/client">
                <span>
                    I don't have an account. Register
                </span>
            </nuxt-link>
        </div>
    </main>
</template>

<script setup lang="ts">
import { ref } from 'vue';
const authStore = useAuthStore();

const email = ref('');
const password = ref('');
const userType = ref('client');
const error = ref('');

const tryLogin = async () => {
    console.log('Logging in...');
    login(email.value, password.value, userType.value).then((res) => {
        console.log(res);
        handleResponse(res);
    });
};

const handleResponse = (res: ApiResponse) => {
    console.log('Handling response...', res);
    if (res.code === 200) {
        console.log('Login successful');
        console.log('Token:', res.data.token);
        authStore.token = res.data.token;
        authStore.uid = res.data.uid;
        authStore.userType = res.data.userType;
        navigateTo('/dashboard')
    }
    else if (res.code === 401) {
        // wrong credentials
        error.value = res.message;
    }
    else if (res.code === 500) {
        // Display the error message
        error.value = res.message;
    }
    else {
        // error is not planned
        error.value = 'An error occured. Please try again later';
    }
}
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
</style>