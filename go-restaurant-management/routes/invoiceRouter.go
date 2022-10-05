package routes

import (
	controller "github.com/JakubPluta/go-projects/go-restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(r *gin.Engine) {
	r.GET("/invoices", controller.GetInvoices())
	r.GET("/invoices/:invoice_id", controller.GetInvoice())
	r.POST("/invoices", controller.CreateInvoice())
	r.PATCH("/invoices/:invoice_id", controller.UpdateInvoice())

}
