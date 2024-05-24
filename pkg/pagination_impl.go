package pkg

import (
	"math"
	"mvhamadiqbalriv/belajar-golang-restful-api/helper"
	"net/http"

	"gorm.io/gorm"
)

type PaginationImpl struct {
    Limit        int         `json:"limit,omitempty;query:limit"`   
    Page         int         `json:"page,omitempty;query:page"` 
    Sort         string      `json:"sort,omitempty;query:sort"` 
    TotalRows    int64       `json:"total_rows"`    
    TotalPages   int         `json:"total_pages"`   
    Rows         interface{} `json:"rows"`
}

func NewPagination() Pagination {
    return &PaginationImpl{}
}

func (pagination *PaginationImpl) GetOffset() int {
    return (pagination.Page - 1) * pagination.Limit
}

func (pagination *PaginationImpl) GetLimit() int {
    if pagination.Limit == 0 {
        return 10
    }
    return pagination.Limit
}

func (pagination *PaginationImpl) GetPage() int {
    if pagination.Page == 0 {
        return 1
    }
    return pagination.Page
}

func (pagination *PaginationImpl) GetSort() string {
    if pagination.Sort == "" {
        return "id desc"
    }
    return pagination.Sort
}

func Paginate(value interface{}, pagination *PaginationImpl, db *gorm.DB) func(db *gorm.DB) *gorm.DB {  

    var totalRows int64 
    db.Model(value).Count(&totalRows)
    pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))
    pagination.TotalPages = totalPages

    return func(db *gorm.DB) *gorm.DB { 
        return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())   
    }   
}   

func ExtractQueryParams(r *http.Request) (PaginationImpl) {
    query := r.URL.Query()

    limit := query.Get("limit")
    if limit == "" {
        limit = "10"
    }

    page := query.Get("page")
    if page == "" {
        page = "1"
    }

    sort := query.Get("sort")
    if sort == "" {
        sort = "id desc"
    }

    pagination := PaginationImpl{
        Limit: helper.StringToInt(limit),
        Page:  helper.StringToInt(page),
        Sort:  sort,
    }

    return pagination
}


