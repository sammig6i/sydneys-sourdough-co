CREATE VIEW IF NOT EXISTS unified_search AS
SELECT 
    'menu_item' as entity_type,
    id,
    name,
    description as context,
    embedding
FROM menu_items
WHERE embedding IS NOT NULL
UNION ALL
SELECT 
    'category' as entity_type,
    id,
    name,
    NULL as context,
    embedding
FROM categories
WHERE embedding IS NOT NULL;
