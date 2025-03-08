package main

import "github.com/Gradient-and-Co/sigma-school/internal/app"

// @title           Sigma School API
// @version         1.0
// @description     This is a Sigma School API.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:80
// @BasePath  /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	app.RunWeb()
}
