package auth

import (
	"encoding/json"
	"net/http"

	"github.com/IsraelTeo/api-jwt-go/db"
	"github.com/IsraelTeo/api-jwt-go/model"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Cretendials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var credentials Cretendials

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := model.User{}
	if err := db.GDB.Where("email = ?", credentials.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Invalid access data", http.StatusUnauthorized)
		} else {
			http.Error(w, "Error when querying user", http.StatusInternalServerError)
		}
		return
	}

	if err := VerifyPassword(user.Password, credentials.Password); err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	jwtToken, err := GenerateToken(user.Email)
	if err != nil {
		http.Error(w, "Invalid access data", http.StatusInternalServerError)
	}

	response := model.NewResponse(model.MessageTypeSuccess, "Authentication success", jwtToken)
	model.ResponseJSON(w, http.StatusOK, response)

}

func VerifyPassword(passwordHashed, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHashed), []byte(password))
}
