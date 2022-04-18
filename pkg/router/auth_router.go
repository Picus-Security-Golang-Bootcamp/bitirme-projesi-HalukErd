package router

import (
	//httpErr "BasketProjectGolang/internal/httpErrors"
	"BasketProjectGolang/internal/model/api"
	"BasketProjectGolang/internal/service"
	"BasketProjectGolang/pkg/config"
	jwtHelper "BasketProjectGolang/pkg/jwt"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"

	//jwtHelper "BasketProjectGolang/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type authRouter struct {
	service service.AuthService
	cfg     config.Config
}

func NewAuthRouter(service service.AuthService, cfg config.Config) *authRouter {
	return &authRouter{
		service: service,
		cfg:     cfg,
	}
}

func (router *authRouter) Register(r *gin.RouterGroup) {
	r.POST("/signup", router.Signup)
	//r.GET("/:id", h.getBasket)
	//r.DELETE("/:id", h.deleteBasket)
	//
	//r.POST("/item", h.addItem)
	//r.DELETE("/:id/item/:itemId", h.deleteItem)
	//r.PUT(":id/item/:item/quantity/:quantity", h.updateItem)
}

// Signup godoc
// @tags auth
// @Accept  json
// @Produce  json
// @Param request body api.SignupRequest true "SignupRequest"
// @Success 201
// @Router /signup [post]
func (router *authRouter) Signup(g *gin.Context) {
	var request api.SignupRequest
	if err := g.Bind(&request); err != nil {
		g.JSON(http.StatusBadRequest, err.Error())
	}
	response, err := router.service.Signup(g.Request.Context(), &request)
	if err != nil {
		g.JSON(http.StatusBadRequest, err.Error())
	} else {
		sessionTime := router.cfg.JWTConfig.SessionTime
		jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userId": response.Id,
			"email":  response.Email,
			"iat":    time.Now().Unix(),
			"iss":    os.Getenv("ENV"),
			"exp":    time.Now().Add(time.Duration(sessionTime) * time.Minute).Unix(),
			"roles":  response.Roles,
		})
		g.Header("Authorization", jwtHelper.GenerateToken(jwtClaims, &router.cfg))
		g.JSON(http.StatusCreated, map[string]string{"id": *response.Id})
	}
}

//
//func (a *authRouter) login(c *gin.Context) {
//	var req api.LoginRequest
//	if err := c.Bind(&req); err != nil {
//		c.JSON(httpErr.ErrorResponse(httpErr.NewRestError(http.StatusBadRequest, "check your request body", nil)))
//	}
//	a.service.
//	user := a.service.Login(*req.Email, *req.Password)
//	if user == nil {
//		c.JSON(httpErr.ErrorResponse(httpErr.NewRestError(http.StatusBadRequest, "user not found", nil)))
//
//	}
//	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
//		"userId": user.Id,
//		"email":  user.Email,
//		"iat":    time.Now().Unix(),
//		"iss":    os.Getenv("ENV"),
//		"exp":    time.Now().Add(24 * time.Hour).Unix(),
//		"roles":  user.Roles,
//	})
//	token := jwtHelper.GenerateToken(jwtClaims, a.cfg.JWTConfig.SecretKey)
//	c.JSON(http.StatusOK, token)
//}
