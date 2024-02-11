import type { ApiResponse } from "./auth";
export async function createHairdresser(saloonID: string, FirstName: string, Speciality: string): Promise<ApiResponse> {
    return fetch(`/api/hairdressers/create`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ saloonID, FirstName, Speciality })
    }).then(res => res.json());
}