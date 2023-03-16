package route

import (
	"github.com/gorilla/mux"
	"goblog/logger"
	"net/http"
)

var route *mux.Router

// Name2URL 通过路由名称来获取 URL
func Name2URL(routeName string, pairs ...string) string {
	url, err := route.Get(routeName).URL(pairs...)
	if err != nil {
		logger.LogError(err)
		return ""
	}
	return url.String()
}

func GetRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}

func SetRoute(r *mux.Router) {
	route = r
}
