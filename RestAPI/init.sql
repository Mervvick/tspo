DROP TABLE IF EXISTS categories;
CREATE TABLE categories (
    id BIGINT PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

DROP TABLE IF EXISTS products;
CREATE TABLE products (
    id BIGINT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    category_id INT NOT NULL,
	price INT,
    FOREIGN KEY (category_id) REFERENCES categories(id)
);

INSERT INTO categories (id, name) VALUES
(1, 'Солнцезащитные очки'),
(2, 'Контактные линзы'),
(3, 'Оправы для очков'),
(4, 'Очки для чтения'),
(5, 'Защитные очки');

INSERT INTO products (id, name, description, category_id, price)
VALUES
(1, 'Солнцезащитные очки Ray-Ban Aviator', 'Классические авиаторы с металлической оправой и темными линзами.', 1, 100000),
(2, 'Контактные линзы Acuvue Oasys', 'Дышащие двухнедельные контактные линзы для комфортного ношения.', 2, 5000),
(3, 'Оправа для очков Gucci GG0010O', 'Стильная оправа из ацетата с фирменным дизайном Gucci.', 3, 12000),
(4, 'Очки для чтения Foster Grant', 'Легкие и удобные очки для чтения с различными диоптриями.', 4, 15000),
(5, 'Солнцезащитные очки Oakley Holbrook', 'Спортивные очки с высокой ударопрочностью и защитой от УФ-лучей.', 1, 6000),
(6, 'Контактные линзы Biofinity', 'Месячные силикон-гидрогелевые линзы с высоким уровнем увлажнения.', 2, 10000),
(7, 'Оправа для очков Ray-Ban Wayfarer', 'Иконическая пластиковая оправа в стиле ретро.', 3, 10000),
(8, 'Компьютерные очки Gunnar Optiks', 'Очки с фильтром синего света для защиты глаз при работе за компьютером.', 5, 7432),
(9, 'Контактные линзы Dailies Total 1', 'Ежедневные одноразовые линзы с уникальной технологией увлажнения.', 2, 13243),
(10, 'Солнцезащитные очки Polaroid PLD 1013/S', 'Очки с поляризованными линзами для четкого и контрастного зрения.', 1, 9000);
