package middleware

import (
	"crypto/rand"
	"crypto/rsa"
	"log"

	jwtware "github.com/gofiber/contrib/jwt"

	"github.com/gofiber/fiber/v2"
)

var (
	// Obviously, this is just a test example. Do not do this in production.
	// In production, you would have the private key and public key pair generated
	// in advance. NEVER add a private key to any GitHub repo.
	privateKey *rsa.PrivateKey
)

var apiJwtMiddleware = jwtware.New(jwtware.Config{
	SigningKey: jwtware.SigningKey{
		JWTAlg: jwtware.RS256,
		Key:    getPrivateKey().Public(),
	},
	ErrorHandler: apiJwtError,
})

var consoleJwtMiddleware = jwtware.New(jwtware.Config{
	SigningKey: jwtware.SigningKey{
		JWTAlg: jwtware.RS256,
		Key:    getPrivateKey().Public(),
	},
	ErrorHandler: consoleJwtError,
})

func getPrivateKey() *rsa.PrivateKey {
	if privateKey == nil {
		rng := rand.Reader
		var err error
		privateKey, err = rsa.GenerateKey(rng, 2048)
		if err != nil {
			log.Fatalf("rsa.GenerateKey: %v", err)
		}
	}
	return privateKey
}

func ApiProtected(c *fiber.Ctx) error {

	return apiJwtMiddleware(c)
}

func ConsoleProtected(c *fiber.Ctx) error {

	return consoleJwtMiddleware(c)
}

func apiJwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}

func consoleJwtError(c *fiber.Ctx, err error) error {
	return c.Redirect("/console/login")
}
