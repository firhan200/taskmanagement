import { useMutation } from "@tanstack/react-query";
import { useState } from "react";
import { Link } from "react-router-dom";
import Button from "src/atoms/Button/Button"
import Loading from "src/atoms/Loading/Loading";
import Typography from "src/atoms/Typography/Typography";
import FormControl from "src/molecules/FormControl/FormControl"
import { login } from "src/services/authService";

export default function LoginForm() {
    const [emailAddress, setEmailAddress] = useState<string>('')
    const [password, setPassword] = useState<string>('')

    const loginMutation = useMutation({
        mutationKey: ['login'],
        mutationFn: () => {
            return login(emailAddress, password)
        }
    })

    const submit = (e: React.FormEvent) => {
        e.preventDefault()

        loginMutation.mutate()
    }

    return (
        <div className="card flex-shrink-0 w-full max-w-sm shadow-2xl bg-base-100">
            <form onSubmit={submit} className="card-body">
                <FormControl disabled={ loginMutation.isLoading } value={emailAddress} onChange={e => setEmailAddress(e.target.value)} title="Email" type="text" placeholder="Email Address" required/>
                <FormControl disabled={ loginMutation.isLoading } value={password} onChange={e => setPassword(e.target.value)} title="Password" type="password" placeholder="Password" required/>
                <div className="text-center my-4">
                    <Link to="/register" className="label-text-alt link link-hover">
                        <Typography size="md">Sign up now &gt;</Typography>
                    </Link>
                </div>
                <Button disabled={ loginMutation.isLoading } colorType="primary" size="md">
                    { loginMutation.isLoading ? <Loading /> : 'Login' }
                </Button>
            </form>
        </div>
    );
}