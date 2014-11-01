package web

import (
	"flag"
	"github.com/gorilla/schema"
	"github.com/zenazn/goji"
	"net/http"
)

var (
	defaults *config
	decoder  = schema.NewDecoder()
)

func start(conf *config) {
	defaults = conf
	flag.Set("bind", conf.Address) // Uh, I guess that's a bit strange

	register("/head/:player", serveHeadNormal)
	register("/head/:size/:player", serveHeadWithSize)

	register("/face/:player", serveFaceNormal)
	register("/face/:size/:player", serveFaceWithSize)

	goji.Get("/*", http.FileServer(http.Dir("www"))) // TODO: How to find the correct dir?

	goji.Serve()
}

func register(pattern string, handler interface{}) {
	goji.Get(pattern, handler)
	goji.Get(pattern+".png", handler)
}