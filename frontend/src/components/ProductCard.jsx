import React from 'react';
import { Link } from 'react-router-dom';

export default function ProductCard({ product }) {
  return (
    <div className="card">
      <img src={product.image} alt={product.name} className="card-image" />
      <h3 className="card-title">{product.name}</h3>
      <p className="card-price">{product.cost} ₽</p>
      <Link to={`/product/${product.id}`} className="card-button">
        Подробнее
      </Link>
    </div>
  );
}
