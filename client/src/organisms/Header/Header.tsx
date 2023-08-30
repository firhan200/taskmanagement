import { Link } from "react-router-dom";
import BottomNavigation from "src/molecules/BottomNavigation/BottomNavigation";
import Navbar from "src/molecules/Navbar/Navbar";
import ThemeToggle from "src/molecules/ThemeToggle/ThemeToggle";

const Header = () => {
	return (
		<>
			<div className="navbar sticky top-0 z-50 bg-base-100 shadow-lg">
				<div className="navbar-start">
					<Link to="/" className="btn btn-ghost normal-case text-xl">Task Management</Link>
				</div>
				<div className="navbar-center hidden lg:flex">
					<Navbar />
				</div>
				<div className="navbar-end">
					<ThemeToggle />
				</div>
			</div>

			<div>
				<BottomNavigation />
			</div>
		</>
	);
}

export default Header;