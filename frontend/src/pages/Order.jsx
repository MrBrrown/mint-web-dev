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

  const handleSubmit = async (e) => {
    e.preventDefault();

    const items = Object.entries(cart).map(([id, qty]) => ({
      id: parseInt(id),
      quantity: qty,
    }));

    if (items.length === 0) {
      alert("Корзина пуста. Добавьте товары перед оформлением заказа.");
      return;
    }

    const orderPayload = {
      items,
      user_info: form,
      status: "pending",
    };

    try {
      const res = await fetch("/orders/", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(orderPayload),
      });

      if (!res.ok) throw new Error("Ошибка при отправке заказа");

      const data = await res.json();
      setOrderNum(data.id);
      clearCart();
    } catch (err) {
      console.error("Ошибка создания заказа:", err);
      alert("Не удалось оформить заказ. Попробуйте позже.");
    }
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
