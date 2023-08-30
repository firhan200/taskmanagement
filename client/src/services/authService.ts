import axios from "axios"

const apiUrl: string = import.meta.env.VITE_API_URL ?? ""

export type LoginResponse = { 
    error: string | undefined
    token: string | undefined
}

export const login = async (emailAddress: string, password: string): Promise<LoginResponse> => {
    const res = await axios.post(`${apiUrl}auth/login`, {
        emailAddress: emailAddress,
        password: password
    });

    const data = <LoginResponse>res.data

    return data
}

export type RegisterResponse = { 
    error: string | undefined
    token: string | undefined
}

export const register = async (fullname: string, emailAddress: string, password: string): Promise<RegisterResponse> => {
    const res = await axios.post(`${apiUrl}auth/register`, {
        fullname: fullname,
        emailAddress: emailAddress,
        password: password
    });

    const data = <RegisterResponse>res.data

    return data
}