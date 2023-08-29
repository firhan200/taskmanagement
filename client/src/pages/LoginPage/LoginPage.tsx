import Typography from "src/atoms/Typography/Typography";
import LoginForm from "src/organisms/LoginForm/LoginForm";

const LoginPage = () => {
    return (
        <div className="hero min-h-screen">
            <div className="hero-content flex-col lg:flex-row-reverse">
                <div className="text-center lg:text-left">
                    <Typography size="lg">Login now!</Typography>
                    <div className="my-6">
                        <Typography size="sm">Provident cupiditate voluptatem et in. Quaerat fugiat ut assumenda excepturi exercitationem quasi. In deleniti eaque aut repudiandae et a id nisi.</Typography>
                    </div>
                </div>
                <LoginForm />
            </div>
        </div>
    );
}

export default LoginPage;