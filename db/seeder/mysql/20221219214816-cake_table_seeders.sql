-- +migrate Up
INSERT INTO
    cakes (id, title, description, image, rating, created_at, updated_at)
VALUES
    (1, 'Carrot and walnut cake', 'Carrot gold — for a vegetable that’s simply tops, you can’t go past this trusty rabbit-magnet!', 'http://www.taste.com.au/recipes/23320/carrot+and+walnut+cake?ref=,', 10, NOW(), NOW()),
    (2, 'New York Baked Cheese Cake', 'Take the cake and master the basics with this creamy baked version of a classic dessert.', 'http://www.taste.com.au/recipes/27853/new+york+baked+cheesecake?ref=,', 8, NOW(), NOW()),
    (3, 'Chocolate coconut cake', 'A hint of coconut takes this moist chocolate cake to a whole new level.', 'http://www.taste.com.au/recipes/21064/chocolate+coconut+cake?ref=,', 9, NOW(), NOW())
;

-- +migrate Down
TRUNCATE cakes;
