package routers

// import (

// "
// 	middlewares "github.com/SwanHtetAungPhyo/esfor/internal/middleware"
// 	"github.com/SwanHtetAungPhyo/esfor/internal/utils"
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/golang-jwt/jwt/v4"
// 	"go.uber.org/zap"
// )

// type RouterService struct {
// 	Rlog *zap.Logger
// }

// // func NewRouterService(logger *zap.Logger)
// func NewRouterService() *RouterService {
// 	return &RouterService{
// 		Rlog: utils.GetLogger(),
// 	}
// }
// func (r *RouterService) UserRouter(app *fiber.App) {
// 	r.Rlog.Info("UserRoute setting up ......")
// 	userGroup := app.Group("/user").Use(middlewares.JwtMiddleware())

// 	userGroup.Get("/")
// }

// func (r *RouterService) PublicRoutes(app *fiber.App) {
// 	r.Rlog.Info("Protected route")
// 	publicGroup := app.Group("/public")
// 	publicGroup.Get("/tokens", handlers.BaseHandle(func(c *fiber.Ctx) (interface{}, error) {
// 		accesstoken, _ := r.geneToken(1)
// 		refreshToken, _ := r.geneToken(48)
// 		c.Cookie(&fiber.Cookie{
// 			Name:     "refresh",
// 			Value:    *refreshToken,
// 			Path:     "/",
// 			MaxAge:   36000,
// 			Secure:   false,
// 			HTTPOnly: true,
// 		})
// 		return accesstoken, nil
// 	}))

// }
// func (r *RouterService) geneToken(duration int) (*string, error) {
// 	token := jwt.New(jwt.SigningMethodHS256)
// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["name"] = " Swan"
// 	claims["admin"] = true
// 	claims["exp"] = time.Now().Add(time.Hour * time.Duration(duration)).Unix()

// 	tokenString, err := token.SignedString([]byte(middlewares.SecretKey))
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &tokenString, nil
// }
