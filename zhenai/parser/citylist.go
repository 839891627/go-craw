package parser

import (
	"crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	results := engine.ParseResult{}
	for _, m := range matches {
		results.Items = append(results.Items, "City "+string(m[2])) // 城市名字作为 item
		results.Requests = append(results.Requests, engine.Request{
			Url:        string(m[1]), // 城市 url
			ParserFunc: ParseCity,
		})
	}
	return results
}
