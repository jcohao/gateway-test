package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"

	"gateway-test/dao"
	"gateway-test/handlers"
	"gateway-test/setting"
)

func main() {
	fmt.Println("project start...")

	if len(os.Args) < 2 {
		fmt.Println("Usage：./gateway-test conf/config.ini")
		return
	}
	// 加载配置文件
	if err := setting.Init(os.Args[1]); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}
	// 连接数据库
	err := dao.InitMySQL(setting.Config.MySQLConfig)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	fmt.Println("database connect success")
	defer dao.Close()

	r := mux.NewRouter()
	r.HandleFunc("/user/add", handlers.CreateUserHandler)
	r.HandleFunc("/user/get", handlers.GetUserHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe()

}
