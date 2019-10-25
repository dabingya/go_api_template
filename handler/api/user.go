package api

import (
	"github.com/gin-gonic/gin"
	"learn_go/go_api_template/handler/defs"
	"learn_go/go_api_template/handler/resp"
	"learn_go/go_api_template/libs/validate"
)

func HelloWorld(c *gin.Context) {
	var u defs.User

	_ = c.ShouldBindJSON(&u)

	if err, eMap := validate.ValidateApi(u); err != nil {
		resp.InvalidParams(c, eMap)

	} else {
		resp.Success(c, "ok")
		//c.JSON(resp.Success("ok"))
	}
}
