-- --------------------------------------------
-- Table structure for `cart`
-- --------------------------------------------
CREATE TABLE IF NOT EXISTS cart (
    user_id UUID NOT NULL,
    item_id INT NOT NULL,
    name VARCHAR(100),
    image_url VARCHAR(255),
    quantity INT NOT NULL
);

-- --------------------------------------------
-- Table structure for `items`
-- --------------------------------------------
CREATE TABLE IF NOT EXISTS items (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    price DECIMAL(10, 2),
    image_url VARCHAR(255)
);

-- Insert data into `items`
INSERT INTO items (id, name, price, image_url) VALUES
    (1, 'LAYS sour_cream&onion', 100.00, 'img/lays.jpg'),
    (2, 'Coca-Cola', 150.00, 'img/cola.jpg'),
    (3, 'Fanta', 120.00, 'img/orange.jpg');

-- --------------------------------------------
-- Table structure for `orders`
-- --------------------------------------------
CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    name VARCHAR(100),
    phone VARCHAR(20),
    address VARCHAR(255) NOT NULL,
    comment TEXT,
    promo_code VARCHAR(50),
    items JSONB DEFAULT NULL,
    is_delivered INT DEFAULT 0,
    delivery_time INT DEFAULT 5
);

-- --------------------------------------------
-- Table structure for `user_profile`
-- --------------------------------------------
CREATE TABLE IF NOT EXISTS user_profile (
    user_id UUID PRIMARY KEY,
    name VARCHAR(100),
    address VARCHAR(255) NOT NULL,
    phone VARCHAR(20)
);
