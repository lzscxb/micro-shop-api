package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitCommonRoute(Route *gin.RouterGroup) {
	commonRoute := Route.Group("base")
	{
		fmt.Println(commonRoute)
	}
}
