interface AlertProps extends React.HTMLAttributes<HTMLDivElement>{
    type: "error" | "success" | "info"
    text: string
}

const Alert: React.FC<AlertProps> = ({ text, type, ...props }) => {
    const getAlertColor = () => {
        if(type == "error"){
            return "alert-error"
        }
        else if(type == "success"){
            return "alert-success"
        }
       
        return ""
    }

    return (
        <div {...props} className={`alert ${getAlertColor()}`}>
            { text }
        </div>
    )
}

export default Alert;