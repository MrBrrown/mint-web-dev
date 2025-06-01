import React, { useEffect, useState } from "react";
import ProductForm from "../components/ProductForm";
import "./ProductsTab.css";

export default function ProductsTab() {
  const [products, setProducts] = useState([]);
  const [selected, setSelected] = useState(null);
  const [mode, setMode] = useState(null); // 'create' | 'edit' | 'view'

  const token = localStorage.getItem("token");

  const fetchProducts = () => {
    fetch("/products/")
      .then((res) => res.json())
      .then(setProducts)
      .catch((err) => console.error("Ошибка загрузки:", err));
  };

  useEffect(fetchProducts, []);

  const handleCreate = () => {
    setSelected(null);
    setMode("create");
  };

  const handleView = (product) => {
    setSelected(product);
    setMode("view");
  };

  const handleEdit = (product) => {
    setSelected(product);
    setMode("edit");
  };

  const handleDelete = (id) => {
    if (!window.confirm("Удалить товар?")) return;
    fetch(`/products/${id}`, {
      method: "DELETE",
      headers: { Authorization: `Bearer ${token}` },
    })
      .then(() => fetchProducts())
      .catch((err) => alert("Ошибка удаления: " + err));
  };

  const handleSubmit = (product) => {
    const isEdit = mode === "edit";
    const url = isEdit ? `/products/${product.id}` : "/products/";
    const method = isEdit ? "PUT" : "POST";

    fetch(url, {
      method,
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(product),
    })
      .then(() => {
        fetchProducts();
        setMode(null);
      })
      .catch((err) => alert("Ошибка сохранения: " + err));
  };

  const closeForm = () => {
    setMode(null);
    setSelected(null);
  };

  return (
    <div className="products-container">
      <div className="products-header">
        <h2>Список товаров</h2>
        <button onClick={handleCreate}>+ Новый товар</button>
      </div>

      {mode ? (
        <ProductForm
          mode={mode}
          initialData={selected}
          onSubmit={handleSubmit}
          onCancel={closeForm}
        />
      ) : products.length === 0 ? (
        <div className="empty-state">
          <p>Пока нет ни одного товара.</p>
          <p>
            Нажмите <strong>«+ Новый товар»</strong>, чтобы добавить первый.
          </p>
        </div>
      ) : (
        <table className="products-table">
          <thead>
            <tr>
              <th>Название</th>
              <th>Описание</th>
              <th>Цена</th>
              <th>ID</th>
              <th></th>
            </tr>
          </thead>
          <tbody>
            {products.map((p) => (
              <tr key={p.id}>
                <td>{p.name}</td>
                <td>{p.desc}</td>
                <td>{p.price} ₽</td>
                <td>{p.id}</td>
                <td className="actions">
                  <div>
                    <button onClick={() => handleView(p)} title="Просмотр">
                      <span className="material-icons">visibility</span>
                    </button>
                    <button onClick={() => handleEdit(p)} title="Редактировать">
                      <span className="material-icons">edit</span>
                    </button>
                    <button onClick={() => handleDelete(p.id)} title="Удалить">
                      <span className="material-icons">delete</span>
                    </button>
                  </div>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
}
