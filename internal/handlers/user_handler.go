package handlers

import (
	"Threddit/internal/helpers"
	"Threddit/internal/models"
	"Threddit/internal/repositories"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"html/template"
	"net/http"
	"os"
	"time"
)

type UserHandler struct {
	repository *repositories.Repository
}

func (h *UserHandler) LoginPage(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")

	if err == nil {
		_, err = helpers.VerifyToken(cookie.Value)

		if err == nil {
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}

	tmpl, err := template.ParseFiles(
		templateRoot+"sections.html",
		templateRoot+"login.html",
		templateRoot+"navbar.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "login", nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")

	user, err := h.repository.FindByUsername(username)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.FormValue("password")))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	expirationTime := time.Now().Add(2 * time.Hour)

	claims := &helpers.Claims{
		Username: user.Username,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
		Path:    "/",
	})

	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *UserHandler) RegisterPage(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")

	if err == nil {
		_, err = helpers.VerifyToken(cookie.Value)

		if err == nil {
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}

	tmpl, err := template.ParseFiles(
		templateRoot+"sections.html",
		templateRoot+"register.html",
		templateRoot+"navbar.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "register", nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirm_password")

	if password != confirmPassword {
		http.Error(w, "Passwords do not match.", http.StatusInternalServerError)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := models.User{
		Email:    r.FormValue("email"),
		Username: r.FormValue("username"),
		Password: string(hashedPassword),
		Role:     models.Redditor,
	}

	if err := h.repository.CreateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	expirationTime := time.Now().Add(2 * time.Hour)

	claims := &helpers.Claims{
		Username: user.Username,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
		Path:    "/",
	})

	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "token",
		Value:  "",
		MaxAge: -1,
		Path:   "/",
	})

	http.Redirect(w, r, "/", http.StatusFound)
}
