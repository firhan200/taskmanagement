import { useMutation } from "@tanstack/react-query";
import { AxiosError } from "axios";
import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import Alert from "src/atoms/Alert/Alert";
import Button from "src/atoms/Button/Button"
import Loading from "src/atoms/Loading/Loading";
import Typography from "src/atoms/Typography/Typography";
import useAuth from "src/hooks/useAuth";
import FormControl from "src/molecules/FormControl/FormControl"
import { login, LoginResponse } from "src/services/authService";

export default function LoginForm() {
    const { authorize } = useAuth()
    const navigate = useNavigate()

    const [emailAddress, setEmailAddress] = useState<string>('test@gmail.com')
    const [password, setPassword] = useState<string>('123456')
    const [error, setError] = useState<string>("")

    const { isLoading, mutate } = useMutation({
        mutationKey: ['login'],
        mutationFn: () => {
            return login(emailAddress, password)
        },
        onMutate: () => {
            setError("")
        },
        onSuccess: (data) => {
            if(typeof data.token !== 'undefined' && data.token !== ""){
                authorize(data.token!)
            }

            setError("Oopps something wrong");
        },
        onError: (err: AxiosError) => {
            const res: LoginResponse = err.response?.data as LoginResponse
            setError(res.error ?? "")
        }
    })

    const submit = (e: React.FormEvent) => {
        e.preventDefault()

        mutate()
    }

    return (
        <div className="card flex-shrink-0 w-full max-w-sm shadow-2xl bg-base-100">
            <form onSubmit={submit} className="card-body">
                {
                    error !== "" ? <Alert type="error" text={error} /> : null
                }
                
                <FormControl disabled={ isLoading } value={emailAddress} onChange={e => setEmailAddress(e.target.value)} title="Email" type="text" placeholder="Email Address" required/>
                <FormControl disabled={ isLoading } value={password} onChange={e => setPassword(e.target.value)} title="Password" type="password" placeholder="Password" required/>
                <div className="text-center my-4">
                    <Link to="/register" className="label-text-alt link link-hover">
                        <Typography size="md">Sign up now &gt;</Typography>
                    </Link>
                </div>
                <Button disabled={ isLoading } colorType="primary" size="md">
                    { isLoading ? <Loading /> : 'Login' }
                </Button>
            </form>
        </div>
    );
}