package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//
//type Resp struct {
//	ErrorCode int    `json:"error_code"`
//	Msg       string `json:"msg"`
//}
//
//type ErrResp struct {
//	Error     string `json:"error"`
//	ErrorCode int    `json:"error_code"`
//	Msg       string `json:"msg"`
//}

//var (
//	InvalidParams = ErrResp{ErrorCode: 1001,Msg:"invalid params",Error:""}
//	ServerError   = ErrResp{ErrorCode: 1002, Error: "server error"}
//)

//func InvalidParams(ss interface{}, msgs ...map[string]interface{}) (int, interface{}) {
//	res := make(map[string]interface{})
//	for _, p := range msgs {
//		for k, v := range p {
//			res[k] = v
//		}
//	}
//	res["msg"] = ss
//	res["errcode"] = 1003
//	//return http.StatusBadRequest, gin.H{"msg": ss, "errcode": 1003}
//	return http.StatusBadRequest, res
//}
//
//func NotFound(ss interface{}) (int, interface{}) {
//
//	return http.StatusNotFound, gin.H{"msg": ss, "errcode": 1004}
//}
//
//func ServerError(ss interface{}) (int, interface{}) {
//	return http.StatusInternalServerError, gin.H{"msg": "server error", "errcode": "1005"}
//}
//
//func Success(ss interface{}, ) (int, interface{}) {
//
//	return http.StatusOK, gin.H{"msg": ss, "errcode": "0"}
//}
//
//func SendSuccessResponse(r_code int, c *gin.Context, ) {
//
//}

func InvalidParams(c *gin.Context, msg interface{}) {

	c.JSON(http.StatusBadRequest, gin.H{"errcode": 1003, "msg": msg, "request": c.Request.URL.Path})
}

func Success(c *gin.Context, msg interface{}) {
	c.JSON(http.StatusOK, gin.H{"errcode": 0, "resp": msg, "request": c.Request.URL.Path})
}

func ServerError(c *gin.Context, msg interface{}) {
	c.JSON(http.StatusInternalServerError, gin.H{"errcode": 1010, "request": c.Request.URL.Path, "msg": msg})
}


