package models

import (
	"fmt"
	"reflect"

	"github.com/astaxie/beego/orm"
)

type GgcmsArticlePages struct {
	Id        int    `orm:"column(id);auto"`
	Articleid int    `orm:"column(articleid);null"`
	Content   string `orm:"column(content);null"`
	Sortid    int    `orm:"column(sortid);null"`
	Title     string `orm:"column(title);size(255);null"`
	OldSortid int    `orm:"-"`
}
type ArticlePages struct {
	Pages []GgcmsArticlePages
}

func (t *GgcmsArticlePages) TableName() string {
	return "ggcms_article_pages"
}

func init() {
	orm.RegisterModel(new(GgcmsArticlePages))
}

// AddGgcmsArticlePages insert a new GgcmsArticlePages into database and returns
// last inserted Id on success.
func AddGgcmsArticlePages(m *GgcmsArticlePages) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetGgcmsArticlePagesById retrieves GgcmsArticlePages by Id. Returns error if
// Id doesn't exist
func GetGgcmsArticlePagesById(id int) (v *GgcmsArticlePages, err error) {
	o := orm.NewOrm()
	v = &GgcmsArticlePages{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}
func GetGgcmsArticlePagesBySortId(aid, sid int) (v *GgcmsArticlePages, err error) {
	o := orm.NewOrm()
	v = &GgcmsArticlePages{Articleid: aid, Sortid: sid}
	if err = o.Read(v, "Articleid", "Sortid"); err == nil {
		return v, nil
	}
	return nil, err
}
func GetCountGgcmsArticlePages(query map[string]string) (count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsArticlePages))
	var strs []string
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return 0, err
	}
	return qs.Count()
}
func ExistGgcmsArticlePages(query map[string]string) (exist bool) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsArticlePages))
	var strs []string
	var err error
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return false
	}
	return qs.Exist()
}

// GetAllGgcmsArticlePages retrieves all GgcmsArticlePages matches certain condition. Returns empty list if
// no records exist
func GetAllGgcmsArticlePages(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64, c bool) (ml []interface{}, count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsArticlePages))
	if qs, err = qsInit(&qs, query, sortby, order); err != nil {
		return nil, 0, err
	}
	count = 0
	var l []GgcmsArticlePages

	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		if c {
			count, _ = qs.Count()
		}
		return ml, count, nil
	}
	return nil, 0, err
}

// UpdateGgcmsArticlePages updates GgcmsArticlePages by Id and returns error if
// the record to be updated doesn't exist
func UpdateGgcmsArticlePagesById(m *GgcmsArticlePages) (err error) {
	o := orm.NewOrm()
	v := GgcmsArticlePages{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteGgcmsArticlePages deletes GgcmsArticlePages by Id and returns error if
// the record to be deleted doesn't exist
func DeleteGgcmsArticlePages(id int) (err error) {
	o := orm.NewOrm()
	v := GgcmsArticlePages{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&GgcmsArticlePages{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
func MultDeleteGgcmsArticlePages(query map[string]string, ids []int) (num int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsArticlePages))
	var strs []string
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return 0, err
	}
	if len(ids) > 0 {
		qs = qs.Filter("id__in", ids)
	}
	return qs.Delete()
}
