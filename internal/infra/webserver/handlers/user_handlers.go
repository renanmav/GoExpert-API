package handlers

import (
	"encoding/json"
	"github.com/go-chi/jwtauth"
	"github.com/renanmav/GoExpert-API/configs"
	"github.com/renanmav/GoExpert-API/internal/dto"
	"github.com/renanmav/GoExpert-API/internal/entity"
	"github.com/renanmav/GoExpert-API/internal/infra/database"
	"net/http"
	"time"
)

type UserHandler struct {
	UserDB       database.UserInterface
	Jwt          *jwtauth.JWTAuth
	JwtExpiresIn int
}

func NewUserHandler(db database.UserInterface, config *configs.Config) *UserHandler {
	return &UserHandler{
		UserDB:       db,
		Jwt:          config.TokenAuth,
		JwtExpiresIn: config.JWTExpiresIn,
	}
}

// GetJWT godoc
// @Summary Get a JWT
// @Description Get a JWT
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.GetJWTInput true "User to authenticate"
// @Success 200 {object} dto.GetJWTOutput
// @Failure 400 {string} string "Bad request"
// @Failure 401 {string} string "Unauthorized"
// @Router /users/generate_token [post]
func (uh *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	var user dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := uh.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	payload := map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Minute * time.Duration(uh.JwtExpiresIn)).Unix(),
	}
	_, tokenString, _ := uh.Jwt.Encode(payload)
	response := dto.GetJWTOutput{AccessToken: tokenString}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// CreateUser godoc
// @Summary Create a user
// @Description Create a user
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.CreateUserInput true "User to create"
// @Success 201 {string} string "User created"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /users [post]
func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = uh.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
