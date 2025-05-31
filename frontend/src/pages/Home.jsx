import React, { useEffect, useState, useContext } from 'react';
import ProductCard from '../components/ProductCard';
import { CartContext } from '../contexts/CartContext';
import { fetchProducts } from '../api';

export default function Home() {
  const [products, setProducts] = useState([]);
  const { addToCart } = useContext(CartContext);

  useEffect(() => {
    fetchProducts().then(setProducts);
  }, []);

  return (
    <main className="container">
      <h1>Магазин лампочек</h1>
      <div className="cards-container">
        {products.map(p => (
          <div key={p.id}>
            <ProductCard product={p} />
            <button className="res-submit-btn" onClick={() => addToCart(p.id)}>
              В корзину
            </button>
          </div>
        ))}
      </div>
    </main>
  );
}
