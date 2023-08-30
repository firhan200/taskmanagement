import { Suspense } from "react";
import { Navigate, Outlet } from "react-router-dom";
import Container from "src/atoms/Container/Container";
import useAuth from "src/hooks/useAuth";
import useTheme from "src/hooks/useTheme";

const UnauthorizeRoute = () => {
	const { isAuthorize } = useAuth();
	const { theme } = useTheme();

	if(isAuthorize){
		return <Navigate to="/"/>
	}

	return (
		<div style={{
			minHeight: '100vh'
		}}  data-theme={theme}>
			<Container>
				<Suspense>
					<Outlet />
				</Suspense>
			</Container>
		</div>
	);
}

export default UnauthorizeRoute;