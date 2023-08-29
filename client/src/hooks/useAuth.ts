const TOKEN_KEY: string = "__task_management_token";

export default function useAuth(){
	const getToken = (): string | null => {
		return localStorage.getItem(TOKEN_KEY);
	}

	//check if jwt token is exist on local storage
	const isAuthorize : boolean = getToken() !== null;

	const authorize = (jwt: string) => {
		localStorage.setItem(TOKEN_KEY, jwt);
	}

	return {
		getToken,
		isAuthorize,
		authorize
	}
}