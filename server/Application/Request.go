package Application

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"server/Models"
)

// ShareResource interface
type ShareResource interface {
	share()
}

type Request struct {
	Context    *gin.Context
	DB         *gorm.DB
	Connection *sql.DB
	User       *Models.User
	IsAuth     bool
}

func (req *Request) share() {}

// Handle Request closure data
func req() func(c *gin.Context) *Request {
	return func(c *gin.Context) *Request {
		var request Request
		request.Context = c
		connectToDB(&request)
		return &request
	}
}

// NewRequest init closure
func NewRequest(c *gin.Context) *Request {
	request := req()
	return request(c)
}

// Response req method to single response
func (req *Request) Response(code int, message string, data interface{}) {
	CloseConnection(req)
	req.Context.JSON(code, gin.H{
		"message": message,
		"data":    data,
	})
}

func (req *Request) Auth() *Request {
	req.IsAuth = false
	authHeader := req.Context.GetHeader("Authorization")
	if authHeader != "" {
		req.DB.Where("token = ?", authHeader).First(&req.User)
		if req.User.ID != 0 {
			req.IsAuth = true
		}
	}
	return req
}
