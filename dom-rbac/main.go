package main

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"

	"casbin_apply/rbac/middleware"
)

func main() {
	// 要使用自己定义的数据库rbac_db,最后的true很重要.默认为false,使用缺省的数据库名casbin,不存在则创建
	a, err := middleware.GetAdapter()
	if err != nil {
		log.Printf("连接数据库错误: %v", err)
		return
	}

	e, err := casbin.NewEnforcer("D:/mark/go/src/casbin_apply/dom-rbac/model.conf", a)
	if err != nil {
		log.Printf("初始化casbin错误: %v", err)
		return
	}
	//从DB加载策略
	e.LoadPolicy()

	//获取router路由对e
	r := gin.New()

	r.POST("/api/v1/add", func(c *gin.Context) {
		fmt.Println("增加Policy")
		if ok, err := e.AddGroupingPolicy("mark", "admin", "domain1"); !ok {
			fmt.Println("GroupingPolicy已经存在")
			fmt.Println("group policy err : ", err)
			return
		}

		if ok, _ := e.AddPolicy("admin", "domain1", "/api/v1/hello", "GET"); !ok {
			fmt.Println("Policy已经存在")
		} else {
			fmt.Println("增加成功")
		}

	})
	//删除policy
	r.DELETE("/api/v1/delete", func(c *gin.Context) {
		fmt.Println("删除Policy")
		if ok, _ := e.RemovePolicy("admin", "domain1", "/api/v1/hello", "GET"); !ok {
			fmt.Println("Policy不存在")
		} else {
			fmt.Println("删除成功")
		}
	})
	//获取policy
	r.GET("/api/v1/get", func(c *gin.Context) {
		fmt.Println("查看policy")
		list := e.GetPolicy()
		for _, vlist := range list {
			for _, v := range vlist {
				fmt.Printf("value: %s, ", v)
			}
		}
	})
	//使用自定义拦截器中间件
	r.Use(middleware.CasbinMiddleware(e))
	//创建请求
	r.GET("/api/v1/hello", func(c *gin.Context) {
		fmt.Println("Hello 接收到GET请求..")
	})

	r.Run(":9000") //参数为空 默认监听8080端口
}
