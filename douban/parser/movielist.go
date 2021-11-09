package parser

import (
	"crawler/engine"
	"regexp"
)

var movieRe = regexp.MustCompile(`<a href="(https://movie.douban.com/subject/\d+\/)"`)

func ParseMovieList(contents []byte) engine.ParseResult {
	matches := movieRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseMovie,
		})
	}
	return result
}
