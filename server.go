package main
import (
	"net/http"
	"github.com/labstack/echo"
	_ "gopkg.in/cq.v1"
	"realworld/controllers"
)
func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "HELLO FROM API")
	})
	e.POST("/createUser", controllers.CreateUser)
	e.POST("/checkUserLogin", controllers.CheckUserLogin)
	e.POST("/addUserInterests", controllers.CreateAddInterests)
	e.POST("/getUserInterests", controllers.FetchInterests)
	e.POST("/sendConnectionReq",controllers.SendConnectionRequest)
	e.POST("/acceptConnectionReq",controllers.AcceptConnectionRequest)
	e.POST("/blockUser",controllers.BlockUser)
	e.POST("/unblockUser",controllers.UnBlockUser)
	e.Logger.Fatal(e.Start(":8000"))
}