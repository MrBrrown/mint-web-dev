import React from 'react';
import { Routes, Route, Navigate } from 'react-router-dom';
import Navbar from './components/Navbar';
import Home from './pages/Home';
import Product from './pages/Product';
import Cart from './pages/Cart';
import Contacts from './pages/Contacts';
import Order from './pages/Order';

export default function App() {
  return (
    <>
      <Navbar />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/product/:id" element={<Product />} />
        <Route path="/cart" element={<Cart />} />
        <Route path="/contacts" element={<Contacts />} />
        <Route path="/order" element={<Order />} />
        <Route path="*" element={<Navigate to="/" replace />} />
      </Routes>
    </>
  );
}