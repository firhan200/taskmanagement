import Input from "src/atoms/Input/Input";

interface FormControlProps extends React.InputHTMLAttributes<HTMLInputElement> {
    title: string
}

const FormControl : React.FC<FormControlProps> = ({ title, ...props }) => {
    return (
        <div className="form-control">
            <label className="label">
                <span className="label-text">{ title }</span>
            </label>
            <Input {...props}/>
        </div>
    );
}

export default FormControl;