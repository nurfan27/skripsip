package qasircore

import "math"

type Pagination struct {
	Pagesize  int
	NowPage   int
	TotalPage int
}

func (this *Pagination) SetPageSize(size int) {
	this.Pagesize = size
}

func (this *Pagination) GetPageSize() int {
	return this.Pagesize
}

func (this *Pagination) SetTotalPage(total int) {
	totalAll := total
	d := float64(totalAll) / float64(this.Pagesize)
	totalPages := int(math.Ceil(d))

	if this.Pagesize == 0 {
		this.Pagesize = 15
	}

	this.TotalPage = totalPages
}

func (this *Pagination) GetTotalPage() int {
	return this.TotalPage
}

func (this *Pagination) SetNowPage(page int) {
	this.NowPage = page
}

func (this *Pagination) GetNowPage() int {
	return this.NowPage
}

func (this *Pagination) GetDataPage() map[string]interface{} {
	return map[string]interface{}{
		"current_page": this.NowPage,
		"page_size":    this.Pagesize,
		"total_page":   this.TotalPage,
	}
}
