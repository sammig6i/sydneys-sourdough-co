INSERT INTO categories (name) VALUES
  ('Breads'),
  ('Pastries'),
  ('Sandwiches'),
  ('Beverages');

INSERT INTO menu_items (name, description, price, category_id, image_url) VALUES
  (
    'Classic Sourdough',
    'Traditional San Francisco style sourdough with a perfectly crispy crust',
    8.99,
    (SELECT id FROM categories WHERE name = 'Breads'),
    'https://images.unsplash.com/photo-1586444248902-2f64eddc13df?auto=format&fit=crop&w=800&q=80'
  ),
  (
    'Butter Croissant',
    'Flaky, buttery layers with a golden-brown exterior',
    4.99,
    (SELECT id FROM categories WHERE name = 'Pastries'),
    'https://images.unsplash.com/photo-1555507036-ab1f4038808a?auto=format&fit=crop&w=800&q=80'
  ),
  (
    'Turkey Avocado Club',
    'Roasted turkey, fresh avocado, crispy bacon on fresh sourdough',
    12.99,
    (SELECT id FROM categories WHERE name = 'Sandwiches'),
    'https://images.unsplash.com/photo-1567234669003-dce7a7a88821?auto=format&fit=crop&w=800&q=80'
  ),
  (
    'Artisan Coffee',
    'House-roasted specialty coffee blend',
    3.99,
    (SELECT id FROM categories WHERE name = 'Beverages'),
    'https://images.unsplash.com/photo-1509785307050-d4066910ec1e?auto=format&fit=crop&w=800&q=80'
  );

      