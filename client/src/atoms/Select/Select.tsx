export type SelectOption = {
    key: string | number
    label: string
}

interface SelectProps extends React.SelectHTMLAttributes<HTMLSelectElement> {
    options: SelectOption[]
}

const Select: React.FC<SelectProps> = ({ options, ...props }) => {
    const baseClass = "select select-bordered"
    if(props.className !== undefined){
        props.className! += " "+baseClass
    }else{
        props.className = baseClass 
    }

    return (
        <select {...props}>
            {
                options.map((opt, i) => (
                    <option key={i} value={opt.key}>
                        { opt.label }
                    </option>
                ))
            }
        </select>
    )
}

export default Select;