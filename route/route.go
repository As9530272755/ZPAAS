package route

import (
	"exec/controllers"
	"net/http"
)

func Register() {
	http.HandleFunc("/", controllers.GetUser)
	http.HandleFunc("/create/", controllers.AddUser)
	http.HandleFunc("/delete/", controllers.DeleteUser)
	http.HandleFunc("/edit/", controllers.Edit)
}
