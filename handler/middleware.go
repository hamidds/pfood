package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/hamidds/pfood/model"
	"github.com/hamidds/pfood/utils"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		tokenString := request.Header.Get("Authorization")
		fmt.Println(tokenString)
		if len(tokenString) == 0 {
			writer.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(writer).Encode(model.NewError(errors.New("missing Authorization Header")))
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := utils.VerifyJWT(tokenString)
		if err != nil {
			writer.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(writer).Encode(model.NewError(errors.New("Error verifying JWT token: " + err.Error())))
			return
		}

		phoneNumber := claims.(jwt.MapClaims)["phone_number"].(string)
		fmt.Println(phoneNumber)
		request.Header.Set("phone_number", phoneNumber)
		next.ServeHTTP(writer, request)
		fmt.Println("End Of middleware")
	}
}
