export type ApiResponse = {
    code: number;
    status: string;
    message: string;
    data: any;
};

export const login = async (email: string, password: string, type: string): Promise<ApiResponse> => {
    const res = await fetch(`/api/auth/${type}_login`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password })
    });

    const data: ApiResponse = await res.json();
    return data;
}
