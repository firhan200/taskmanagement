interface TypographyProps extends React.DOMAttributes<HTMLDivElement> {
    size: "sm" | "md" | "lg"
}

const Typography: React.FC<TypographyProps> = ({ children, size, ...props })  => {
    const renderSize = () => {
        if(size == "sm"){
            return "text-md"
        }
        else if(size == "md"){
            return "text-lg"
        }
        else if(size == "lg"){
            return "text-5xl font-bold"
        }
    }

    return (
        <div {...props} className={renderSize()}>
            { children }
        </div>
    )
}

export default Typography;