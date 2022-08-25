package framework

import (
	"log"
	"net/http"
)

type Core struct {
	router map[string]ControllerHandler
}

func NewCore() *Core {
	return &Core{router: map[string]ControllerHandler{}}
}

func (c *Core) Get(url string, handler ControllerHandler) {
	c.router[url] = handler

}

//!框架核心实现Hanlder接口
func (c *Core) ServeHTTP(reponse http.ResponseWriter, request *http.Request) {
	// fmt.Fprintf(reponse, "Hello my frame")

	log.Println("core.ServerHTTP ", request.URL.Path)
	ctx := NewContext(request, reponse)

	//?简单路由
	router := c.router["foo"] //*根据url获取配皮对应路由
	if router == nil {        //*习惯问题
		log.Println("url can't match the router")
		return
	}
	log.Println("core.router")

	router(ctx)
}
