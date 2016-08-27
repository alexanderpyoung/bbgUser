package main

import (
	"github.com/alexanderpyoung/bbgUser/api"
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", api.Handlers())
}
