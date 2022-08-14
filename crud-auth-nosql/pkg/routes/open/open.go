package open

import (
	"crud-auth-nosql/pkg/controllers"

	"github.com/gorilla/mux"
)

/*
Following routes can be accessed on HTTP and HTTPS Schemes
*/
var RegisterStatsRoutes = func(r mux.Router) {
	// Returns the total number of users signed up and names and pics of public profiles
	r.HandleFunc("/user", controllers.UserStats).Methods("GET")
	// Returns the statistics surrounding Todos
	r.HandleFunc("/todo", controllers.TodoStats).Methods("GET")
}

var RegisterBrowseRoutes = func(r mux.Router) {
	// Fetches the User profile's public elements if profile is public
	r.HandleFunc("/user/{id}", controllers.GetUserPublicProfile).Methods("GET")
	// Fetches the User's public todos
	r.HandleFunc("/user/{userId}/todo", controllers.GetUserPublicTodos).Methods("GET")
	// Fetches the User's public todo details
	r.HandleFunc("/user/{userId}/todo/{todoId}", controllers.GetUserPublicTodoDetails).Methods("GET")
}
