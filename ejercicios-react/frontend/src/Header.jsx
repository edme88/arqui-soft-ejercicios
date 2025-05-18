import { useNavigate } from "react-router-dom";

const Header = () => {
    const navigate = useNavigate();
    const isLoggedIn = localStorage.getItem("isLoggedIn") === "true";

    const logout = () => {
        localStorage.removeItem("isLoggedIn");
        navigate("/");
    };

    return (
        <header>
            <h1>Gym</h1>
            <nav>
                <a href="/">Home</a>
                {isLoggedIn ? (
                    <>
                        <button onClick={logout}>Cerrar sesión</button>
                    </>
                ) : (
                    <a href="/login">Login</a>
                )}
                <a href="/actividades">Actividades</a>
            </nav>
        </header>
    );
}

export default Header;