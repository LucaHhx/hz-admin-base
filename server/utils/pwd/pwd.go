package pwd

import (
	"context"
	"hab/global"
	systemReq "hab/model/system/request"
	"hab/utils"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func SetPwdToken(id uint) string {
	idStr := utils.Itoa(id)
	key := "pwd_login:" + idStr
	token := idStr + "-" + uuid.New().String()
	exp := 7 * 24 * time.Hour

	global.HAB_REDIS.Set(context.Background(), key, token, exp)
	global.BlackCache.Set(key, token, exp)
	return token
}

var tokenErr = errors.New("token is invalid or has expired")

func CheckPwdToken(c *gin.Context) (uint, error) {
	token := c.GetHeader("x-token")
	idStr := ""
	if i := strings.Index(token, "-"); i >= 0 {
		idStr = token[:i]
	}
	if idStr == "" {
		return 0, tokenErr
	}
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, tokenErr
	}
	id := uint(idInt)

	key := "pwd_login:" + idStr
	t, ok := global.BlackCache.Get(key)
	if ok {
		if t != token {
			return 0, tokenErr
		}
		SetClaims(c)
		return id, nil
	}

	t, err = global.HAB_REDIS.Get(context.Background(), key).Result()
	if err != nil {
		global.HAB_LOG.Error("GetPwdToken Error " + err.Error())
		return 0, tokenErr
	}
	if t != token {
		return 0, tokenErr
	}
	global.BlackCache.Set(key, token, 7*24*time.Hour)
	SetClaims(c)
	return id, nil
}

func SetClaims(c *gin.Context) {
	c.Set("claims", &systemReq.CustomClaims{
		BaseClaims:       systemReq.BaseClaims{AuthorityId: 1},
		BufferTime:       0,
		RegisteredClaims: jwt.RegisteredClaims{},
	})
}
