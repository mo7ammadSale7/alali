package Routes

import "server/Application"

type RouterApp struct {
	*Application.Application
}

func InitRoutes(app *Application.Application) {
	routerApp := RouterApp{app}
	routerApp.Routing()
}

func (app RouterApp) Routing() {
	app.VisitorsRoutes()
}
