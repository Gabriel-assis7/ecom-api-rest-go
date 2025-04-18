package user

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gabriel-assis7/ecom-api-rest-go/service/auth"
	"github.com/gabriel-assis7/ecom-api-rest-go/types"
	"github.com/gabriel-assis7/ecom-api-rest-go/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var payload types.LoginUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		error := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("validation error: %v", error))
		return
	}

	u, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid credentials: %v", err))
		return
	}

	if err = auth.CheckPasswordHash(payload.Password, u.Password); err != nil {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid credentials: %v", err))
		return
	}

	token, err := auth.CreateJwt([]byte(os.Getenv("JWT_SECRET")), u.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("could not create token: %v", err))
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		error := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("validation error: %v", error))
		return
	}

	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user already exists with email: %s", payload.Email))
		return
	} else if err.Error() != "user not found" {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("could not check if user exists: %v", err))
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("could not hash password: %v", err))
		return
	}

	if err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	}); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("could not create user: %v", err))
		return
	}
}
