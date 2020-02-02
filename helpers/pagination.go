package helpers

import (
	"github.com/gin-gonic/gin"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
)

// Param 分页参数
type Param struct {
	DB      *gorm.DB
	Page    int
	Limit   int
	OrderBy []string
	Filters []string
	Preload []string
	ShowSQL bool
}

// Paginator 分页返回
type Paginator struct {
	TotalRecord int         `json:"total_record"`
	TotalPage   int         `json:"total_page"`
	Records     interface{} `json:"records"`
	Offset      int         `json:"offset"`
	Limit       int         `json:"limit"`
	Page        int         `json:"page"`
	PrevPage    int         `json:"prev_page"`
	NextPage    int         `json:"next_page"`
}

// Paging 分页
func Paging(p *Param, result interface{}) *Paginator {
	db := p.DB

	if p.ShowSQL {
		db = db.Debug()
	}

	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit == 0 {
		p.Limit = 10
	}

	if len(p.OrderBy) > 0 {
		for _, o := range p.OrderBy {
			db = db.Order(o)
		}
	}

	if len(p.Filters) > 0 {
		for _, o := range p.Filters {
			db = db.Where(o)
		}
	}

	if len(p.Preload) > 0 {
		db = PreloadD(db, p.Preload)
	}

	done := make(chan bool, 1)
	var paginator Paginator
	var count int
	var offset int

	go countRecords(db, result, done, &count)

	if p.Page == 1 {
		offset = 0
	} else {
		offset = (p.Page - 1) * p.Limit
	}

	db.Limit(p.Limit).Offset(offset).Find(result)
	<-done

	paginator.TotalRecord = count
	paginator.Records = result
	paginator.Page = p.Page

	paginator.Offset = offset
	paginator.Limit = p.Limit
	paginator.TotalPage = int(math.Ceil(float64(count) / float64(p.Limit)))

	if p.Page > 1 {
		paginator.PrevPage = p.Page - 1
	} else {
		paginator.PrevPage = p.Page
	}

	if p.Page == paginator.TotalPage {
		paginator.NextPage = p.Page
	} else {
		paginator.NextPage = p.Page + 1
	}
	return &paginator
}

func countRecords(db *gorm.DB, anyType interface{}, done chan bool, count *int) {
	db.Model(anyType).Count(count)
	done <- true
}

func Page(g *gin.Context) int {
	page, _ := strconv.Atoi(g.DefaultQuery("page", "1"))
	return page
}

func Limit(g *gin.Context) int {
	page, _ := strconv.Atoi(g.DefaultQuery("limit", os.Getenv("LIMIT")))
	return page
}

func Order(g *gin.Context, order ...string) []string {
	var o []string

	if g.Query("order") != "" {

		order := strings.SplitN(g.Query("order"), "|", -1)

		for _, orderBy := range order {
			o = append(o, orderBy)
		}
		return o

	} else {

		for _, orderBy := range order {
			o = append(o, orderBy)
		}
		return o

	}
}

//func Order(order ...string)[]string   {
//	var o []string
//	for _ , orderBy := range order{
//		o = append(o , orderBy)
//	}
//	return o
//}
