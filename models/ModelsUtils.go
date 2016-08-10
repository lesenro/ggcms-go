package models

import (
	"errors"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
)

func qsInit(qs *orm.QuerySeter, query map[string]string, sortby []string, order []string) (qq orm.QuerySeter, err error) {
	// query k=v
	qq = *qs
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		//in查询
		if strings.Contains(k, "__in") {
			vs := make([]interface{}, 0)
			ss := strings.Split(v, "|")
			for _, s := range ss {
				if i, err := strconv.Atoi(s); err == nil {
					vs = append(vs, i)
				} else {
					vs = append(vs, s)
				}
			}
			qq = qq.Filter(k, vs)
		} else {
			qq = qq.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qq = qq.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}
	qq = qq.OrderBy(sortFields...)
	return qq, nil
}

func RunSql(query string, args []interface{}, fields string) (maps []orm.Params, err error) {
	strFillds := strings.Split(fields, ",")
	o := orm.NewOrm()
	_, err = o.Raw(query, args...).Values(&maps, strFillds...)
	return maps, err
}
