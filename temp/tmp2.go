package main

import "fmt"

type iOrder interface {
	GetID() string
}

type Order struct {
	ID string
}

func (o *Order) GetID() string {
	return o.ID

}

func main() {
	order := createOrder()

	fmt.Println(order == nil)
	fmt.Printf("ORDER: %v\n", order)
	fmt.Printf("ORDER ID: %v\n", order.GetID())
}

func createOrder() iOrder {
	var order *Order

	return order
}
