package gosql

//Product Query
const (
	ProductPaginationFunctionQuery = `
	CREATE OR REPLACE FUNCTION GetProductPage(page_size INT, last_id INT)
	RETURNS SETOF products AS $$
	BEGIN
		RETURN QUERY
		SELECT * FROM products WHERE id > last_id
		ORDER BY id
		LIMIT page_size;
	
		RETURN;
	END;
	$$ LANGUAGE plpgsql;
	`

	SelectProducts = `SELECT 
		id, name, price, stock, description, shop_id, created_at, updated_at 
		FROM products
		`

	SelectProductById = `
		SELECT * FROM products 
		WHERE id = $1
	`

	InsertProductQuery = `
		INSERT INTO products 
			(name, price, stock, description, shop_id, category_id) 
			VALUES 
			($1, $2, $3, $4, $5, $6) 
			RETURNING ID
	`
)

//User Query
const (
	InsertUserQuery = `
		INSERT INTO users 
		(name, email, password, birthday, phone_number) 
		VALUES 
		($1, $2, $3, $4, $5) 
		RETURNING ID
	`

	SelectUserById = `
	SELECT * FROM users WHERE id = $1
	`

	SelectUserByEmail = `
		SELECT id, name, email, birthday, phone_number 
		FROM users 
		WHERE email = $1
	`
)

//Category Query
const (
	InsertCategoryQuery = `
		INSERT INTO categories (category)
		VALUES
		($1) RETURNING ID
	`

	SelectCategories = `
		SELECT id, category, created_at, updated_at FROM categories
	`

	SelectCategoryById = `
		SELECT category FROM categories WHERE id = $1
	`

	SelectCategoryDetail = `
		SELECT c.category, p.name, p.price, p.stock, p.description
		FROM
		categories c JOIN products p
		ON
		c.id = p.category_id
		WHERE c.id = $1
	`
)
