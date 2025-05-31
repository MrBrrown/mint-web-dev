import React, { createContext, useState, useEffect } from 'react';

const CART_KEY = 'cartCounts';
export const CartContext = createContext();

export function CartProvider({ children }) {
  const [cart, setCart] = useState(() => {
    try {
      return JSON.parse(localStorage.getItem(CART_KEY)) || {};
    } catch {
      return {};
    }
  });

  useEffect(() => {
    localStorage.setItem(CART_KEY, JSON.stringify(cart));
  }, [cart]);

  const addToCart = (id) =>
    setCart(prev => ({ ...prev, [id]: (prev[id] || 0) + 1 }));

  const removeFromCart = (id) => {
    setCart(prev => {
      const next = { ...prev };
      delete next[id];
      return next;
    });
  };

  const clearCart = () => setCart({});

  return (
    <CartContext.Provider value={{ cart, addToCart, removeFromCart, clearCart }}>
      {children}
    </CartContext.Provider>
  );
}
