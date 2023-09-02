import { useQueryClient } from "@tanstack/react-query"
import { Link, useNavigate } from "react-router-dom"
import useAuth from "src/hooks/useAuth"

export default function Navbar() {
    const { logout } = useAuth()

    const queryClient = useQueryClient()

    const onLogoutClick = () => {
        var confirm = window.confirm("logout from app?")
        if(confirm){
            //clear tasks cache
            queryClient.clear()
            logout()
        }
    }

    return (
        <ul className="menu menu-horizontal px-1">
            <li><Link to="/">My Tasks</Link></li>
            <li><a onClick={() => onLogoutClick()}>Logout</a></li>
        </ul>
    )
}