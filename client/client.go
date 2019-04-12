package client

import (
	utils "myFirstapp/dbutils"
	u "myFirstapp/utility"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type Client struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
}

func (client *Client) Validate() (map[string]interface{}, bool) {

	if !strings.Contains(client.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if len(client.Password) < 6 {
		return u.Message(false, "Password is required"), false
	}

	temp := &Client{}

	//check for errors and duplicate emails
	db := utils.GetDB()
	err := db.Table("clients").Where("email = ?", client.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if temp.Email != "" {
		return u.Message(false, "Email address already in use by another user."), false
	}

	return u.Message(false, "Requirement passed"), true
}

func (client *Client) Create() map[string]interface{} {

	if resp, ok := client.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(client.Password), bcrypt.DefaultCost)
	client.Password = string(hashedPassword)

	db := utils.GetDB()
	db.Create(client)

	if client.ID <= 0 {
		return u.Message(false, "Failed to create account, connection error.")
	}
	client.Password = "" //delete password

	response := u.Message(true, "Account has been created")
	response["client"] = client
	return response
}

func (client *Client) Login(email, password string) map[string]interface{} {

	err := utils.GetDB().Table("clients").Where("email = ?", email).First(client).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return u.Message(false, "Invalid login credentials. Please try again")
	}
	//Worked! Logged In
	client.Password = ""

	//Create JWT token
	tk := &Token{UserId: client.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	client.Token = tokenString //Store the token in the response

	resp := u.Message(true, "Logged In")
	resp["client"] = client
	return resp
}

// func CreateClientHandlers(router *mux.Router) {
// 	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")

// 	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
// }
