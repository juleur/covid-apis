package main

import (
	covidfetcher "backend-covid-api/covid-fetch"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

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

	app.Get("/geojson_departements", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json")
		c.Set("Access-Control-Allow-Origin", "https://www.coro21-jl.xyz")
		return c.Send(geojson)
	})

	app.Get("/covid-data", func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "https://www.coro21-jl.xyz")
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
