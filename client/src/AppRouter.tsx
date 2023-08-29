/* init page routing */
import { lazy } from "react";
import {
	RouterProvider,
	createBrowserRouter
} from "react-router-dom";
import UnauthorizeRoute from "src/templates/UnauthorizeRoute";
import ProtectedRoute from "src/templates/ProtectedRoute";
// import GenericTemplate from "./templates/GenericTemplate";

//init every pages using lazy
const dashboardPageLazy = lazy(() => import('src/pages/DashboardPage/DashboardPage'));
const loginPageLazy = lazy(() => import('src/pages/LoginPage/LoginPage'));
const registerPageLazy = lazy(() => import('src/pages/RegisterPage/RegisterPage'));

//create root router
const router = createBrowserRouter([
	{
		element: <UnauthorizeRoute />,
		children: [
			{
				path: "/login",
				Component: loginPageLazy
			},
			{
				path: "/register",
				Component: registerPageLazy
			},
		]
	},
    {
		element: <ProtectedRoute />,
		children: [
			{
				path: "/",
				Component: dashboardPageLazy
			},
		]
	},
]);
/* init page routing */

const AppRouter = () => {
	return (
		<RouterProvider router={router}/>
	);
}

export default AppRouter;