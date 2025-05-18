import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import './Login.css';

const Login = () => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const navigate = useNavigate();

    const handleLogin = async (e) => {
        e.preventDefault();
        
        if(username=="admin" && password=="admin"){
            console.log("Login OK");
            navigate("/actividades");
        }else{
            console.log("Login incorrecto");
        }
    }

    return (
        <div className="login-container">
            <form className="login-form" onSubmit={ handleLogin }>
                <h2> Iniciar Sesión </h2>
                <input
                    type="text"
                    placeholder="Usuario"
                    value={ username }
                    onChange={(e) => setUsername(e.target.value)}
                    required
                />
                <input
                    type="password"
                    placeholder="Contraseña"
                    value={ password }
                    onChange={(e) => setPassword(e.target.value)}
                    required
                />
                <button type="submit"> Ingresar </button>
            </form>
        </div>
    )
}

export default Login;