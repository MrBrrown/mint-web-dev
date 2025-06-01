import React from 'react'
import './Navbar.css';
import { NavLink } from 'react-router-dom'

export default function Navbar() {
  return (
    <nav className="navbar">
      <ul className="navbar-list">
        <li>
          <NavLink to="/" className={({ isActive }) => (isActive ? 'active' : '')}>
            Главная
          </NavLink>
        </li>
        <li>
          <NavLink to="/cart" className={({ isActive }) => (isActive ? 'active' : '')}>
            Корзина
          </NavLink>
        </li>
        <li>
          <NavLink to="/order" className={({ isActive }) => (isActive ? 'active' : '')}>
            Оформление
          </NavLink>
        </li>
        <li className="navbar-login">
          <NavLink to="/login" className={({ isActive }) => (isActive ? 'active' : '')}>
            Войти
          </NavLink>
        </li>
      </ul>
    </nav>
  )
}
