interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
    size: "sm" | "md" | "lg"
    colorType: "primary" | "secondary" | "danger"
}

const Button: React.FC<ButtonProps> = ({ children, size, colorType, ...props })  => {
    const renderSize = () => {
        if(size == "sm"){
            return "btn-sm"
        }
        else if(size == "md"){
            return "btn-md"
        }
        else if(size == "lg"){
            return "btn-lg"
        }
    }

    const renderColorType = () => {
        if(colorType == "primary"){
            return "bg-base-300"
        }
        else if(colorType == "secondary"){
            return "bg-base-200"
        }
        else if(colorType == "danger"){
            return "bg-red-200 hover:bg-red-300 dark:hover:bg-red-700 dark:bg-red-800"
        }
    }

    return (
        <button {...props} className={`btn ${renderSize()} ${renderColorType()}`}>
            { children }
        </button>
    )
}

export default Button;