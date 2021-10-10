package middleware

import (
    "fmt"
    "encoding/json"
    "github.com/unrolled/render"
    "github.com/go-xorm/xorm"
)

type BaseController struct {
    Render    *render.Render
    DB        *xorm.Engine
}

type ApiResponse struct {
    Code    int         `json:"code"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}

func (a BaseController) Log(v ...interface{}) {
    fmt.Println(v)
}

func (ctrl BaseController) NewResponse() *ApiResponse {
    return &ApiResponse{}
}

func (ctrl *ApiResponse) ToJSON() []byte {
    result, _ := json.Marshal(ctrl)
    return result
}
