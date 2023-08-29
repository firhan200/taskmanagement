import { Suspense } from "react";
import { Navigate, Outlet } from "react-router-dom";
import Container from "src/atoms/Container/Container";
import useAuth from "src/hooks/useAuth";

const UnauthorizeRoute = () => {
	const { isAuthorize } = useAuth();

	if(!isAuthorize){
		return <Navigate to={`/login`}/>
	}

	return (
		<div style={{
			minHeight: '100vh'
		}}>
			<Container>
				<Suspense>
					<Outlet />
				</Suspense>
			</Container>
		</div>
	);
}

export default UnauthorizeRoute;