package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func NewAPIController() *APIController {
	return &APIController{}
}

type LogContext struct {
	RequestID   string    `json:"request_id"`
	RequestTime time.Time `json:"request_time"`
	IP          string    `json:"ip"`
}

type APIController struct {
	*gin.Context
	LogContext LogContext
}

func (s *APIController) CJSON(code int, data interface{}) {
	s.Logger()
	if code != 200 {
		resp := data
		fmt.Println("error", data)
		s.JSON(code, resp)
		return
	} else {
		resp := make(map[string]interface{})
		resp["data"] = data
		s.JSON(code, resp)
		return
	}
}
func (s *APIController) Handler(method func()) gin.HandlerFunc {
	return func(c *gin.Context) {
		s.LogContext.RequestID = c.GetHeader("request_id")
		if s.LogContext.RequestID == "" {
			s.LogContext.RequestID = "test"
		}
		s.LogContext.RequestTime = time.Now()
		s.LogContext.IP = c.ClientIP()
		s.Context = c
		method()
	}
}

func (s *APIController) Logger() {
	log.Printf("ip:%s", s.LogContext.IP)
	log.Printf("request_id:%s", s.LogContext.RequestID)
	log.Printf("request_time:%v", s.LogContext.RequestTime)
}
