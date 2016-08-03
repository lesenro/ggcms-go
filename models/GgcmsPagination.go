package models

import (
	"strconv"
	"strings"
)

type GgcmsPagination struct {
	PageNum       int   //当前页
	RowTotal      int64 //总记录数
	PageSize      int   //每页显示条数
	NavigatePages int   //导航显示页数
	UrlTemplate   string

	PageTotal       int      //共多少页
	PrePage         int      //上一页
	NextPage        int      //下一页
	IsFirstPage     bool     //是第一页
	IsLastPage      bool     //是最后一页
	HasPreviousPage bool     //有上一页
	HasNextPage     bool     //有下一页
	ShowPagination  bool     //无记录时不显示
	PageList        []int    //分页列表
	PageUrls        []string //分页url
	FirstPageUrl    string   //第一页url
	LastPageUrl     string   //最后一页
	PrevPageUrl     string   //上页url
	NextPageUrl     string   //下页url
}

func (this *GgcmsPagination) CalcPages() {
	if this.RowTotal == 0 {
		this.ShowPagination = false
		return
	}
	this.ShowPagination = true
	//计算总页数
	this.PageTotal = int(this.RowTotal / int64(this.PageSize))
	if this.RowTotal%int64(this.PageSize) > 0 {
		this.PageTotal = this.PageTotal + 1
	}
	//设置导航显示页数，缺省为5
	if this.NavigatePages == 0 {
		this.NavigatePages = 5
	}
	//默认分页url模板
	if len(this.UrlTemplate) == 0 {
		this.UrlTemplate = "[n]"
	}
	//设置缺省值
	this.PrePage = this.PageNum - 1
	this.NextPage = this.PageNum + 1
	this.IsFirstPage = false
	this.HasPreviousPage = true
	this.IsLastPage = false
	this.HasNextPage = true
	this.FirstPageUrl = this.getUrl(1)
	this.LastPageUrl = this.getUrl(this.PageTotal)
	this.PrevPageUrl = this.getUrl(this.PrePage)
	this.NextPageUrl = this.getUrl(this.NextPage)
	//第一页
	if this.PageNum <= 1 {
		this.PageNum = 1
		this.PrePage = 1
		this.IsFirstPage = true
		this.HasPreviousPage = false
	}
	//最后一页
	if this.PageNum >= this.PageTotal {
		this.PageNum = this.PageTotal
		this.NextPage = this.PageTotal
		this.IsLastPage = true
		this.HasNextPage = false
	}
	//分页列表
	var startPage int
	startPage = this.PageNum - this.NavigatePages/2
	if this.PageTotal-startPage < this.NavigatePages {
		startPage = this.PageTotal - this.NavigatePages + 1
	}
	if startPage < 1 {
		startPage = 1
	}
	if this.PageTotal > this.NavigatePages && this.PageTotal-startPage < this.NavigatePages {
		startPage = this.PageTotal - this.NavigatePages + 1
	}
	for i := 0; i < this.NavigatePages; i++ {
		page := i + startPage
		this.PageList = append(this.PageList, page)
		this.PageUrls = append(this.PageUrls, this.getUrl(page))
		if page == this.PageTotal {
			break
		}
	}
}
func (this *GgcmsPagination) getUrl(pn int) string {
	return strings.Replace(this.UrlTemplate, "[n]", strconv.Itoa(pn), -1)
}
