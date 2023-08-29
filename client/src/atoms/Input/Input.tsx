interface InputProps extends React.InputHTMLAttributes<HTMLInputElement> {
}

const Input: React.FC<InputProps> = ({ ...props })  => {
    return (
        <input className="input input-bordered" {...props}/>
    )
}

export default Input;