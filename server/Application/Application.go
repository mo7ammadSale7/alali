package Application

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Application struct {
	Gin        *gin.Engine
	DB         *gorm.DB
	Connection *sql.DB
}

func (app *Application) share() {}

func app() func() *Application {
	return func() *Application {
		var application Application
		application.Gin = gin.Default()
		connectToDB(&application)
		return &application
	}
}

func NewApp() *Application {
	application := app()
	return application()
}
