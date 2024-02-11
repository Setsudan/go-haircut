<template>
  <div>
    <h1>Réservez votre rendez-vous chez {{ salonDetails.name }}</h1>
    <div v-if="hours.length > 0">
      <div v-for="hour in hours" :key="hour">
        <button @click="bookAppointment(hour)">{{ hour }}:00</button>
      </div>
    </div>
    <div v-else>Chargement des heures disponibles...</div>
  </div>
</template>

<script>
import { createAppointment } from "~/utils/appointment";

export default {
  async asyncData({ params, $http }) {
    const salonId = params.salonId;
    let salonDetails = {};
    let hours = [];

    try {
      const response = await $http.$get(`/saloons/${salonId}`);
      if (response.status === "Success") {
        salonDetails = response.data;
        hours = generateHoursFromOpeningTimes(
          salonDetails.openingTime,
          salonDetails.closingTime
        );
      }
    } catch (error) {
      console.error("Error fetching salon details:", error);
    }

    return { salonId, salonDetails, hours };
  },
  methods: {
    generateHoursFromOpeningTimes(openingTime, closingTime) {
      const startHour = parseInt(openingTime.split(":")[0]);
      const endHour = parseInt(closingTime.split(":")[0]);
      const hours = [];
      for (let hour = startHour; hour < endHour; hour++) {
        hours.push(hour);
      }
      return hours;
    },
    async bookAppointment(hour) {
      const startHour =
        new Date().toISOString().split("T")[0] + `T${hour}:00:00Z`;
      const appointmentData = {
        saloonId: this.salonId,
        clientId: "Mettre l'ID du client svp je le mettais manuellement", // Mettre l'ID du client svp
        hairdresserId:
          "Mettre l'ID du coiffeur svp je le mettais manuellement aussi", // Mettre l'ID du coiffeur svp
        startHour: startHour,
      };
      const response = await createAppointment(appointmentData);
      alert(response.message);
    },
  },
};
</script>
