CREATE TABLE stock (
  id INTEGER PRIMARY KEY,
  shop_id VARCHAR(255) NOT NULL,
  prod_id VARCHAR(255) NOT NULL,
  operation_type TEXT CHECK (operationType IN ('in', 'out')),
  amount DECIMAL(10,2) NOT NULL CHECK (amount >= 0),
  description TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);