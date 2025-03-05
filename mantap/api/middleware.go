package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Zulhaidir/microservice/mantap/token"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_type"
)

func authMiddleWare(tokenMaker token.Maker) gin.HandlerFunc {
	// gin.HandlerFunc adalah type yant menggunakan anonymous function, sehingga juga harus menggunakan anonymous function
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		// handling untuk panjang authorizationHeader jika 0 atau tidak ada
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		// untuk menyimpan payload ke context
		ctx.Set(authorizationPayloadKey, payload)
		// untuk meneruskan permintaan ke handler lain
		ctx.Next()
	}
}
