CREATE EXTENSION IF NOT EXISTS vector;

ALTER TABLE menu_items 
ADD COLUMN IF NOT EXISTS embedding vector(384);

CREATE INDEX IF NOT EXISTS idx_menu_items_embedding 
ON menu_items USING ivfflat (embedding vector_cosine_ops);

