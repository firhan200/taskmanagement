import { Suspense } from "react";
import { Navigate, Outlet } from "react-router-dom";
import Container from "src/atoms/Container/Container";
import useAuth from "src/hooks/useAuth";
import useTheme from "src/hooks/useTheme";
import Modal from "src/molecules/Modal/Modal";
import Footer from "src/organisms/Footer/Footer";
import Header from "src/organisms/Header/Header";

const UnauthorizeRoute = () => {
	const { isAuthorize } = useAuth();
	const { theme } = useTheme()

	if (!isAuthorize) {
		return <Navigate to={`/login`} />
	}

	return (
		<>
			<Modal />
			<div style={{
				minHeight: '100vh'
			}} data-theme={theme}>
				<Header />
				<Container>
					<Suspense>
						<Outlet />
					</Suspense>
				</Container>
				<Footer />
			</div>
		</>
	);
}

export default UnauthorizeRoute;