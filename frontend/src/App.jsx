import React from 'react';
import { Routes, Route, Navigate } from 'react-router-dom';
import Navbar from './components/Navbar';
import Home from './pages/Home';
import Product from './pages/Product';
import Cart from './pages/Cart';
import Order from './pages/Order';
import LoginPage from './pages/Login'
import AdminPage from './pages/Admin'

export default function App() {
  return (
    <>
      <Navbar />
      <Routes>
        <Route path="/" element={<Home />} />
         <Route path="/login" element={<LoginPage />} />
         <Route path="/admin" element={<AdminPage />} />
        <Route path="/product/:id" element={<Product />} />
        <Route path="/cart" element={<Cart />} />
        <Route path="/order" element={<Order />} />
        <Route path="*" element={<Navigate to="/" replace />} />
      </Routes>
    </>
  );
}