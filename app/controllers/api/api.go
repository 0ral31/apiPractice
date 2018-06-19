package api

import (
	"github.com/revel/revel"
	"net/http"
)

type Api struct {
	*revel.Controller
}

func (c *Api) SetMessage(s string) *Api {
	c.Args["Message"] = s
	return c
}

func (c *Api) GetMessage() string {
	message, ok := c.Args["Message"].(string)
	if !ok {
		message = ""
	}
	return message
}

// エラーの際に返す Json 用の構造体
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Bad Request Error を返すやつ
func (c *Api) HandleBadRequestError() revel.Result {
	c.Response.Status = http.StatusBadRequest
	return c.RenderJSON(nil)
}

// 作成しようとしたリソースが既にある場合
func (c *Api) HandleConflictError() revel.Result {
	c.Response.Status = http.StatusConflict
	if c.GetMessage() == "" {
		c.SetMessage("作成しようとしたリソースがすでに存在しています")
	}
	return c.RenderJSON(nil)
}

// 指定できないIDを指定したとき
func (c *Api) UnSpecifiedId() *Api {
	c.SetMessage("idを指定することはできません。")
	return c
}

// Not Found Error を返すやつ
func (c *Api) HandleNotFoundError() revel.Result {
	c.Response.Status = http.StatusNotFound
	return c.RenderJSON(nil)
}

type ApiResultJSON struct {
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
}

func (c *Api) RenderJSON(o interface{}) revel.Result {
	if c.Response.Status == 0 {
		c.Response.Status = http.StatusOK
	}

	api_result_json := ApiResultJSON{
		Data:    o,
		Code:    c.Response.Status,
		Message: c.GetMessage(),
	}
	return c.Controller.RenderJSON(api_result_json)
}