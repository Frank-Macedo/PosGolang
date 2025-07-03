package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Frank-Macedo/PosGoLang/cursoGo/Aulas/9-APIS/internal/dto"
	"github.com/Frank-Macedo/PosGoLang/cursoGo/Aulas/9-APIS/internal/entity"
	"github.com/go-chi/jwtauth"

	"github.com/Frank-Macedo/PosGoLang/cursoGo/Aulas/9-APIS/internal/infra/database"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserDb database.UserInterface
}

func NewUserHandler(db database.UserInterface) *UserHandler {
	return &UserHandler{
		UserDb: db,
	}
}

// GetJWT godoc
// @Summary      Get JWT token
// @Description  Generate JWT token for user authentication
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      dto.GetJWTInput  true  "User credentials"
// @Success      200  {object}  dto.GetJWTOutput
// @Failure      400  {object}  Error
// @Failure      401  {object}  Error
// @Failure      500  {object}  Error
// @Router       /users/generate_token [post]
func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {

	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("jwtExpiresIn").(int)

	var user dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := h.UserDb.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	_, token, err := jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	accessToken := dto.GetJWTOutput{
		AccessToken: token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)

}

// Create user godoc
// @Summary      Create a new user
// @Description  Create a new user with name, email and password
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      dto.CreateUserInput  true  "User data"
// @Failure      500  {object}  Error
// @Router	   /users [post]
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)

	fmt.Printf("Nome:%s  Email:%s PassWord:%s", user.Name, user.Email, user.Password)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := entity.NewUser(user.Name, user.Password, user.Email)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.UserDb.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
