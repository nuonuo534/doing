package main

import (
	"fmt"
	"net/http"

	"server/algorithm"
	"server/routers"
	"server/setting"
)

func main() {
	arr := []int{6, 7, 2, 34, 5, 8, 9, 1}
	algorithm.BubbleSort(arr)
	fmt.Printf("arr : %v\n", arr)

	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
