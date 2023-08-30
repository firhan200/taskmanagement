interface TextareaProps extends React.TextareaHTMLAttributes<HTMLTextAreaElement> {
}

const Textarea: React.FC<TextareaProps> = ({ ...props })  => {
    return (
        <textarea rows={5} className="textarea textarea-bordered" {...props}/>
    )
}

export default Textarea;