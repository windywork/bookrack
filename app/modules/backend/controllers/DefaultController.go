package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/unrolled/render"
    "net/http"
    "github.com/windywork/bookrack/app/middleware"
)

type DefaultController struct {
    middleware.BaseController
}

func (ctrl DefaultController) Index(c *gin.Context) {
    ctrl.Render.HTML(c.Writer, http.StatusOK, "backend/default/index", gin.H{
        "Title": "backend",
    }, render.HTMLOptions{Layout: "backend/_layouts/main"})
}

func (ctrl DefaultController) Login(c *gin.Context) {
    ctrl.Render.HTML(c.Writer, http.StatusOK, "backend/default/login", gin.H{
        "Title": "backend",
    }, render.HTMLOptions{Layout: "backend/_layouts/empty"})
}
