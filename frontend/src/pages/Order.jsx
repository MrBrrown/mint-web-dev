import React, { useState, useEffect, useContext } from "react";
import { CartContext } from "../contexts/CartContext";
import { fetchProducts } from "../api";

export default function Order() {
  const { cart, clearCart } = useContext(CartContext);
  const [form, setForm] = useState({
    fullName: "",
    address: "",
    comment: "",
    phone: "",
  });
  const [products, setProducts] = useState([]);
  const [orderNum, setOrderNum] = useState(null);

  useEffect(() => {
    fetchProducts().then(setProducts);
  }, []);

  const handleChange = (e) =>
    setForm((f) => ({ ...f, [e.target.name]: e.target.value }));
  const handleSubmit = (e) => {
    e.preventDefault();
    const items = Object.entries(cart).map(([id, qty]) => ({
      product_id: parseInt(id),
      quantity: qty,
    }));
    console.log("Mock send order:", { customer: form, items });
    setOrderNum(Math.floor(Math.random() * 100000));
    clearCart();
  };

  if (orderNum)
    return (
      <main className="container">
        <h1>Заказ оформлен</h1>
        <p>Номер заказа: {orderNum}</p>
      </main>
    );

  return (
    <main className="container">
      <h1>Оформление заказа</h1>
      <form onSubmit={handleSubmit} className="order-form">
        <div className="form-group">
          <label htmlFor="fullName">ФИО:</label>
          <input
            id="fullName"
            name="fullName"
            value={form.fullName}
            onChange={handleChange}
            required
          />
        </div>

        <div className="form-group">
          <label htmlFor="address">Адрес:</label>
          <input
            id="address"
            name="address"
            value={form.address}
            onChange={handleChange}
            required
          />
        </div>
        <div className="form-group">
          <label htmlFor="phone">Телефон:</label>
          <input
            id="phone"
            name="phone"
            type="tel"
            placeholder="+7 (___) ___-__-__"
            inputMode="tel"
            value={form.phone}
            onChange={handleChange}
            required
          />
        </div>

        <div className="form-group">
          <label htmlFor="comment">Комментарий:</label>
          <textarea
            id="comment"
            name="comment"
            value={form.comment}
            onChange={handleChange}
          />
        </div>

        <button type="submit" className="res-submit-btn">
          Отправить
        </button>
      </form>
    </main>
  );
}
