package middleware

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rabduljamal/backend-iot/lib"
)

func DecryptRequestMiddleware(c *fiber.Ctx) error {
	// Mendapatkan data permintaan dari c.Body
	requestBody := c.Body()

	// Mendekripsi permintaan untuk mengambil objek JSON
	requestData := map[string]interface{}{}
	if err := json.Unmarshal([]byte(requestBody), &requestData); err != nil {
		return err
	}

	// Mendekripsi bagian "data" dalam objek JSON
	encryptedData := requestData["data"].(string)

	timestamp := c.Get("X-Snip-Timestamp")
	decryptedData, err := lib.DecryptAES(encryptedData, timestamp)
	if err != nil {
		return err
	}

	fmt.Println(string(decryptedData))

	var decryptedRequest map[string]interface{}

	if err := json.Unmarshal(decryptedData, &decryptedRequest); err != nil {
		return err
	}

	fmt.Println(decryptedRequest)

	// Mengganti c.Body dengan data yang sudah didekripsi
	c.Request().Header.Set("Content-Type", "application/json")
	c.Request().SetBody(decryptedRequest) // Mengganti body dengan data yang sudah didekripsi

	return c.Next()
}
