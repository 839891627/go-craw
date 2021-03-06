package parser

import (
	"crawler/engine"
	"regexp"
)

var profileRe = regexp.MustCompile(`<a href="(http://localhost:8888/mock/album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)

// 更多城市（下一页）
var cityUrlRe = regexp.MustCompile(`<a href="(http://localhost:8888/mock/www.zhenai.com/zhenghun/[^"]+)"`)

func ParseCity(contents []byte) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		//result.Items = append(result.Items, "User "+name)

		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name)
			},
		})
	}
	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}
	return result
}
