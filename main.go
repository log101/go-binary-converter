package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// convert binary to integer
func binaryToInt(binary string) int {
	// create a variable to hold the integer value
	var integer int

	// loop through the binary string
	for i := 0; i < len(binary); i++ {
		// convert the binary string to integer
		integer += int(binary[i]-'0') << uint(len(binary)-i-1)
	}

	// return the integer value
	return integer
}

// convert binary string to ascii
func binaryToAscii(binary string) []byte {
	// create a byte slice to hold the ascii values
	ascii := make([]byte, len(binary)/8)

	// loop through the binary string
	for i := 0; i < len(binary); i += 8 {
		// convert the binary string to ascii
		ascii[i/8] = byte(binaryToInt(binary[i : i+8]))
	}

	// return the ascii string
	return ascii
}

func main() {
	app := fiber.New()

	app.Static("/", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// add a new route to convert binary to ascii
	app.Get("/ascii", func(c *fiber.Ctx) error {
		// get the binary string from the query string
		binary := c.Query("binary")

		// convert the binary string to ascii
		ascii := string(binaryToAscii(binary))

		// return the ascii string
		return c.SendString(ascii)
	})

	// add a new route to convert ascii to binary
	app.Get("/binary", func(c *fiber.Ctx) error {
		// get the ascii string from the query string
		ascii := c.Query("ascii")

		// create a byte slice to hold the binary string
		binary := make([]byte, 0)

		// loop through the ascii string
		for i := 0; i < len(ascii); i++ {
			// convert the ascii string to binary
			binary = append(binary, []byte(fmt.Sprintf("%08b", ascii[i]))...)
		}

		// return the binary string
		return c.SendString(string(binary))
	})

	app.Listen(":3000")
}
