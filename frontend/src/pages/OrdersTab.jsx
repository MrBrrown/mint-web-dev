import React, { useEffect, useState } from "react";
import "./OrdersTab.css";

const STATUS_OPTIONS = [
  "pending",
  "processing",
  "paid",
  "shipped",
  "delivered",
  "cancelled",
];

export default function OrdersTab() {
  const [orders, setOrders] = useState([]);
  const [selected, setSelected] = useState(null);
  const [mode, setMode] = useState(null);
  const [newStatus, setNewStatus] = useState("");

  const token = localStorage.getItem("token");

  const fetchOrders = () => {
    fetch("/orders/")
      .then((res) => res.json())
      .then(setOrders)
      .catch((err) => console.error("Ошибка загрузки заказов:", err));
  };

  useEffect(fetchOrders, []);

  const handleView = (order) => {
    setSelected(order);
    setNewStatus(order.status);
    setMode("view");
  };

  const handleStatusChange = (e) => {
    setNewStatus(e.target.value);
  };

  const saveStatus = () => {
    fetch(`/orders/${selected.id}/status`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify({ status: newStatus }),
    })
      .then((res) => {
        if (!res.ok) throw new Error("Failed to update status");
        return res.json();
      })
      .then((updated) => {
        setOrders((prev) =>
          prev.map((o) => (o.id === updated.id ? updated : o))
        );
        setSelected(updated);
        alert("Статус обновлён");
      })
      .catch((err) => {
        console.error("Ошибка обновления статуса:", err);
        alert("Не удалось обновить статус");
      });
  };

  const closeView = () => {
    setSelected(null);
    setMode(null);
  };

  return (
    <div className="orders-container">
      <div className="orders-header">
        <h2>Список заказов</h2>
      </div>

      {mode === "view" && selected ? (
        <div className="order-details">
          <h3>Заказ #{selected.id}</h3>
          <p><strong>Сумма:</strong> {selected.total_price} ₽</p>
          <p><strong>Пользователь:</strong> {JSON.stringify(selected.user_info)}</p>
          <p><strong>Товары:</strong></p>
          <ul>
            {selected.items.map((item, i) => (
              <li key={i}>{JSON.stringify(item)}</li>
            ))}
          </ul>

          <div className="status-edit">
            <label><strong>Статус:</strong></label>
            <select value={newStatus} onChange={handleStatusChange}>
              {STATUS_OPTIONS.map((status) => (
                <option key={status} value={status}>{status}</option>
              ))}
            </select>
            <button onClick={saveStatus}>Сохранить</button>
          </div>

          <button onClick={closeView}>Закрыть</button>
        </div>
      ) : orders.length === 0 ? (
        <div className="empty-state">
          <p>Пока нет заказов.</p>
        </div>
      ) : (
        <table className="orders-table">
          <thead>
            <tr>
              <th>ID</th>
              <th>Сумма</th>
              <th>Статус</th>
              <th>Дата</th>
              <th></th>
            </tr>
          </thead>
          <tbody>
            {orders.map((order) => (
              <tr key={order.id}>
                <td>{order.id}</td>
                <td>{order.total_price} ₽</td>
                <td>{order.status}</td>
                <td>{new Date(order.created_at).toLocaleString()}</td>
                <td className="actions">
                  <button onClick={() => handleView(order)} title="Просмотр">
                    <span className="material-icons">visibility</span>
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
}
