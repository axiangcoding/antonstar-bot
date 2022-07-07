package app

type Pagination struct {
	// 分页页码，从1开始
	PageNum int `json:"page_num" binding:"required,gte=1" form:"page_num"`
	// 分页大小，最大为1000
	PageSize int `json:"page_size" binding:"required,gte=1,lte=1000" form:"page_size"`
}

func (p Pagination) ToOffsetLimit() (int, int) {
	offset := (p.PageNum - 1) * p.PageSize
	limit := p.PageSize
	return offset, limit
}
