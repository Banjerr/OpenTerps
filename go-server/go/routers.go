/*
 * OpenTerps
 *
 * An Open API that contains information about terpenes, the effects, and the cannabis varieties that contain them.
 *
 * API version: 0.0.2
 * Contact: benjamminredden@gmail.com
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
		"/CountryFriedCoders/OpenTerps/0.0.1/",
		Index,
	},

	Route{
		"AddCannabisStrain",
		strings.ToUpper("Post"),
		"/CountryFriedCoders/OpenTerps/0.0.1/cannabis",
		AddCannabisStrain,
	},

	Route{
		"DeleteCannabisStrain",
		strings.ToUpper("Delete"),
		"/CountryFriedCoders/OpenTerps/0.0.1/cannabis/{cannabisId}",
		DeleteCannabisStrain,
	},

	Route{
		"FindCannabisStrainsByTerpene",
		strings.ToUpper("Get"),
		"/CountryFriedCoders/OpenTerps/0.0.1/cannabis/findByTerpene",
		FindCannabisStrainsByTerpene,
	},

	Route{
		"GetCannabisStrainById",
		strings.ToUpper("Get"),
		"/CountryFriedCoders/OpenTerps/0.0.1/cannabis/{cannabisId}",
		GetCannabisStrainById,
	},

	Route{
		"UpdateCannabisStrain",
		strings.ToUpper("Put"),
		"/CountryFriedCoders/OpenTerps/0.0.1/cannabis",
		UpdateCannabisStrain,
	},

	Route{
		"UpdateCannabisStrainByID",
		strings.ToUpper("Post"),
		"/CountryFriedCoders/OpenTerps/0.0.1/cannabis/{cannabisId}",
		UpdateCannabisStrainByID,
	},

	Route{
		"AddEffect",
		strings.ToUpper("Post"),
		"/CountryFriedCoders/OpenTerps/0.0.1/effect",
		AddEffect,
	},

	Route{
		"DeleteEffect",
		strings.ToUpper("Delete"),
		"/CountryFriedCoders/OpenTerps/0.0.1/effect/{effectId}",
		DeleteEffect,
	},

	Route{
		"GetEffectById",
		strings.ToUpper("Get"),
		"/CountryFriedCoders/OpenTerps/0.0.1/effect/{effectId}",
		GetEffectById,
	},

	Route{
		"UpdateEffect",
		strings.ToUpper("Put"),
		"/CountryFriedCoders/OpenTerps/0.0.1/effect",
		UpdateEffect,
	},

	Route{
		"UpdateEffectByID",
		strings.ToUpper("Post"),
		"/CountryFriedCoders/OpenTerps/0.0.1/effect/{effectId}",
		UpdateEffectByID,
	},

	Route{
		"AddTerpene",
		strings.ToUpper("Post"),
		"/CountryFriedCoders/OpenTerps/0.0.1/terpene",
		AddTerpene,
	},

	Route{
		"DeleteTerpene",
		strings.ToUpper("Delete"),
		"/CountryFriedCoders/OpenTerps/0.0.1/terpene/{terpeneId}",
		DeleteTerpene,
	},

	Route{
		"FindTerpenesByEffect",
		strings.ToUpper("Get"),
		"/CountryFriedCoders/OpenTerps/0.0.1/terpene/findByEffect",
		FindTerpenesByEffect,
	},

	Route{
		"GetTerpeneById",
		strings.ToUpper("Get"),
		"/CountryFriedCoders/OpenTerps/0.0.1/terpene/{terpeneId}",
		GetTerpeneById,
	},

	Route{
		"UpdateTerpene",
		strings.ToUpper("Put"),
		"/CountryFriedCoders/OpenTerps/0.0.1/terpene",
		UpdateTerpene,
	},

	Route{
		"UpdateTerpeneByID",
		strings.ToUpper("Post"),
		"/CountryFriedCoders/OpenTerps/0.0.1/terpene/{terpeneId}",
		UpdateTerpeneByID,
	},

	Route{
		"CreateUser",
		strings.ToUpper("Post"),
		"/CountryFriedCoders/OpenTerps/0.0.1/user",
		CreateUser,
	},

	Route{
		"CreateUsersWithArrayInput",
		strings.ToUpper("Post"),
		"/CountryFriedCoders/OpenTerps/0.0.1/user/createWithArray",
		CreateUsersWithArrayInput,
	},
}
