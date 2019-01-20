package pagination

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cast"
	"math"
)

const MaxPerPage = 100
const DefaultPerPage = 10

type Params struct {
	Page int
	PerPage int
}

type Pagination struct {
	Items interface{} `json:"items"`
	FirstItem int64 `json:"first_item"`
	LastItem int64 `json:"last_item"`
	Total int64 `json:"total"`
	Page int `json:"page"`
	TotalPage int `json:"total_page"`
	PerPage int `json:"per_page"`
}

func (p *Params) Normalize() {
	if p.Page < 1 {
		p.Page = 1
	}

	if p.PerPage > MaxPerPage {
		p.PerPage = MaxPerPage
	}
}

func getQueryTotal(q *gorm.DB, total chan int64) {
	var c int64
	q.Offset(-1).Count(&c)
	total <- c
}

func getQueryResults(q *gorm.DB, limit int, offset int64, out interface{}) {
	q.Limit(limit).Offset(offset).Find(out)
}

func (p *Params) Paginate(q *gorm.DB, rows interface{}) *Pagination {
	p.Normalize()

	pagination := &Pagination{}

	offset := int64((p.Page - 1) * p.PerPage)
	cTotal := make(chan int64)
	go getQueryTotal(q, cTotal)
	getQueryResults(q, p.PerPage, offset, rows)
	total := <- cTotal

	lastItem := offset + int64(p.PerPage)
	// In case last page have number of items smaller than per_page
	maxLastItem := total - 1
	if lastItem > maxLastItem {
		lastItem = maxLastItem
	}

	pagination.Total = total
	pagination.Items = rows
	pagination.FirstItem = offset
	pagination.LastItem = lastItem
	pagination.PerPage = p.PerPage
	pagination.Page = p.Page
	pagination.TotalPage = int(math.Ceil(float64(total / int64(p.PerPage))))

	return pagination
}

func GetParamsContext(c *gin.Context) *Params {
	q := c.Request.URL.Query()
	perPage := cast.ToInt(q.Get("per_page"))
	page := cast.ToInt(q.Get("page"))

	if perPage == 0 {
		perPage = DefaultPerPage
	}

	if page == 0 {
		page = 1
	}

	return &Params{
		Page: page,
		PerPage: perPage,
	}
}
