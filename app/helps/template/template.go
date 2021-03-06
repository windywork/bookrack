package template

import (
    "html/template"
    "time"
)

var AppHelpers = template.FuncMap{
    "unescaped": unescaped,
    "date": date,
    "time": dateTime,
    "ResourceVersion": func() string {return "10"},
    "Slogan": func() string {return ""},
}

func unescaped(x string) interface{} {
    return template.HTML(x)
}

func date(d time.Time) string {
    return d.Format("2006-01-02")
}

func dateTime(d time.Time) string {
    return d.Format("2006-01-02 15:04:05")
}
