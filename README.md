# compressible-go

Compressible Content-Type / mime checking for Go.

## Demo

```go
import (
  "github.com/itchenyi/compressible-go"
)

fmt.Println(compressible.Is("text/html"))
// -> true

fmt.Println(compressible.Is("image/png"))
// -> false

var wt compressible.WithTrashold = 1024

fmt.Println(wt.Compressible("text/html", 1023))
// -> false
```

**Work with akita:**

```go
package main

import (
  "github.com/itchenyi/compressible-go"
  "github.com/itchenyi/akita"
  "github.com/itchenyi/akita/middleware/static"
)

func main() {
  app := akita.New()
  app.Set(akita.SetCompress, compressible.WithThreshold(1024))

  // Add a static middleware
  app.Use(static.New(static.Options{
    Root:   "./",
    Prefix: "/",
  }))
  app.Error(app.Listen(":3000")) // http://127.0.0.1:3000/
}
```

## License

MIT
