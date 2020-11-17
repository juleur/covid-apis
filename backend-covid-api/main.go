package main

import (
	covidfetcher "backend-covid-api/covid-fetch"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gofiber/cors"
	fiber "github.com/gofiber/fiber/v2"
)

var (
	portArg uint64
)

func init() {
	flag.Uint64Var(&portArg, "port", 1029, "server port")
	flag.Parse()
}

func main() {
	geojson := openGeoJSON()
	covidfetcher.CovidDataBackup()
	covidStore := covidfetcher.CovidStore

	app := fiber.New(fiber.Config{
		GETOnly: true,
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://www.coro21-jl.xyz"},
		AllowHeaders: []string{"Origin, Content-Type, Accept"},
	}))

	app.Get("/geojson_departements", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json")
		return c.Send(geojson)
	})

	app.Get("/covid-data", func(c *fiber.Ctx) error {
		return c.JSON(covidStore.LastSevenDataDays)
	})

	log.Fatalln(app.Listen(fmt.Sprintf(":%d", portArg)))
}

func openGeoJSON() []byte {
	geo, err := ioutil.ReadFile("departements.geojson")
	if err != nil {
		log.Fatal(err)
	}
	return geo
}
