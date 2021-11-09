package parser

import (
	"crawler/douban/model"
	"crawler/engine"
	"regexp"
	strconv "strconv"
	"strings"
)

var titleRe = regexp.MustCompile(`<span property="v:itemreviewed">([^<]+)</span>`)
var yearRe = regexp.MustCompile(`<span class="year">\((\d+)\)</span>`)
var imgRe = regexp.MustCompile(`<img src="(https://img2.doubanio.com/view/photo/s_ratio_poster/public/[^"]+)"`)
var descRe = regexp.MustCompile(`<span property="v:summary" class="">([^<]+)</span>`)
var typeRe = regexp.MustCompile(`<span property="v:genre">([^<]+)</span>`)
var scoreRe = regexp.MustCompile(`<strong class="ll rating_num" property="v:average">([^<]+)</strong>`)
var scoreNumRe = regexp.MustCompile(`<span property="v:votes">(\d+)</span>`)

func ParseMovie(contents []byte) engine.ParseResult {
	movie := model.Movie{}

	movie.Title = extractString(contents, titleRe, false)
	movie.Year = extractString(contents, yearRe, false)
	movie.Img = extractString(contents, imgRe, false)
	movie.Desc = extractString(contents, descRe, false)
	movie.Type = extractString(contents, typeRe, true)
	if score, err := strconv.ParseFloat(extractString(contents, scoreRe, false), 64); err == nil {
		movie.Score = score
	}
	if scoreNum, err := strconv.Atoi(extractString(contents, scoreNumRe, false)); err == nil {
		movie.ScoreNum = scoreNum
	}

	result := engine.ParseResult{
		Items: []interface{}{movie},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp, multi bool) string {
	if !multi {
		match := re.FindSubmatch(contents)
		if len(match) >= 2 {
			return string(match[1])
		} else {
			return ""
		}
	} else {
		matches := re.FindAllSubmatch(contents, -1)
		s := ""
		for _, m := range matches {
			s += " / " + string(m[1])
		}
		return strings.TrimPrefix(s, " / ")
	}
}
