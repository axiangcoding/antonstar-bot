package app

import (
	"axiangcoding/antonstar/api-system/logging"
	"encoding/json"
	"fmt"
	"strings"
	"text/template"
)

type Pagination struct {
	// 分页页码，从1开始
	PageNum int `json:"page_num" binding:"required,gte=1" form:"page_num"`
	// 分页大小，最大为1000
	PageSize int `json:"page_size" binding:"required,gte=1,lte=1000" form:"page_size"`
	// 过滤，json格式
	Filter string `json:"filter" binding:"omitempty,json" form:"filter"`
	// 排序，json格式
	Sort string `json:"sort" binding:"omitempty,json" form:"sort"`
}

func (p Pagination) ToOffsetLimit() (int, int) {
	offset := (p.PageNum - 1) * p.PageSize
	limit := p.PageSize
	return offset, limit
}

// GetSortSql 获取sort的sql语句
func (p Pagination) GetSortSql() string {
	m := make(map[string]int)
	err := json.Unmarshal([]byte(p.Sort), &m)
	if err != nil {
		logging.Warn("get pagination sort failed")
		return ""
	}
	var item []string
	for k, v := range m {
		if v != 1 && v != -1 {
			logging.Warnf("unsupported sort value %d", v)
		}
		if v == 1 {
			item = append(item, k+" ASC")
		} else if v == -1 {
			item = append(item, k+" DESC")
		}
	}
	return strings.Join(item, ",")
}

// GetFilterSql 获取filter的sql语句。直接使用有sql注入的风险！
func (p Pagination) GetFilterSql() string {
	m := make(map[string]string)
	err := json.Unmarshal([]byte(p.Filter), &m)
	if err != nil {
		logging.Warn("get pagination filter failed")
		return ""
	}
	var item []string
	for k, v := range m {
		var opera string
		var isStr bool
		split := strings.SplitN(v, ":", 3)
		switch split[0] {
		case "eq":
			opera = "="
			break
		case "gt":
			opera = ">"
			break
		case "like":
			opera = "like"
		default:
			logging.Warnf("currently unsupported filter value %s", v)
		}
		switch split[1] {
		case "bool":
			isStr = false
			break
		case "num":
			isStr = false
			break
		default:
			isStr = true
		}
		val := template.HTMLEscapeString(split[2])
		if isStr {
			val = "\"" + val + "\""
		}
		pattern := "%s %s %s"
		item = append(item, fmt.Sprintf(pattern, k, opera, val))
	}
	return strings.Join(item, " AND ")
}
