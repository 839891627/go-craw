package main

import (
	"crawler/engine"
	"crawler/zhenai/parser"
)

func main() {
	url := "http://127.0.0.1:8080/mock/www.zhenai.com/zhenghun"

	engine.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseCityList,
	})
}
