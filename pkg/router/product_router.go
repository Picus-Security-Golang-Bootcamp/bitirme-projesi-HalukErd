package router

import (
	"BasketProjectGolang/internal/model/api"
	"BasketProjectGolang/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type productRouter struct {
	productService service.ProductService
	authService    service.AuthService
}

func NewProductRouter(productService service.ProductService, authService service.AuthService) *productRouter {
	return &productRouter{
		productService: productService,
		authService:    authService,
	}
}

func (router *productRouter) Register(r *gin.RouterGroup) {
	r.POST("/products", router.Create)
	//r.GET("/:id", h.getBasket)
	//r.DELETE("/:id", h.deleteBasket)
	//
	//r.POST("/item", h.addItem)
	//r.DELETE("/:id/item/:itemId", h.deleteItem)
	//r.PUT(":id/item/:item/quantity/:quantity", h.updateItem)
}

// Create godoc
// @tags products
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param request body api.CreateProductRequest true "CreateProductRequest"
// @Success 200
// @Router /api/v1/basket-api/products [post]
func (router *productRouter) Create(g *gin.Context) {
	err := router.authService.PermissionCheck(g, "admin")
	if err != nil {
		g.JSON(http.StatusForbidden, err.Error())
	}
	var request api.CreateProductRequest
	if err := g.Bind(&request); err != nil {
		g.JSON(http.StatusBadRequest, err.Error())
	}
	response, err := router.productService.Create(g.Request.Context(), &request)
	if err != nil {
		g.JSON(http.StatusBadRequest, err.Error())
	} else {
		g.JSON(http.StatusCreated, map[string]string{"id": response.ID})
	}
}
