CREATE TABLE 'categories' (
    'id' varchar(36) NOT NULL,
    'name' varchar(255) NOT NULL,
    primary key (id)
);

CREATE TABLE 'products' (
    'id' varchar(36) NOT NULL,
    'name' varchar(255) NOT NULL,
    'description' varchar(255) NOT NULL,
    'price' decimal(10,2) NOT NULL,
    'category_id' varchar(36) NOT NULL,
    'image_url' varchar(255) NOT NULL,
    PRIMARY KEY ('ID'),
    KEY 'fk_products_categories_idx' ('category_id')
);