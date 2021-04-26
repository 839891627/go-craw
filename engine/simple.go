package engine

import (
	"crawler/fetcher"
	"log"
)

type SimpleEngine struct{}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}

// fetch + parse 抽成 worker
func worker(r Request) (ParseResult, error) {
	body, err := fetcher.Fetch(r.Url)
	log.Printf("Fetching %s", r.Url)

	if err != nil {
		log.Printf("Fetcher： error fetcheing url %s: %v", r.Url, err)
		return ParseResult{}, err
	}

	return r.ParserFunc(body), nil
}
