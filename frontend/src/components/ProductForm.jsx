import React, { useState, useEffect } from 'react';
import './ProductForm.css';

export default function ProductForm({ mode, initialData = {}, onSubmit, onCancel }) {
  const isView = mode === 'view';
  const [product, setProduct] = useState({
    name: '',
    desc: '',
    price: '',
    attribs: '{}',
    ...initialData
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setProduct(prev => ({ ...prev, [name]: value }));
  };

  const handleSubmit = (e) => {
    e.preventDefault();

    try {
      const parsedAttribs = JSON.parse(product.attribs || '{}');
      onSubmit({ ...product, price: parseFloat(product.price), attribs: parsedAttribs });
    } catch (err) {
      alert('Невалидный JSON в поле атрибутов');
    }
  };

  return (
    <form className="product-form" onSubmit={handleSubmit}>
      <h3>{mode === 'edit' ? 'Редактировать' : mode === 'view' ? 'Просмотр' : 'Новый товар'}</h3>

      <label>Название:</label>
      <input type="text" name="name" value={product.name} onChange={handleChange} disabled={isView} required />

      <label>Описание:</label>
      <textarea name="desc" value={product.desc} onChange={handleChange} disabled={isView} required />

      <label>Цена:</label>
      <input type="number" name="price" value={product.price} onChange={handleChange} disabled={isView} required />

      <label>Атрибуты (JSON):</label>
      <textarea name="attribs" value={JSON.stringify(product.attribs || {}, null, 2)} onChange={handleChange} disabled={isView} />

      {!isView && <button type="submit">{mode === 'edit' ? 'Сохранить' : 'Создать'}</button>}
      <button type="button" onClick={onCancel}>Отмена</button>
    </form>
  );
}
