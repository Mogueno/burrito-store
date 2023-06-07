package dbqueries

const (
	GET_BURRITOS = "SELECT id, name, size, price FROM burritos"
	SAVE_BURRITO = "INSERT INTO burritos (name, size, price) VALUES (?, ?, ?)"
	GET_BURRITO = "SELECT id, name, size, price FROM burritos WHERE id = ?"
	GET_ORDERS = "SELECT id, total_cost FROM orders"
	GET_ORDER = "SELECT id, total_cost FROM orders WHERE id = ?"
	SAVE_ORDER = "INSERT INTO orders (total_cost) VALUES (?)"
	GET_ORDER_ITEMS = "SELECT id, burrito_id, quantity FROM order_items WHERE order_id = ?"
	SAVE_ORDER_ITEM = "INSERT INTO order_items (order_id, burrito_id, quantity) VALUES (?, ?, ?)"
)