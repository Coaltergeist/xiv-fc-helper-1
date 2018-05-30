package xivdb

import "fmt"

var searchTypeEndpoints = map[RequestType]string{
	CHARACTER: "one=characters&",
	ITEM:      "one=items&",
}

var queryTypeEndpoints = map[RequestType]string{
	CHARACTER: "character/",
	ITEM:      "item/",
}

const searchURL = baseURL + "search?"
const baseURL = "https://api.xivdb.com/"

func (r *SearchRequest) generateEndpoint() {
	if r.request.requestType < 0 {
		panic("invalid requestType (< 0)")
	}
	if _, ok := searchTypeEndpoints[r.request.requestType]; !ok {
		panic("invalid requestType (not found)")
	}
	if len(r.searchString) <= 0 {
		panic("invalid searchString")
	}
	r.request.endpoint = fmt.Sprintf("%s%sstring=%s", searchURL, searchTypeEndpoints[r.request.requestType], r.searchString)
}

func (r *QueryRequest) generateEndpoint() {
	if r.request.requestType < 0 {
		panic("invalid requestType (< 0)")
	}
	if _, ok := queryTypeEndpoints[r.request.requestType]; !ok {
		panic("invalid requestType (not found)")
	}
	if r.id < 0 {
		panic("invalid request id")
	}
	r.request.endpoint = fmt.Sprintf("%s%s%d", baseURL, queryTypeEndpoints[r.request.requestType], r.id)
}
