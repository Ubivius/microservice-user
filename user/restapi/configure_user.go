// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	"Ubivius/microservices/user-microservice/user/models"
	"Ubivius/microservices/user-microservice/user/restapi/operations"
	"Ubivius/microservices/user-microservice/user/restapi/operations/users"
)

//go:generate swagger generate server --target ..\..\user --name User --spec ..\swagger.yml --principal interface{}

var exampleFlags = struct {
	Example1 string `long:"example1" description:"Sample for showing how to configure cmd-line flags"`
	Example2 string `long:"example2" description:"Further info at https://github.com/jessevdk/go-flags"`
}{}

func configureFlags(api *operations.UserAPI) {
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		swag.CommandLineOptionsGroup{
			ShortDescription: "Example Flags",
			LongDescription:  "",
			Options:          &exampleFlags,
		},
	}
}

var items = make(map[int64]*models.User)
var lastID int64

var itemsLock = &sync.Mutex{}

func newItemID() int64 {
	return atomic.AddInt64(&lastID, 1)
}

func addItem(item *models.User) error {
	if item == nil {
		return errors.New(500, "Please specify a user")
	}

	itemsLock.Lock()
	defer itemsLock.Unlock()

	newID := newItemID()
	item.ID = newID
	items[newID] = item

	return nil
}

func allItems() (result []*models.User) {
	result = make([]*models.User, 0)
	for _, item := range items {
		result = append(result, item)
	}
	return
}

func configureAPI(api *operations.UserAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.UsersAddUserHandler = users.AddUserHandlerFunc(func(params users.AddUserParams) middleware.Responder {
		if err := addItem(params.Body); err != nil {
			return users.NewAddUserDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		}
		return users.NewAddUserCreated().WithPayload(params.Body)
	})

	api.UsersGetUsersHandler = users.GetUsersHandlerFunc(func(params users.GetUsersParams) middleware.Responder {
		return users.NewGetUsersOK().WithPayload(allItems())
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}
	println(exampleFlags.Example1)
	println(exampleFlags.Example2)

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
	if exampleFlags.Example1 != "something" {
		fmt.Print("example1 argument is not something")
	}
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
