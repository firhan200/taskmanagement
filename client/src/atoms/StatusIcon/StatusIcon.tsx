import useTheme from "src/hooks/useTheme"

type StatusIconProps = {
    status: "Not urgent" | "Overdue" | "Due soon"
}

export default function StatusIcon({ status }: StatusIconProps) {
    const { theme } = useTheme()

    const getIcon = () => {
        if (status == "Overdue") {
            const color = theme === "dark" ? "#888888" : "#cecece"

            return (
                <svg width="64px" height="64px" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg"><g id="SVGRepo_bgCarrier" strokeWidth="0"></g><g id="SVGRepo_tracerCarrier" strokeLinecap="round" strokeLinejoin="round"></g><g id="SVGRepo_iconCarrier"> <path d="M3 5.5L5 3.5M21 5.5L19 3.5M9 9.5L15 15.5M15 9.5L9 15.5M20 12.5C20 16.9183 16.4183 20.5 12 20.5C7.58172 20.5 4 16.9183 4 12.5C4 8.08172 7.58172 4.5 12 4.5C16.4183 4.5 20 8.08172 20 12.5Z" stroke={color} strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"></path> </g></svg>

            )
        }

        if (status == "Due soon") {
            return (
                <svg width="64px" height="64px" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" transform="rotate(0)" stroke="#bd6565"><g id="SVGRepo_bgCarrier" strokeWidth="0"></g><g id="SVGRepo_tracerCarrier" strokeLinecap="round" strokeLinejoin="round" stroke="#CCCCCC" strokeWidth="0.048"></g><g id="SVGRepo_iconCarrier"> <path d="M3 5.5L5 3.5M21 5.5L19 3.5M12 8.5V12.5L14 14.5M20 12.5C20 16.9183 16.4183 20.5 12 20.5C7.58172 20.5 4 16.9183 4 12.5C4 8.08172 7.58172 4.5 12 4.5C16.4183 4.5 20 8.08172 20 12.5Z" stroke="#ce3b3b" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"></path> </g></svg>
            )
        }

        if(status == "Not urgent") {
            const color = theme === "dark" ? "#fff" : "#000"

            return (
                <svg width="64px" height="64px" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg"><g id="SVGRepo_bgCarrier" strokeWidth="0"></g><g id="SVGRepo_tracerCarrier" strokeLinecap="round" strokeLinejoin="round"></g><g id="SVGRepo_iconCarrier"> <path d="M9 12L11 14L15 10M12 3L13.9101 4.87147L16.5 4.20577L17.2184 6.78155L19.7942 7.5L19.1285 10.0899L21 12L19.1285 13.9101L19.7942 16.5L17.2184 17.2184L16.5 19.7942L13.9101 19.1285L12 21L10.0899 19.1285L7.5 19.7942L6.78155 17.2184L4.20577 16.5L4.87147 13.9101L3 12L4.87147 10.0899L4.20577 7.5L6.78155 6.78155L7.5 4.20577L10.0899 4.87147L12 3Z" stroke={color} strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"></path> </g></svg>
            )
        }
    }

    return (
        <div className="tooltip" data-tip={ status }>
            { getIcon() }
        </div>
    )
}