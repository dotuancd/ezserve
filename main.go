package main

import (
	. "bitbucket.org/ezserve/ezserve/app"
	//"database/sql"
	"fmt"
	"net/http"
	//"time"
)

func main() {
	a := NewApp()

	bootApp(a)
	defer shutdown(a)

	registerFileRoutes(a)

	err := http.ListenAndServe(":8000", a.Routes)

	if err != nil {
		panic(fmt.Errorf("[ERROR] - Cannot start the application %s", err))
	}
}

func shutdown(app *App) {
	app.DB.Close()
}
