package utils

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/gin-gonic/gin"
)

func Resp(message string, err error, data any) gin.H {
	return gin.H{
		"message": message,
		"error":   err,
		"data":    data,
	}
}

func StoreBodyToCtx(c *gin.Context) error {

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}

	c.Set("body", body)
	return nil
}

func LoadBodyFromCtx(c *gin.Context, dest any) error {

	dataAny, exists := c.Get("body")
	if !exists {
		return errors.New("键body不存在")
	}

	data, ok := dataAny.([]byte)
	if !ok {
		return errors.New("键body值的类型不是[]byte")
	}

	return json.Unmarshal(data, dest)
}
