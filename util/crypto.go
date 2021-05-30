package util

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	logger "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"time"
)

var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

var tokenGenerator *jwt.Token


func init(){
	privateKeyFile, err := ioutil.ReadFile("private.key")
	if err!=nil{
		logger.Error("Error occurred while reading privateKey file: ",err)
	}
	
	publicKeyFile, err := ioutil.ReadFile("public.key")
	if err!=nil{
		logger.Error("Error occurred while reading publicKey file: ",err)
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateKeyFile)
	if err!=nil{
		logger.Error("Error occurred while parsing private key: ",err)
	}
	
	publicKey,err =jwt.ParseRSAPublicKeyFromPEM(publicKeyFile)

	if err!=nil{
		logger.Error("Error occurred while parsing public key: ",err)
	}

	tokenGenerator = jwt.New(jwt.SigningMethodRS256)
}

func EncryptPassword(plainPassword string)(string,error){

	password, err := bcrypt.GenerateFromPassword([]byte(plainPassword), 10)

	encryptedPassword :=string(password[:])

	return encryptedPassword,err
}

type Claim struct {
	UserId    int    `json:"userId"`
	TokenType string `json:"tokenType"`
	jwt.StandardClaims
}



func GenerateToken( userId int,tokenType string)(string,error){

	expiresAt := time.Now().Add(time.Minute * 30).Unix()

	claim:=&Claim{
		UserId:         userId,
		TokenType:      tokenType,
		StandardClaims: jwt.StandardClaims{ExpiresAt: expiresAt},
		}

	tokenGenerator.Claims=claim

	token, err := tokenGenerator.SignedString(privateKey)

	return token, err
}

func VerifyToken(stringToken string) (*Claim,error){

	claim:=&Claim{}

	_, err := jwt.ParseWithClaims(stringToken, claim, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err!=nil {
		logger.Error("Error while verifying token :",err)
	}

	return claim,err
}

