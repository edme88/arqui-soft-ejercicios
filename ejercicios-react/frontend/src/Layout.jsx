import Header from './Header.jsx';
import Footer from './Footer.jsx';
import { Outlet } from 'react-router-dom';

const Layout = () => {
    return (
        <>
            <Header />
            <main>
                <Outlet /> {/* Aquí se renderiza el componente de la ruta actual */}
            </main>
            <Footer />
        </>
    );
};

export default Layout;
