CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    total_price NUMERIC(10, 2) NOT NULL,
    items JSONB NOT NULL DEFAULT '[]',         -- Список товаров в заказе
    user_info JSONB NOT NULL DEFAULT '{}',     -- Информация о заказчике (имя, телефон, адрес и т.п.)
    status VARCHAR(32) NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
