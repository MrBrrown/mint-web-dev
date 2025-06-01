import React, { useState } from 'react';
import ProductsTab from './ProductsTab';
import OrdersTab from './OrdersTab';
import './Admin.css';

export default function AdminPage() {
  const [activeTab, setActiveTab] = useState('products');

  return (
    <div className="admin-container">
      <h1 className="admin-title">Панель администратора</h1>

      <div className="tab-buttons">
        <button
          className={activeTab === 'products' ? 'active' : ''}
          onClick={() => setActiveTab('products')}
        >
          Товары
        </button>
        <button
          className={activeTab === 'orders' ? 'active' : ''}
          onClick={() => setActiveTab('orders')}
        >
          Заказы
        </button>
      </div>

      <div className="tab-content">
        {activeTab === 'products' ? <ProductsTab /> : <OrdersTab />}
      </div>
    </div>
  );
}
