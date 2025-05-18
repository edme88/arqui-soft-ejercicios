import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { BrowserRouter, Routes, Route } from "react-router";
import './index.css';
import Login from './Login.jsx';
import Activities from './Activities.jsx';
import Layout from './Layout.jsx';
import Home from './Home.jsx';

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <BrowserRouter>
      <Routes>
        <Route path='/' element={ <Layout />} >
          <Route index element={ <Home />} />
          <Route path='/login' element={ <Login /> } />
          <Route path='/actividades' element={ <Activities /> } />
        </Route>
      </Routes>
    </BrowserRouter>
  </StrictMode>,
)
