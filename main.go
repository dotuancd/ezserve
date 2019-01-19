package main

import (
	. "github.com/dotuancd/ezserve/app"
	//"database/sql"
	"fmt"
	//"time"
)

func main() {
	a := DefaultApp()

	bootApp(a)
	defer shutdown(a)

	//registerRoutes(a)

	err := a.Routes.Run(":8000")
	//err := http.ListenAndServe(":8000", a.Routes)

	if err != nil {
		panic(fmt.Errorf("[ERROR] - Cannot start the application %s", err))
	}
}

func shutdown(app *App) {
	app.DB.Close()
}
