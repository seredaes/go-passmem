package Web

import (
	"fmt"
	"log"
	"net/http"
	"time"

	AuthController "seredaes/go-passmem/App/Controllers/AuthController"
	"seredaes/go-passmem/App/Controllers/CredentialController"
	RegistrationController "seredaes/go-passmem/App/Controllers/RegistrationController"
	"seredaes/go-passmem/App/Middleware"

	"github.com/gorilla/mux"
)

func StartWebServer(WebServer string) {

	router := mux.NewRouter()
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./public/assets/"))))

	api := router.PathPrefix("/api").Subrouter()
	api.Use(Middleware.AuthMiddleware)
	api.HandleFunc("/registration", RegistrationController.Registration).Methods("POST").Schemes("http").Name("registration")
	api.HandleFunc("/auth", AuthController.Auth).Methods("POST").Schemes("http").Name("auth")
	api.HandleFunc("/credentials", CredentialController.CredentialList).Methods("GET").Schemes("http").Name("credentials")
	api.HandleFunc("/credential", CredentialController.CreateCredential).Methods("POST").Schemes("http").Name("credentials")
	api.HandleFunc("/credential", CredentialController.UpdateCredential).Methods("PATCH").Schemes("http").Name("credentials")
	api.HandleFunc("/credential", CredentialController.DeleteCredential).Methods("DELETE").Schemes("http").Name("credentials")

	srv := &http.Server{
		Handler:      router,
		Addr:         WebServer,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	msg := fmt.Sprintf("PassMem Service v.0.1 \n\rMade in Geniward (seredaes@yandex.com)\nServer started at: http://%s\n", WebServer)
	fmt.Println(msg)

	log.Fatal(srv.ListenAndServe())

}
