package generator

import (
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/ranjeetk489/ecom-service/internal/models"
)

func GenerateFakeOrder() models.Order {
	quantity, _ := faker.RandomInt(1, 10)
	return models.Order{
		ID:         faker.UUIDDigit(),
		CustomerID: faker.UUIDDigit(),
		ProductID:  faker.UUIDDigit(),
		Quantity:   quantity[0],
		TotalPrice: faker.Currency(),
		OrderDate:  time.Now(),
	}
}
