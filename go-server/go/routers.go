/*
 * Product aggregation composite interface
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/MyTeam889/Life/0.0.1/",
		Index,
	},

	Route{
		"DeleteSubscription",
		strings.ToUpper("Delete"),
		"/MyTeam889/Life/0.0.1/API/v1/subscriptions/{serviceid}",
		DeleteSubscription,
	},

	Route{
		"GetOfferss",
		strings.ToUpper("Get"),
		"/MyTeam889/Life/0.0.1/API/v1/offerslist",
		GetOfferss,
	},

	Route{
		"GetSubscriptions",
		strings.ToUpper("Get"),
		"/MyTeam889/Life/0.0.1/API/v1/subscriptionslist",
		GetSubscriptions,
	},

	Route{
		"PostSubscription",
		strings.ToUpper("Post"),
		"/MyTeam889/Life/0.0.1/API/v1/subscriptions/{serviceid}",
		PostSubscription,
	},
}
