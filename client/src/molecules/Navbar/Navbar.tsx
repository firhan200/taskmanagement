import { Link, useNavigate } from "react-router-dom"
import useAuth from "src/hooks/useAuth"

export default function Navbar() {
    const { logout } = useAuth()

    return (
        <ul className="menu menu-horizontal px-1">
            <li><Link to="/">My Tasks</Link></li>
            <li><a onClick={() => logout()}>Logout</a></li>
        </ul>
    )
}