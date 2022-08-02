package request

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	// MaxLimit defines query max value to field limit
	MaxLimit uint64 = 100
)

// Params used for requests when
// some controls params are allowed
type Params struct {
	// Fields returns as request queries
	Fields []string
	
	// Filters returns all filters of a request in map format
	Filters map[string][]string
	
	// Limit limit is used to set the size of a paginated list
	Limit uint64
	
	// Offset offset is used to determine which page 
	// should be in the paginated list
	Offset uint64
	
	// Total Total is used to fetch the total 
	// amount of a paginated list
	Total bool
}

// NewParams builds an empty dummy 
// `Params` object for utility usage
func NewParams() Params {
	return Params{
		Limit: MaxLimit,
		Filters: make(map[string][]string),
	}
}

// ParseParams receives the gin.Context and parse the query params for the request
func ParseParams(ctx *gin.Context) (params Params, err error) {
	var limit, offset int
	
	if limit, err = strconv.Atoi(ctx.DefaultQuery("limit", "15")); err != nil {
		return params, err
	}

	if limit <= 0 || limit > int(MaxLimit) {
		limit = int(MaxLimit)
	}
	params.Limit = uint64(limit)

	if offset, err = strconv.Atoi(ctx.DefaultQuery("offset", "0")); err != nil {
		return params, err
	}
	params.Offset = uint64(offset)

	params.Fields, _ = ctx.GetQueryArray("field")
	if params.Total, err = strconv.ParseBool(ctx.DefaultQuery("total", "false")); err != nil {
		return params, err
	}

	params.Filters = map[string][]string{}
	for k, v := range ctx.Request.URL.Query() {
		if k == "limit" || k == "offset" || k == "field" {
			continue
		}

		if len(v) > 0 {
			params.Filters[k] = append(params.Filters[k], v...)
		}
	}

	return
}

// AddFilter add a filter to request params
func (p *Params) AddFilter(name string, values ...string) *Params {
	if len(values) > 0 {
		if p.Filters == nil {
			p.Filters = map[string][]string{}
		}
		p.Filters[name] = values
	}
	return p
}

// ClearFilters clear the filters from the request params
func (p *Params) ClearFilters() *Params {
	p.Filters = make(map[string][]string)
	return p
}

// RemoveFilters remove filter values
func (p *Params) RemoveFilters(filters ...string) *Params {
	for i := range filters {
		delete(p.Filters, filters[i])
	}
	return p
}

// HasFilter returns true if the filter is searched for in the filter list
func (p *Params) HasFilter(filter string) (hasFilter bool) {
	if _, ok := p.Filters[filter]; ok {
		return true
	}
	return
}
