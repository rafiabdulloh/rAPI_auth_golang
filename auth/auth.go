package auth

import (
	"auth_with_token/models"
	"crypto/sha1"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

const (
	Secret = "secret"
)

var (
	Sign *jwt.Token
	PW   string
)

func DoHash(pw string) string {
	var sha = sha1.New()
	sha.Write([]byte(pw))

	result := sha.Sum(nil)
	return fmt.Sprintf("%x", result)
}

func LoginHandler(c *gin.Context, db *gorm.DB) {

	// lakukan penghubungan antara struct dan json
	var user models.Users
	var dataUser models.Users
	err := c.Bind(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"pesan": err.Error()})
	}

	if err := db.Where("username = ?", user.Username).First(&dataUser).Error; err != nil { //find just first row
		c.JSON(http.StatusBadRequest, gin.H{"data not found": user})
		return
	}

	user.Password = DoHash(user.Password)

	if user.Username != dataUser.Username {
		c.JSON(http.StatusBadRequest, gin.H{"pesan": "wrong user"})
		return
	} else {
		if user.Password != dataUser.Password {
			c.JSON(http.StatusUnauthorized, gin.H{
				"pesan": "wrong password ",
			})
			return
		}
	}

	// Mengatur waktu kadaluarsa token
	expirationTime := time.Now().Add(1 * time.Minute)

	// Membuat token dengan waktu kadaluarsa
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Menandatangani token dengan kunci rahasia (secret key)
	secret := []byte(dataUser.Password)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	PW = dataUser.Password

	c.JSON(200, gin.H{
		"token": signedToken,
	})
}
