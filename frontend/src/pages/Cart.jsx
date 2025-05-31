import React, { useEffect, useState, useContext } from 'react';
import { useNavigate } from 'react-router-dom';
import { CartContext } from '../contexts/CartContext';
import { fetchProducts } from '../api';

export default function Cart() {
  const { cart, removeFromCart, clearCart } = useContext(CartContext);
  const [products, setProducts] = useState([]);
  const navigate = useNavigate();

  useEffect(() => {
    fetchProducts().then(setProducts);
  }, []);

  const cartItems = Object.entries(cart)
    .map(([id, qty]) => {
      const p = products.find(prod => prod.id === parseInt(id));
      return p && { ...p, quantity: qty };
    })
    .filter(Boolean);

  const total = cartItems.reduce((sum, item) => sum + item.cost * item.quantity, 0);

  return (
    <main className="container">
      <h1>Корзина</h1>
      {cartItems.length === 0 ? (
        <p>Ваша корзина пуста.</p>
      ) : (
        <>
          <table className="cart-table">
            <thead>
              <tr>
                <th>Название</th>
                <th>Цена</th>
                <th>Кол-во</th>
                <th>Сумма</th>
                <th>Действие</th>
              </tr>
            </thead>
            <tbody>
              {cartItems.map(item => (
                <tr key={item.id}>
                  <td>{item.name}</td>
                  <td>{item.cost} ₽</td>
                  <td>{item.quantity}</td>
                  <td>{item.cost * item.quantity} ₽</td>
                  <td>
                    <button onClick={() => removeFromCart(item.id)}>Удалить</button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
          <div className="cart-summary">
            <span>
              Итого: <strong>{total} ₽</strong>
            </span>
            <button onClick={() => navigate('/order')} className="order-btn">
              Оформить
            </button>
            <button onClick={clearCart} className="order-btn">
              Очистить
            </button>
          </div>
        </>
      )}
    </main>
  );
}
