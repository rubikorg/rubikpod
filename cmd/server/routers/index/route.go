package index

import (
	r "github.com/rubikorg/rubik"
)

// Router is index's router
var Router = r.Create("/")

func init() {
	loginRoute := r.Route{
		Path:       "/login",
		Controller: loginCtl,
	}
	Router.Add(loginRoute)
}
