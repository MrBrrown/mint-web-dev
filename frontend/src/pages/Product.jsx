import React, { useEffect, useState, useContext } from 'react';
import { useParams } from 'react-router-dom';
import { CartContext } from '../contexts/CartContext';
import { fetchProductById } from '../api';

export default function Product() {
  const { id } = useParams();
  const [product, setProduct] = useState(null);
  const { addToCart } = useContext(CartContext);

  useEffect(() => {
    fetchProductById(parseInt(id)).then(setProduct);
  }, [id]);

  if (!product) return <p>Загрузка...</p>;

  return (
    <main className="container">
      <div className="product-details">
        <img src={product.image} alt={product.name} className="product-image" />
        <h1>{product.name}</h1>
        <p>{product.description}</p>
        <p>{product.cost} ₽</p>
        <button className="res-submit-btn" onClick={() => addToCart(product.id)}>
          В корзину
        </button>
      </div>
    </main>
  );
}
