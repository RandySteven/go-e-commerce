package gosql

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
)
