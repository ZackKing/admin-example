package middleware

import (
	"admin-api/internal/dao"
	"admin-api/internal/dao/model"
	"admin-api/internal/services/log"
	"admin-api/internal/services/uuid"
	"admin-api/pkg/lib/def"
	"admin-api/pkg/lib/encoding/json"

	"github.com/gin-gonic/gin"
)

func Access(c *gin.Context) {
	c.Set("sys_access_id", uuid.LogNode.Generate().String())

	c.Next()

	if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "DELETE" {
		accessLog(c)
	}
}

func accessLog(c *gin.Context) {
	q := dao.GetQuery()

	l := model.AccessLog{
		AccessID: c.GetString("sys_access_id"),
		UID:      int32(c.GetInt("uid")),
		Method:   c.Request.Method,
		Path:     c.Request.URL.Path,
		Header:   getHaederJson(c),
		Query:    getQueryJson(c),
		Body:     "{}",
		IP:       c.ClientIP(),
		Response: "{}",
	}

	if err := q.AccessLog.Create(&l); err != nil {
		log.Error("accescc_log_create_error", err, def.KV{"data": l})
	}

}

func getHaederJson(c *gin.Context) string {
	s := json.MustEncode2Str(c.Request.Header)
	if s == "" {
		s = "{}"
	}
	return s
}

func getQueryJson(c *gin.Context) string {
	s := json.MustEncode2Str(c.Request.URL.Query())
	if s == "" {
		s = "{}"
	}
	return s
}

// TODO: should get before c.Next(), wait for thinking
// func getBodyJson(c *gin.Context) string {
// 	b, _ := c.GetRawData()

// 	if len(b) == 0 {
// 		return "{}"
// 	}
// 	s := string(b)
// 	if ok := json.Valid(s); ok {
// 		return s
// 	} else {
// 		return json.MustEncode2Str(map[string]string{
// 			"raw": s,
// 		})
// 	}
// }
