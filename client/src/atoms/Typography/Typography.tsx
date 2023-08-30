interface TypographyProps extends React.HTMLAttributes<HTMLDivElement> {
    size: "sm" | "md" | "lg"
}

const Typography: React.FC<TypographyProps> = ({ children, size, ...props })  => {
    const renderSize = () => {
        if(size == "sm"){
            return "text-lg"
        }
        else if(size == "md"){
            return "text-xl"
        }
        else if(size == "lg"){
            return "text-5xl font-bold"
        }
    }

    if(props.className !== undefined){
        props.className! += " "+renderSize()
    }else{
        props.className = renderSize()
    }

    return (
        <div {...props}>
            { children }
        </div>
    )
}

export default Typography;