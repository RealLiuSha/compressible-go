package compressible_test

import (
	"fmt"

	"github.com/itchenyi/compressible-go"
	"github.com/itchenyi/akita"
	"github.com/itchenyi/akita/middleware/static"
)

func Example() {
	fmt.Println(compressible.Is("text/html"))
	// -> true

	fmt.Println(compressible.Is("image/png"))
	// -> false

	var wt compressible.WithThreshold = 1024

	fmt.Println(wt.Compressible("text/html", 1023))
	// -> false
}

func ExampleWithAkita() {
	app := akita.New()
	app.Set("AppCompress", compressible.WithThreshold(1024))

	// Add a static middleware
	app.Use(static.New(static.Options{
		Root:   "./",
		Prefix: "/",
	}))
	app.Error(app.Listen(":3000")) // http://127.0.0.1:3000/
}
