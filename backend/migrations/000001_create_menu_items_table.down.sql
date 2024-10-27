DROP INDEX IF EXISTS idx_menu_items_category_id;
DROP INDEX IF EXISTS idx_menu_items_embedding;
DROP INDEX IF EXISTS idx_categories_embedding;
DROP TABLE IF EXISTS menu_items;
DROP TABLE IF EXISTS categories;
DROP EXTENSION IF EXISTS vector;
