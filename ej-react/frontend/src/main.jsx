import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { BrowserRouter, Routes, Route } from "react-router";
import './index.css';
import Login from './Login.jsx';
import Layout from './Layout.jsx';
import Home from './Home.jsx';
import Activities from './Activities.jsx';
import Vinilos from './Vinilos.jsx';
import AddVinyl from './AddVinyl.jsx';

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <BrowserRouter>
    <Routes>
      <Route path='/' element={ <Layout /> }>
        <Route index element={ <Home />} />
        <Route path="/login" element={<Login />} />
        <Route path='/actividades' element={<Activities /> } />
        <Route path='/vinilos' element={<Vinilos /> } />
        <Route path='/add-vinyl' element={<AddVinyl /> } />
      </Route>
    </Routes>
  </BrowserRouter>
  </StrictMode>,
)
