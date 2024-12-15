CREATE EXTENSION IF NOT EXISTS vector;

CREATE TABLE categories (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) UNIQUE NOT NULL,
  embedding vector(384)
);

CREATE TYPE status AS ENUM ('active', 'inactive', 'archived');

CREATE TABLE menu_items (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) UNIQUE NOT NULL, 
  description TEXT, 
  price DECIMAL(10, 2) NOT NULL,
  category_id INTEGER NOT NULL,
  image_url VARCHAR(255), 
  embedding vector(384),
  status status NOT NULL DEFAULT 'active',
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (category_id) REFERENCES categories(id) 
);


CREATE INDEX idx_menu_items_category_id ON menu_items(category_id);
CREATE INDEX idx_menu_items_embedding ON menu_items USING ivfflat (embedding vector_cosine_ops);
CREATE INDEX idx_categories_embedding ON categories USING ivfflat (embedding vector_cosine_ops);

