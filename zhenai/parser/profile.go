package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"
)

var nameRe = regexp.MustCompile(`<h1 class="ceiling-name ib fl fs24 lh32 blue">([^<]+)</h1>`)
var descRe = regexp.MustCompile(`<p class="ceiling-age fs14 lh22 c5e">([^<]+)</p>`)
var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>(\d+)岁</td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>(\d+)CM</td>`)
var guessRe = regexp.MustCompile(`href="(http://localhost:8888/mock/album.zhenai.com/u/[^"+])">([^<]+)</a>`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}

	profile.Name = name
	profile.Desc = extractString(contents, descRe)

	if age, err := strconv.Atoi(extractString(contents, ageRe)); err == nil {
		profile.Age = age
	}
	if height, err := strconv.Atoi(extractString(contents, heightRe)); err == nil {
		profile.Height = height
	}

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	matches := guessRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		name := string(m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(bytes []byte) engine.ParseResult {
				return ParseProfile(bytes, name)
			},
		})
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
