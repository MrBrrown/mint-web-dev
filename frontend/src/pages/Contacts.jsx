import React from 'react';

export default function Contacts() {
  return (
    <main className="container">
      <header class="menu">
        <div class="container">
            <menu class="header-menu">
                <li class="main-btn"><a href="main.html">Главная</a></li>
                <li class="center-btn"><a href="contacts.html">Контакты</a></li>
                <li class="cart-btn"><a href="cart.html">Корзина</a></li>
            </menu>
        </div>
    </header>

    <main class="contacts">
        <h1>Контакты</h1>
        <p><strong>Адрес:</strong> г. Санкт-Петербург, пл. Дворцовая, д. 1</p>
        <p><strong>Телефон:</strong> 8-812-550-12-12</p>
        <p>
            <strong>Мы в телеграм:</strong>
            <a target="_blank" class="telegram-link">
                <i class="fab fa-telegram"></i> 
            </a>
        </p>
        <h2>Реквизиты</h2>
        <p>ООО “Световые решения”</p>
        <p>ИНН: 11111111</p>
        <p>Банк: Сбер</p>
        <p>Р/С: 111111111111</p>
    </main>
    </main>
  );
}
