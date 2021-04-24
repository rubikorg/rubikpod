package main

import (
	r "github.com/rubikorg/rubik"
	"rubikpod/cmd/server/app"
	"rubikpod/cmd/server/routers"
)

func main() {
	var config app.ProjectConfig
	err := r.Load(&config)

	if err != nil {
		panic(err)
	}

	// TODO: set your one-time application level dependency here
	// eg: DB Connection, Logger etc..
	err = app.SetDep(app.Dependency{
		ProjectConfig: config,
	})
	if err != nil {
		panic(err)
	}

	routers.Import()
	panic(r.Run())

}
