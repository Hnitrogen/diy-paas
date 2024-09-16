package server

import (
	"diy-paas/internal/app/routers"
	"diy-paas/internal/global"
	"fmt"
	"github.com/emicklei/go-restful/v3"
	//"github.com/emicklei/go-restful-swagger12"
	"log"
	"net/http"
)

func InitRestFulServer() {
	//restful.DefaultContainer.Router(restful.CurlyRouter{})
	global.DefaultContainer = restful.NewContainer()
	global.DefaultContainer.ServeMux = http.DefaultServeMux
	global.DefaultContainer.Router(restful.CurlyRouter{})
	// 用global.DefaultContainer 替换 restful.DefaultContainer --- 制造全局路由容器

	routers.RegisterBaseRoutes()
	// 打印所有的api
	allApis := global.DefaultContainer.RegisteredWebServices()
	fmt.Println("Registered APIs:")
	for _, api := range allApis {
		//fmt.Printf("API: %s\n", api.Path("/"))
		for _, route := range api.Routes() {
			fmt.Printf("  - Method: %s, Path: %s\n", route.Method, route.Path)
		}
	}
	// 启动swagger --- 最新的go-restful已经不支持swagger-ui
	//config := swagger.Config{
	//	WebServices:    restful.DefaultContainer.RegisteredWebServices(),
	//	WebServicesUrl: "http://localhost:8080",
	//	ApiPath:        "/apidocs.json",
	//
	//	// Optionally, specify where the UI is located
	//	SwaggerPath:     "/apidocs/",
	//	SwaggerFilePath: "Users/17914/GolandProjects/diy-paas/swagger-dist"}
	//swagger.RegisterSwaggerService(config, restful.DefaultContainer)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
