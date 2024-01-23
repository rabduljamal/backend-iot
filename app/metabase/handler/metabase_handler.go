package handler

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rabduljamal/backend-iot/app/metabase/repository"
	"github.com/rabduljamal/backend-iot/lib"
)

func GetMetabases(c *fiber.Ctx) error {

	requestBody := c.Body()

	// Mencetak body input ke konsol
	fmt.Println("Request Body:", string(requestBody))
	dataInput := new(repository.MetabaseParam)
	if err := c.BodyParser(dataInput); err != nil {
		return err // Handle parsing error
	}

	data, err := repository.GetMetabaseData(dataInput)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error processing request")
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error processing request")
	}

	var timestamp = time.Now().Format("20060102150405")

	responseData, _ := lib.Success(string(jsonData), timestamp)
	return c.Status(200).JSON(responseData)
}
