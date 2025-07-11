package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// type RefreshTokenClaims struct {
// 	CustomerOnlineID int    `json:"customer_online_id"`
// 	CustomerID       int    `json:"customer_id"`
// 	ShopID           int    `json:"shop_id"`
// 	LineID           string `json:"line_id"`
// 	CitizenID        string `json:"citizen_id"`
// 	jwt.RegisteredClaims
// }

type AccessTokenClaims struct {
	CustomerOnlineID int    `json:"customer_online_id"`
	CustomerID       int    `json:"customer_id"`
	ShopID           int    `json:"shop_id"`
	ShopMotherID     int    `json:"shop_mother_id"`
	LineID           string `json:"line_id"`
	CitizenID        string `json:"citizen_id"`
	jwt.RegisteredClaims
}

// var rftkKey = []byte(os.Getenv("JWT_RF_KEY"))
var actkKey = []byte(os.Getenv("JWT_AC_KEY"))

func CheckPublicKey(c *gin.Context) {
	bearer := c.Request.Header.Get("Authorization")
	pbtkToken := strings.TrimPrefix(bearer, "Bearer ")
	if pbtkToken != os.Getenv("TK_PUBPLIC_KEY") {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Invalid token.",
		})
		c.Abort()
		return
	} else {
		c.Next()
		return
	}
}

func CheckTelePublicKey(c *gin.Context) {
	bearer := c.Request.Header.Get("Authorization")
	pbtkToken := strings.TrimPrefix(bearer, "Bearer ")
	if pbtkToken != os.Getenv("TK_TELE_KEY") {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Invalid token.",
		})
		c.Abort()
		return
	} else {
		c.Next()
		return
	}
}

func CreateAccessToken(customerOnlineID int, customerID int, shopID int, shopMID int, lineID string, citizenID string) string {
	addMinuteExpire, errENV := strconv.Atoi(os.Getenv("JWT_AC_EXPIRE"))
	if errENV != nil {
		return ""
	} else {
		timeCreated := time.Now()
		// timeCreatedInt := timeCreated.Unix()
		timeExpire := timeCreated.Add(time.Duration(addMinuteExpire) * time.Minute)
		// timeExpireInt := timeExpire.Unix()
		claims := &AccessTokenClaims{
			CustomerOnlineID: customerOnlineID,
			CustomerID:       customerID,
			ShopID:           shopID,
			ShopMotherID:     shopMID,
			LineID:           lineID,
			CitizenID:        citizenID,
			RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(timeExpire)},
		}
		actkClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		actkToken, errSS := actkClaims.SignedString(actkKey)
		if errSS != nil {
			return ""
		}
		return actkToken
	}
}

func CheckAccessToken(c *gin.Context) {
	// check access token is valid
	bearer := c.Request.Header.Get("Authorization")
	actkToken := strings.TrimPrefix(bearer, "Bearer ")
	actkClaims := &AccessTokenClaims{}
	_, errPWC := jwt.ParseWithClaims(actkToken, actkClaims, func(token *jwt.Token) (interface{}, error) {
		return actkKey, nil
	})
	if errPWC != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "Invalid or expire token.",
		})
		c.Abort()
		return
	} else {
		c.AddParam("customerOnlineId", strconv.Itoa(actkClaims.CustomerOnlineID))
		c.AddParam("customerId", strconv.Itoa(actkClaims.CustomerID))
		c.AddParam("shopId", strconv.Itoa(actkClaims.ShopID))
		c.AddParam("shopMotherId", strconv.Itoa(actkClaims.ShopMotherID))
		c.AddParam("lineId", actkClaims.LineID)
		c.AddParam("citizenId", actkClaims.CitizenID)
		c.Next()
		return
	}
}

func DecodeToken(tokenString string) (AccessTokenClaims, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return AccessTokenClaims{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return AccessTokenClaims{}, fmt.Errorf("invalid claims type")
	}

	fmt.Println("Decoded claims:", claims)

	jsonData, _ := json.Marshal(claims)
	var tokenData AccessTokenClaims
	_ = json.Unmarshal(jsonData, &tokenData)

	return tokenData, nil
}
