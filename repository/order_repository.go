package repository

import (
	"log"

	"github.com/mogueno/burrito-shop/models"
	"github.com/mogueno/burrito-shop/repository/dbqueries"
	database "github.com/mogueno/burrito-shop/utils"
)

func GetOrders() []models.Order {
	rows, err := database.DB.Query(dbqueries.GET_ORDERS)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		err := rows.Scan(&order.ID, &order.TotalCost)
		if err != nil {
			log.Fatal(err)
		}

		order.Items = GetOrderItems(order.ID)
		orders = append(orders, order)
	}

	return orders
}

func GetOrderItems(orderID uint) []models.OrderItem {
	rows, err := database.DB.Query(dbqueries.GET_ORDER_ITEMS, orderID)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var orderItems []models.OrderItem
	for rows.Next() {
		var orderItem models.OrderItem
		err := rows.Scan(&orderItem.ID, &orderItem.Burrito.ID, &orderItem.Quantity)
		if err != nil {
			log.Fatal(err)
		}

		orderItem.Burrito = GetBurrito(orderItem.Burrito.ID)

		orderItems = append(orderItems, orderItem)
	}

	return orderItems
}

func GetOrder(orderID uint) models.Order {
	var order models.Order
	row := database.DB.QueryRow(dbqueries.GET_ORDER, orderID)
	err := row.Scan(&order.ID, &order.TotalCost)
	if err != nil {
		log.Fatal(err)
	}


	order.Items = GetOrderItems(order.ID)
	return order
}

func SaveOrder(order models.Order) error {
	// Prepare the SQL statement for inserting the order into the database
	stmt, err := database.DB.Prepare(dbqueries.SAVE_ORDER)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement to insert the order
	res, err := stmt.Exec(order.TotalCost)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Retrieve the ID of the last inserted order
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Insert the order items into the database
	for _, item := range order.Items {
		burritoID, err := SaveBurrito(item.Burrito)
		if err != nil {
			log.Fatal(err)
			return err
		}
		item.Burrito.ID = uint(burritoID)
		err = saveOrderItem(lastInsertID, item)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	return nil
}

func saveOrderItem(orderID int64, item models.OrderItem) error {
	// Prepare the SQL statement for inserting the order item into the database
	stmt, err := database.DB.Prepare(dbqueries.SAVE_ORDER_ITEM)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement to insert the order item
	_, err = stmt.Exec(orderID, item.Burrito.ID, item.Quantity)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}