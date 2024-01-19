package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lu1s-souza/anGolar-scaffolding/controllers"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	projectRoutes := a.Group("/api/project")

	// Routes for GET method:
	projectRoutes.Get("", controllers.GetProjects)                                          // get list of all books
	projectRoutes.Get(":id", controllers.GetProject)                                        // get one book by ID
	projectRoutes.Get(":id/modules", controllers.GetProjectModules)                         // get one book by ID
	projectRoutes.Get(":id/module/:moduleName", controllers.GetModuleData)                  // get one book by ID
	projectRoutes.Get(":id/module/:moduleName/components", controllers.GetModuleComponents) // get one book by ID
	projectRoutes.Get(":id/module/:moduleName/services", controllers.GetModuleServices)     // get one book by ID
	projectRoutes.Get(":id/module/:moduleName/tests", controllers.GetModuleTests)           // get one book by ID
	projectRoutes.Get(":id/module/:moduleName/store", controllers.GetModuleStore)           // get one book by ID

}
