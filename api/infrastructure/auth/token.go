package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type Token struct{}

func NewToken() *Token {
	return &Token{}
}

type TokenInterface interface {
	CreateToken(userid uint64) (*TokenDetails, error)
	ExtractTokenMetadata(*http.Request) (*AccessDetails, error)
}

func (t *Token) CreateToken(userid uint64) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AccessTokenExpires = time.Now().Add(time.Hour * 5).Unix()
	td.TokenUuid = uuid.NewV4().String()

	td.RefreshTokenExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = td.TokenUuid + "++" + strconv.Itoa(int(userid))

	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.TokenUuid
	atClaims["user_id"] = strconv.Itoa(int(userid))
	atClaims["exp"] = td.AccessTokenExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}
	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = strconv.Itoa(int(userid))
	rtClaims["exp"] = td.RefreshTokenExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return td, nil
}

func TokenValid(r *http.Request, authInterface AuthInterface) (uint64, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return 0, err
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return 0, err
	}

	_, err = authInterface.FetchAuth(claim["access_uuid"].(string))
	if err != nil {
		return 0, err
	}

	tmp, err := strconv.Atoi(claim["user_id"].(string))
	if err != nil {
		return 0, nil
	}
	return uint64(tmp), nil
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

//get the token from the request body
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func (t *Token) ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	fmt.Println("WE ENTERED METADATA")
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId, err := strconv.ParseUint(claims["user_id"].(string), 10, 64)
		if err != nil {
			return nil, err
		}
		return &AccessDetails{
			TokenUuid: accessUuid,
			UserId:    userId,
		}, nil
	}
	return nil, err
}
