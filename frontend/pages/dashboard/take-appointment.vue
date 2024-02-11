<script setup lang="ts">
import { ref, onMounted } from 'vue';

interface HairSaloon {
    uid: string;
    name: string;
    adress: string;
    email: string;
    phone: string;
    openingTime: string;
    closingTime: string;
}
interface Hairdresser {
    saloonID: string;
    uid: string;
    firstName: string;
    speciality: string;
}

const saloons = ref<HairSaloon[]>([]);
const selectedSaloon = ref('');
const isChoosingSaloon = ref(true);
const saloonOpeningTime = ref('');
const saloonClosingTime = ref('');
const hairdressers = ref<Hairdresser[]>([]);
const choosenHairdresser = ref('');
const isChoosingHairdresser = ref(false);
const appointmentDate = ref('');
const appointmentHour = ref('');
const isChoosingHour = ref(false);

const getAllSaloons = () => {
    fetch('/api/saloons/all')
        .then(res => res.json())
        .then(data => {
            saloons.value = data.data;
        });
};

const getHairdressers = (saloonId: string) => {
    fetch(`/api/saloons/${saloonId}/hairdressers`)
        .then(res => res.json())
        .then(data => {
            hairdressers.value = data.data;
        });
};

onMounted(() => {
    getAllSaloons();
});

const handleChooseSaloon = (saloonId: string) => {
    getHairdressers(saloonId);
    saloonOpeningTime.value = saloons.value.find(saloon => saloon.uid === saloonId)?.openingTime ?? '';
    saloonClosingTime.value = saloons.value.find(saloon => saloon.uid === saloonId)?.closingTime ?? '';
    selectedSaloon.value = saloonId;
    isChoosingSaloon.value = false;
    isChoosingHairdresser.value = true;
};

const handleChooseHairdresser = (hairdresserId: string) => {
    isChoosingHairdresser.value = false;
    choosenHairdresser.value = hairdresserId;
    isChoosingHour.value = true;
};

const handleChooseHour = () => {
    console.log('Appointment date', appointmentDate.value);
    console.log('Appointment hour', appointmentHour.value);
    // Here we can book the appointment
    checkIfAppointmentIsPossible();
};

const debug = () => {
    console.log("saloon", saloons.value);
    console.log("hairdressers", hairdressers.value);
    console.log("OpeningTime", saloonOpeningTime.value);
    console.log("saloonClosingTime", saloonClosingTime.value);
    console.log('appointment date', appointmentDate.value);
    console.log('appointment hour', appointmentHour.value);
};

async function checkIfAppointmentIsPossible() {
    const hairdresserUID = choosenHairdresser.value;
    const date = appointmentDate.value;
    const hour = appointmentHour.value;

    fetch('/api/hairdressers/isAvailable', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            hairdresserID: hairdresserUID,
            startHour: hour,
            appointmentDate: date
        })
    })
        .then(res => res.json())
        .then(data => {
            if (data.code === 200) {
                bookAppointment();
            }
            else {
                console.log('Appointment is not possible');
            }
        })
};

function bookAppointment() {
    const hairdresserUID = choosenHairdresser.value;
    const date = appointmentDate.value;
    const hour = appointmentHour.value;

    console.log({
        clientId: useAuthStore().uid,
        saloonId: selectedSaloon.value,
        hairdresserID: hairdresserUID,
        startHour: hour,
        appointmentDate: date
    })

    fetch('/api/appointments/create', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            clientId: useAuthStore().uid,
            saloonId: selectedSaloon.value,
            hairdresserID: hairdresserUID,
            startHour: hour,
            appointmentDate: date
        })
    })
        .then(res => res.json())
        .then(data => {
            console.log(data);
        });
};

onMounted(() => {
    const userType = useAuthStore().userType;
    if (userType !== 'client') {
        navigateTo('/dashboard');
    }
});
</script>

<template>
    <main>
        <h1>Take an appointment</h1>
        <button @click="debug">Debug</button>

        <div v-if="isChoosingSaloon">
            <div v-for="saloon in saloons" :key="saloon.uid">
                <h2>{{ saloon.name }}</h2>
                <p>{{ saloon.adress }}</p>
                <p>{{ saloon.email }}</p>
                <p>{{ saloon.phone }}</p>
                <p>{{ saloon.openingTime }} - {{ saloon.closingTime }}</p>
                <button @click="handleChooseSaloon(saloon.uid)">Choose</button>
            </div>
        </div>
        <span v-else v-show="isChoosingSaloon">Loading saloons...</span>
        <div v-if="hairdressers.length > 0" v-show="isChoosingHairdresser">
            <div v-for="hairdresser in hairdressers" :key="hairdresser.uid">
                <h2>{{ hairdresser.firstName }}</h2>
                <p>{{ hairdresser.speciality }}</p>
                <button @click="handleChooseHairdresser(hairdresser.uid)">Choose</button>
            </div>
        </div>

        <div v-show="isChoosingHour">
            <h2>
                At what time do you want to take an appointment?
            </h2>
            <input type="date" v-model="appointmentDate" required min="{{ new Date().toISOString().split('T')[0] }}">
            <input type="time" v-model="appointmentHour" required>
            <button @click="handleChooseHour">Choose</button>
        </div>
    </main>
</template>