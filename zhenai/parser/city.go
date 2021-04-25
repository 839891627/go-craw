package parser

import (
	"crawler/engine"
	"regexp"
)

const cityRe = `href="(http://localhost:8080/mock/album.zhenai.com/u/[0-9]+[^>]*)">([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	results := engine.ParseResult{}

	limit := 5 // 数组太多，只拿前几个试试
	for _, m := range matches {
		name := string(m[2])                                // 必须要赋值，不然在 ParserFunc 运行到的时候，指向最后一个 m[2]
		results.Items = append(results.Items, "User "+name) // 用户名
		results.Requests = append(results.Requests, engine.Request{
			Url: string(m[1]), // 个人页
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name)
			},
		})
		limit--
		if limit == 0 {
			break
		}
	}
	return results
}
