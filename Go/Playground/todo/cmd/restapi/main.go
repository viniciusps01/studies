package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/viniciusps01/cmd/app"
	"github.com/viniciusps01/internal/ui/restapi/handler"
	"github.com/viniciusps01/internal/ui/restapi/middleware"
	"github.com/viniciusps01/internal/ui/restapi/routes"
)

func main() {
	app := app.New()

	handler.SetUp(app)
	middleware.SetUp(app)

	defer app.Conn.Close()

	fmt.Printf("Server is up on %v:%d\n",
		app.Env.Server.Host,
		app.Env.Server.Port,
	)

	err := http.ListenAndServe(app.Env.Server.Address(), routes.All())

	if err != nil {
		log.Fatal(err)
	}

}
