package index

import (
	r "github.com/rubikorg/rubik"
)

func loginCtl(req *r.Request) {
	req.Respond("Hello, world", r.Type.HTML)
}
