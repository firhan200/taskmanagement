import { create } from "zustand";

const TOKEN_KEY: string = "__task_management_token";

export const getToken = (): string | null => {
	return localStorage.getItem(TOKEN_KEY);
}

type AuthStoreState = {
    token: string | null
    isAuth: boolean
    authorize: (token: string) => void
	logout: () => void
}

export const authStore = create<AuthStoreState>()((set) => ({
    token: getToken(),
    isAuth: getToken() !== null,
    authorize: (token: string) => set(() => {
		localStorage.setItem(TOKEN_KEY, token);
        return {
            token: token,
			isAuth: true
        }
    }),
	logout: () => set(() => {
		localStorage.removeItem(TOKEN_KEY)
        return {
            token: null,
			isAuth: false
        }
    })
}))

export default function useAuth(){
	const [token, isAuthorize, authorize, logout] = authStore((state) => [
		state.token,
		state.isAuth,
		state.authorize,
		state.logout,
	]) 

	return {
		token,
		isAuthorize,
		authorize,
		logout
	}
}