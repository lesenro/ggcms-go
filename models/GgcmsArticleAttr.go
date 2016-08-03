package models

import (
	"fmt"
	"reflect"
	"time"

	"github.com/astaxie/beego/orm"
)

type GgcmsArticleAttr struct {
	Id           int       `orm:"column(id);auto"`
	Articleid    int       `orm:"column(articleid);null"`
	AttrUrl      string    `orm:"column(attrUrl);size(255)"`
	Info         string    `orm:"column(info);size(100)"`
	Size         int       `orm:"column(size);null"`
	Originalname string    `orm:"column(originalname);size(100);null"`
	Addtime      time.Time `orm:"column(addtime);type(datetime);null"`
}

func (t *GgcmsArticleAttr) TableName() string {
	return "ggcms_article_attr"
}

func init() {
	orm.RegisterModel(new(GgcmsArticleAttr))
}

// AddGgcmsArticleAttr insert a new GgcmsArticleAttr into database and returns
// last inserted Id on success.
func AddGgcmsArticleAttr(m *GgcmsArticleAttr) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetGgcmsArticleAttrById retrieves GgcmsArticleAttr by Id. Returns error if
// Id doesn't exist
func GetGgcmsArticleAttrById(id int) (v *GgcmsArticleAttr, err error) {
	o := orm.NewOrm()
	v = &GgcmsArticleAttr{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}
func GetCountGgcmsArticleAttr(query map[string]string) (count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsArticleAttr))
	var strs []string
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return 0, err
	}
	return qs.Count()
}
func ExistGgcmsArticleAttr(query map[string]string) (exist bool) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsArticleAttr))
	var strs []string
	var err error
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return false
	}
	return qs.Exist()
}

// GetAllGgcmsArticleAttr retrieves all GgcmsArticleAttr matches certain condition. Returns empty list if
// no records exist
func GetAllGgcmsArticleAttr(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64, c bool) (ml []interface{}, count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsArticleAttr))
	if qs, err = qsInit(&qs, query, sortby, order); err != nil {
		return nil, 0, err
	}
	count = 0
	var l []GgcmsArticleAttr

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

// UpdateGgcmsArticleAttr updates GgcmsArticleAttr by Id and returns error if
// the record to be updated doesn't exist
func UpdateGgcmsArticleAttrById(m *GgcmsArticleAttr) (err error) {
	o := orm.NewOrm()
	v := GgcmsArticleAttr{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteGgcmsArticleAttr deletes GgcmsArticleAttr by Id and returns error if
// the record to be deleted doesn't exist
func DeleteGgcmsArticleAttr(id int) (err error) {
	o := orm.NewOrm()
	v := GgcmsArticleAttr{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&GgcmsArticleAttr{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
func MultDeleteGgcmsArticleAttr(query map[string]string, ids []int) (num int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsArticleAttr))
	var strs []string
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return 0, err
	}
	if len(ids) > 0 {
		qs = qs.Filter("id__in", ids)
	}
	return qs.Delete()
}
