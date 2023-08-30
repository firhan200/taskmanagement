import Typography from "src/atoms/Typography/Typography";
import RegisterForm from "src/organisms/RegisterForm/RegisterForm";

const RegisterPage = () => {
    return (
        <div className="hero min-h-screen">
            <div className="hero-content flex-col lg:flex-row-reverse">
                <div className="text-center lg:text-left">
                    <Typography size="lg">Join Us!</Typography>
                    <div className="my-6">
                        <Typography size="sm">Provident cupiditate voluptatem et in. Quaerat fugiat ut assumenda excepturi exercitationem quasi. In deleniti eaque aut repudiandae et a id nisi.</Typography>
                    </div>
                </div>
                <RegisterForm />
            </div>
        </div>
    );
}

export default RegisterPage;