export type ApiResponse = {
  code: number;
  status: string;
  message: string;
  data: any;
};

type AppointmentData = {
  saloonId: string;
  clientId: string;
  hairdresserId: string;
  startHour: string;
};

export const createAppointment = async (
  appointmentData: AppointmentData
): Promise<ApiResponse> => {
  const res = await fetch(`/appointments/create`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(appointmentData),
  });

  const data: ApiResponse = await res.json();
  return data;
};
