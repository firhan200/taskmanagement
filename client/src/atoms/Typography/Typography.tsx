interface TypographyProps extends React.HTMLAttributes<HTMLDivElement> {
    size: "sm" | "md" | "lg"
    truncate?: undefined | boolean
}

const Typography: React.FC<TypographyProps> = ({ children, truncate, size, ...props })  => {
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

    if(typeof truncate !== 'undefined' && truncate){
        props.className += " truncate"
    }

    props.className += " break-words"

    return (
        <div {...props}>
            { children }
        </div>
    )
}

export default Typography;