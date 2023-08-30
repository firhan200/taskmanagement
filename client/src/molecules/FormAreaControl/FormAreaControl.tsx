import Textarea from "src/atoms/Textarea/Textarea";

interface FormAreaControlProps extends React.TextareaHTMLAttributes<HTMLTextAreaElement> {
    title: string
}

const FormAreaControl : React.FC<FormAreaControlProps> = ({ title, ...props }) => {
    return (
        <div className="form-control">
            <label className="label">
                <span className="label-text">{ title }</span>
            </label>
            <Textarea {...props}/>
        </div>
    );
}

export default FormAreaControl;