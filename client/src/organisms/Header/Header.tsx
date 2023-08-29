import { Link } from "react-router-dom";
import BottomNavigation from "src/molecules/BottomNavigation/BottomNavigation";
import ThemeToggle from "src/molecules/ThemeToggle/ThemeToggle";

const Header = () => {
	return (
		<>
			<div className="navbar sticky top-0 z-50 bg-base-100 shadow-lg">
				<div className="navbar-start">
					<Link to="/" className="btn btn-ghost normal-case text-xl">Catalogue</Link>
				</div>
				<div className="navbar-center hidden lg:flex">
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