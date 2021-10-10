package main

import (
    "fmt"
    "github.com/fvbock/endless"
    "github.com/gin-gonic/contrib/static"
    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"
    "github.com/unrolled/render"
    "html/template"
    "github.com/windywork/bookrack/app/middleware"
    "github.com/windywork/bookrack/app/helps/setting"
    et "github.com/windywork/bookrack/app/helps/template"
    "github.com/windywork/bookrack/app/routers"
)
var (
    router        *gin.Engine
    genericRender *render.Render
    database      *xorm.Engine
)
func main() {
    setting.Setup()
    gin.SetMode(setting.ServerSetting.RunMode)
    router = gin.Default()
    router.Use(static.Serve("/static", static.LocalFile("static", true)))
    appCtrl := middleware.BaseController{}

    var err error
    database, err = xorm.NewEngine(setting.DatabaseSetting.Driver, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
                setting.DatabaseSetting.User,
                setting.DatabaseSetting.Password,
                setting.DatabaseSetting.Host,
                setting.DatabaseSetting.Name))
    if err != nil {
        fmt.Printf("链接数据库错误: %v", err)
    }
    appCtrl.DB = database

    genericRender = render.New(render.Options{
        Directory:  "templates",
        Layout:     "_layout/main",
        Extensions: []string{".tmpl", ".html"},
        Funcs:     []template.FuncMap{et.AppHelpers},
        Delims:     render.Delims{"{{", "}}"},
        Charset:    "UTF-8",
        IndentJSON: false,
        IndentXML:  false,
        //PrefixJSON: []byte(")]}',\n"),
        PrefixXML: []byte("<?xml version='1.0' encoding='UTF-8'?>"),
        HTMLContentType: "text/html",
        IsDevelopment:             setting.ServerSetting.RunMode == gin.DebugMode,
        UnEscapeHTML:              true,
        StreamingJSON:             true,
        RequirePartials:           true,
        DisableHTTPErrorRendering: true,
    })
    appCtrl.Render = genericRender

    routers.InitRoutes(router.Group("/"), appCtrl)
    _ = endless.ListenAndServe(":" + setting.ServerSetting.HttpPort, router)
}
