package users

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/kirantiloh/gameroom/pkg/auth"
)

type UserHandler struct {
	service UserService
}

func NewHandler(service UserService) UserHandler {
	return UserHandler{
		service: service,
	}
}

func (h *UserHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil {
		defer r.Body.Close()
	}

	userDto := &UserDto{}

	if err := json.NewDecoder(r.Body).Decode(userDto); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Bad Request",
		})
		return
	}

	if err := h.service.RegisterUser(userDto); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	verificationLink := fmt.Sprintf("%s/auth/verify/%s", os.Getenv("BASE_URL"), base64.URLEncoding.EncodeToString([]byte(userDto.Email)))

	mailBody := fmt.Sprintf(`
Hi %s,

Thanks for trying out Kiranti's Gameroom!

We need a little more information to complete your registration, including a confirmation of your email address.

Click below to confirm your email address:

<a href="%s">%s</a>

If you have problems, please paste the above URL into your web browser.

    <img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSHYmPwSqJ1DVMdGma5h-W5APgI3ic67qYTAQ&s" />
    `, html.EscapeString(userDto.Name), verificationLink, verificationLink)

	queue := make(chan struct{}, 5)

	go func() {
		queue <- struct{}{}
		defer func() {
			<-queue
		}()
		if err := auth.SendMail(userDto.Email, mailBody); err != nil {
			log.Printf("Error sending verification email: %v", err)
		}
	}()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Register successfull! Please check your email to verify your account",
	})
}

func (h *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil {
		defer r.Body.Close()
	}

	loginDto := &LoginDto{}

	if err := json.NewDecoder(r.Body).Decode(loginDto); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Bad Request",
		})
		return
	}

	user, err := h.service.VerifyUserCredentials(loginDto)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	if !user.IsVerified {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Account isn't verified yet, please check your inbox to verify your account",
		})
		return
	}

	token, err := auth.EncodeJWT(*user.toUserData())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Internal Server Error",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message":    "Login successfull",
		"auth_token": token,
	})

}

func (h *UserHandler) VerifyAccountHandler(w http.ResponseWriter, r *http.Request) {
	hash := chi.URLParam(r, "hash")

	email, _ := base64.URLEncoding.DecodeString(hash)

	user, err := h.service.GetUser(string(email))

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Not found",
		})
		return
	}

	user.IsVerified = true

	if err := h.service.UpdateUser(user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Internal Server Error",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Account verification is successfull! Enjoy",
	})
}
