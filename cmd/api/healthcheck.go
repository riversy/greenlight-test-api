package main

import (
	"fmt"
	"net/http"
)

// healthcheckHandler responds for a healthcheck
// request if the server is still alive
func (app application) healthcheckHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "status: available")
	fmt.Fprintf(writer, "environment: %s\n", app.config.env)
	fmt.Fprintf(writer, "version: %s\n", version)
}
