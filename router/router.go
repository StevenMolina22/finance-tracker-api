package router

import (
	"database/sql"

	"github.com/StevenMolina22/fiber-turso/controller"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(db *sql.DB, app *fiber.App) {
	ctlr := controller.NewController(db)

	app.Get("/", ctlr.GetRoot)

	// Users
	users := app.Group("/users")
	users.Get("/", ctlr.ListUsers)
	users.Post("/", ctlr.CreateUser)
	users.Get("/:id", ctlr.GetUser)
	users.Put("/:id", ctlr.UpdateUser)
	users.Delete("/:id", ctlr.DeleteUser)

	// Transactions TODO: Add type to transactions (income, expense)
	transactions := app.Group("/transactions")
	transactions.Get("/", ctlr.GetTransactions)
	transactions.Post("/", ctlr.CreateTransaction)
	transactions.Get("/:id", ctlr.GetTransaction)
	transactions.Put("/:id", ctlr.UpdateTransaction)
	transactions.Delete("/:id", ctlr.DeleteTransaction)

	// Categories
	categories := app.Group("/categories")
	categories.Get("/", ctlr.ListCategories)
	categories.Post("/", ctlr.CreateCategory)
	categories.Get("/:id", ctlr.GetCategory)
	categories.Put("/:id", ctlr.UpdateCategory)
	categories.Delete("/:id", ctlr.DeleteCategory)

	// Assets
	assets := app.Group("/assets")
	assets.Get("/", ctlr.ListAssets)
	assets.Post("/", ctlr.CreateAsset)
	assets.Get("/:id", ctlr.GetAsset)
	assets.Put("/:id", ctlr.UpdateAsset)
	assets.Delete("/:id", ctlr.DeleteAsset)

	// Liabilities
	liabilities := app.Group("/liabilities")
	liabilities.Get("/", ctlr.ListLiabilities)
	liabilities.Post("/", ctlr.CreateLiability)
	liabilities.Get("/:id", ctlr.GetLiability)
	liabilities.Put("/:id", ctlr.UpdateLiability)
	liabilities.Delete("/:id", ctlr.DeleteLiability)

	// Budgets
	budgets := app.Group("/budgets")
	budgets.Get("/", ctlr.ListBudgets)
	budgets.Post("/", ctlr.CreateBudget)
	budgets.Get("/:id", ctlr.GetBudget)
	budgets.Put("/:id", ctlr.UpdateBudget)
	budgets.Delete("/:id", ctlr.DeleteBudget)
}
