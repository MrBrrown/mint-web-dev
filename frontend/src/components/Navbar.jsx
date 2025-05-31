import React from 'react';
import { NavLink } from 'react-router-dom';

export default function Navbar() {
  return (
    <nav className="navbar">
      <ul className="navbar-list">
        <li>
          <NavLink to="/" className={({ isActive }) => isActive ? 'active' : ''}>
            Главная
          </NavLink>
        </li>
        <li>
          <NavLink to="/cart" className={({ isActive }) => isActive ? 'active' : ''}>
            Корзина
          </NavLink>
        </li>
        <li>
          <NavLink to="/order" className={({ isActive }) => isActive ? 'active' : ''}>
            Оформление
          </NavLink>
        </li>
      </ul>
    </nav>
  );
}