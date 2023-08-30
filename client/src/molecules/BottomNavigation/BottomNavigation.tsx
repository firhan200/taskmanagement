import { Link } from "react-router-dom";
import useAuth from "src/hooks/useAuth";

export default function BottomNavigation() {
    const { logout } = useAuth()

    return (
        <div className="lg:hidden btm-nav z-50">
            <Link to="/">
                My Tasks
            </Link>
            <a onClick={() => logout()}>
                Logout
            </a>
        </div>
    );
}