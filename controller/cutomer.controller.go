package controller

import (
	"github.com/gofiber/fiber"
	"github.com/indrabpn12/laundry-api-golang.git/database"
	"github.com/indrabpn12/laundry-api-golang.git/model"
)

func GetCustomerAll(c *fiber.Ctx) error {
	var customers []model.Customer

	database.DB.Find(&customers)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "get all customer",
		"data":    customers,
	})
}

func GetCustomer(c *fiber.Ctx) error {
	id := c.Params("id")

	var customer model.Customer

	database.DB.Find(&customer, id)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "get customer by id",
		"data":    customer,
	})
}

func CreateCustomer(c *fiber.Ctx) error {
	customer := new(model.Customer)

	if err := c.BodyParser(customer); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "cannot create customer",
			"data":    err,
		})
	}

	database.DB.Create(&customer)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "create customer",
		"data":    customer,
	})
}

func UpdateCustomer(c *fiber.Ctx) error {
	id := c.Params("id")

	customer := new(model.Customer)

	if err := c.BodyParser(customer); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "cannot update customer",
			"data":    err,
		})
	}

	database.DB.Model(&customer).Where("id = ?", id).Updates(customer)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "update customer",
		"data":    customer,
	})
}

func DeleteCustomer(c *fiber.Ctx) error {
	id := c.Params("id")

	var customer model.Customer

	database.DB.Delete(&customer, id)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "delete customer",
	})
}
