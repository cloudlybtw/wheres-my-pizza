-- Orders table
CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    number TEXT UNIQUE NOT NULL,
    customer_name TEXT NOT NULL,
    type TEXT NOT NULL CHECK (type IN ('dine_in', 'takeout', 'delivery')),
    table_number INTEGER,
    delivery_address TEXT,
    total_amount DECIMAL(10,2) NOT NULL,
    priority INTEGER DEFAULT 1,
    status TEXT DEFAULT 'received',
    processed_by TEXT,
    completed_at TIMESTAMPTZ
    );

-- Order items table
CREATE TABLE IF NOT EXISTS order_items (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    order_id INTEGER REFERENCES orders(id),
    name TEXT NOT NULL,
    quantity INTEGER NOT NULL,
    price DECIMAL(8,2) NOT NULL
    );

-- Order status log table
CREATE TABLE IF NOT EXISTS order_status_log (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    order_id INTEGER REFERENCES orders(id),
    status TEXT,
    changed_by TEXT,
    changed_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    notes TEXT
    );


-- Indexes for better performance
CREATE INDEX IF NOT EXISTS idx_orders_number ON orders(number);
CREATE INDEX IF NOT EXISTS idx_orders_status ON orders(status);
CREATE INDEX IF NOT EXISTS idx_orders_created_date ON orders(created_at);
CREATE INDEX IF NOT EXISTS idx_order_items_order_id ON order_items(order_id);
CREATE INDEX IF NOT EXISTS idx_order_status_log_order_id ON order_status_log(order_id);