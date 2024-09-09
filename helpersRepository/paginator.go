package helpersRepository

import (
	"math"

	"gorm.io/gorm"
)

type Paginator struct {
	Limit      int         `json:"limit" query:"limit"`
	Page       int         `json:"page" query:"page"`
	Sort       string      `json:"sort" query:"sort"`
	TotalRows  int64       `json:"total_rows" query:"total_rows"`
	TotalPages int         `json:"total_pages" query:"total_pages"`
	Rows       interface{} `json:"rows"`
}

/*
BuildPaginator marshal the model to a response serializer struct

	If `serializer` is nil, then will return all data from model
*/
func BuildPaginator(serializer interface{}) *Paginator {
	if serializer != nil {
		return &Paginator{
			Rows: serializer,
		}
	}
	return &Paginator{}
}

func (p *Paginator) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Paginator) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	if p.Limit > 100 {
		p.Limit = 100
	}
	return p.Limit
}

func (p *Paginator) GetPage() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Paginator) GetSort() string {
	if p.Sort == "" {
		p.Sort = "created_at desc"
	}
	return p.Sort
}

func (p *Paginator) Paginate(db *gorm.DB, value interface{}) *gorm.DB {
	var totalRows int64

	db = db.Model(value).Count(&totalRows)

	p.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(p.GetLimit())))
	p.TotalPages = totalPages

	return db.Offset(p.GetOffset()).Limit(p.GetLimit()).Order(p.GetSort())

}
