package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Frank-Macedo/PosGoLang/cursoGo/Aulas/9-APIS/configs"
	_ "github.com/Frank-Macedo/PosGoLang/cursoGo/Aulas/9-APIS/docs"
	"github.com/Frank-Macedo/PosGoLang/cursoGo/Aulas/9-APIS/internal/entity"
	"github.com/Frank-Macedo/PosGoLang/cursoGo/Aulas/9-APIS/internal/infra/database"
	"github.com/Frank-Macedo/PosGoLang/cursoGo/Aulas/9-APIS/internal/infra/webServer/handlers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title           Go expert api
// @version         1.0
// @description     Product api with authentication
// @termsOfService  http://swagger.io/terms/
// @contact.name    Franklin Maceddo
// @contact.url     http://www.swagger.io/support
// @contact.email
// @license.name    Apache 2.0
// @license.url     http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8080
// @BasePath  /
// @schemes  http
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	UserDB := database.NewUser(db)
	UserHandler := handlers.NewUserHandler(UserDB)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("jwtExpiresIn", configs.JWTExpiresIn))

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
		r.Get("/", productHandler.GetProducts)
	})

	r.Post("/users", UserHandler.Create)
	r.Post("/users/generate_token", UserHandler.GetJWT)

	r.Get("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/docs/doc.json"),
	))
	http.ListenAndServe(":8080", r)
}

func LogRequest(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		fmt.Println("cabou o midlleware")

	})
}
