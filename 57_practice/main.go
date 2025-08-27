package main

import (
	"log"
	backtobasic "master_go_programming/57_practice/backTobasic"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	// SampleJson "practice/json"
)

func init() {
	// Initialize the database
	// db.ConnectDB()
}

func main() {
	// // Create an instance of Gotchas
	// gotchas := diffuc.Gotchas{}
	// // Pass the instance to ExampleGotchas
	// diffuc.ExampleGotchas(gotchas)
	// diffuc.SliceArrays()
	// diffuc.StandardLib()
	// diffuc.Print()
	// diffuc.LoopingFor()
	// diffuc.Booleany()
	// diffuc.SwitchCase()
	// diffuc.ExampleFunction()
	// diffuc.VariadicFunction()
	// diffuc.ArrayExample()
	// diffuc.ExampleSlice()
	// diffuc.ExampleMaps()
	// diffuc.ExampleStruct()
	// diffuc.ExampleInterface()
	// diffuc.ExampleGoroutines()
	// diffuc.ExampleChannel()
	// diffuc.ExampleLog()
	// diffuc.ExampleFileDirectory()
	// diffuc.ReadDifferentFile()
	// diffuc.ExampleRegex()
	// diffuc.ExampleFindDNS()
	// diffuc.ExampleCryptography()
	// diffuc.SampleReflection()

	// photogallery.PhotoGallery()
	// barcode.ExampleBarcode()
	// SampleJson.SampleJson()

	// beginer.SamplePanic()
	// beginer.SampleCheckStruct()
	// beginer.SampleInitializeStruct()
	// beginer.SampleEmptyInterface()
	// beginer.ExampleRecoverPanic()
	// beginer.StructToMapString()
	// beginer.ExampleCreateSlice()
	// beginer.SampleMapstructAppend()
	// StringManipulation.StringManipulate()
	// methodobjects.EntryObjectMethod()
	// sampleInterface.SampleInterface()
	// backtobasic.SampleFunction()
	backtobasic.SampleCheckMap()

	// Use the global database instance
	// err := db.GlobalDB.DB.AutoMigrate(&model.Customer{})
	// if err != nil {
	// 	log.Fatalf("Failed to migrate database: %v", err)
	// }

	// Initialize gofiber
	app := fiber.New(fiber.Config{
		CaseSensitive:    true,
		DisableKeepalive: true,
	})

	// CORS configuration
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET,POST,PUT,DELETE"},
		AllowHeaders: []string{"Origin, Content-Type, Accept, Authorization"},
	}))

	// Apply rate limiting middleware
	// app.Use(limiter.New(limiter.Config{
	// 	Max:        5, // Maximum 5 requests
	// 	Expiration: 1 * time.Minute,
	// 	KeyGenerator: func(c fiber.Ctx) string {
	// 		return c.IP() // Rate limit by IP
	// 	},
	// 	LimitReached: func(c fiber.Ctx) error {
	// 		return helper.JSONResponse(c, respcode.ERR_CODE_405, "Too many requests, please try again later.")
	// 	},
	// }))

	app.Use(logger.New())
	app.Use(recover.New())
	log.Fatal(app.Listen(":8007"))
}
