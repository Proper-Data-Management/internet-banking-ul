package server

import (
	"database/sql"
	"time"

	jsoniter "github.com/json-iterator/go"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	companyPersonHandlers "github.com/internet-banking-ul/internal/handlers/company_person"
	customerHandlers "github.com/internet-banking-ul/internal/handlers/customer"
	"github.com/internet-banking-ul/internal/middles"
	companyPersonService "github.com/internet-banking-ul/internal/modules/company_person/services"
	customerService "github.com/internet-banking-ul/internal/modules/customer/services"
	"github.com/internet-banking-ul/modules/logger"
)

//NewServer all rest api
func NewServer(db *sql.DB) *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "AL-HILAL-CORE",
		Immutable:     true,
		JSONEncoder: func(v interface{}) ([]byte, error) {
			return jsoniter.ConfigFastest.Marshal(v)
		},
	})

	app.Use(pprof.New())

	app.Use(etag.New(etag.Config{
		Weak: false,
		Next: nil,
	}))

	app.Use(middles.Logger(3*time.Second, logger.WorkLogger))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // Compress everything
	}))

	v1 := app.Group("/api/v1")
	customerHandlers.NewCustomerHandler(customerService.NewCustomerService(db)).RegisterCustomer(v1)
	companyPersonHandlers.NewCompanyPersonHandler(companyPersonService.NewCompanyPersonService(db)).RegisterCompanyPerson(v1)

	return app
}
