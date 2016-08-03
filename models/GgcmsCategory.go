package models

import (
	"fmt"
	"reflect"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type GgcmsCategory struct {
	Id           int       `orm:"column(id);auto"`
	Categoryname string    `orm:"column(categoryname);size(255);null"`
	Orderid      int       `orm:"column(orderid);null"`
	Articlenum   int       `orm:"column(articlenum);null"`
	Imgartnum    int       `orm:"column(imgartnum);null"`
	Logo         string    `orm:"column(logo);size(255);null"`
	Styledir     string    `orm:"column(styledir);size(100)"`
	Lastupdate   time.Time `orm:"column(lastupdate);type(datetime);null"`
	Pid          int       `orm:"column(pid);null"`
	Ppath        string    `orm:"column(ppath);size(255);null"`
	Depth        int       `orm:"column(depth);null"`
	Child        int       `orm:"column(child);null"`
	Ctempname    string    `orm:"column(ctempname);size(100)"`
	Atempname    string    `orm:"column(atempname);size(100)"`
	Curl         string    `orm:"column(curl);size(255);null"`
	Pagesize     int       `orm:"column(pagesize);null"`
	NavPages     int       `orm:"column(navpages);null"`
	Imgwidth     int       `orm:"column(imgwidth);null"`
	Imgheight    int       `orm:"column(imgheight);null"`
	Rssfeed      string    `orm:"column(rssfeed);size(255);null"`
	Keywords     string    `orm:"column(Keywords);size(255);null"`
	Description  string    `orm:"column(description);size(255);null"`
	Content      string    `orm:"column(content);null"`
	Extattrib    string    `orm:"column(extattrib);size(255);null"`
	Siteid       int       `orm:"column(siteid);null"`
	Ctype        int8      `orm:"column(ctype);null"`
	Mid          int       `orm:"column(mid);null"`
}

func (t *GgcmsCategory) TableName() string {
	return "ggcms_category"
}

func init() {
	orm.RegisterModel(new(GgcmsCategory))
}

// AddGgcmsCategory insert a new GgcmsCategory into database and returns
// last inserted Id on success.
func AddGgcmsCategory(m *GgcmsCategory) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetGgcmsCategoryById retrieves GgcmsCategory by Id. Returns error if
// Id doesn't exist
func GetGgcmsCategoryById(id int) (v *GgcmsCategory, err error) {
	o := orm.NewOrm()
	v = &GgcmsCategory{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}
func GetCountGgcmsCategory(query map[string]string) (count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsCategory))
	var strs []string
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return 0, err
	}
	return qs.Count()
}
func ExistGgcmsCategory(query map[string]string) (exist bool) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsCategory))
	var strs []string
	var err error
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return false
	}
	return qs.Exist()
}

// GetAllGgcmsCategory retrieves all GgcmsCategory matches certain condition. Returns empty list if
// no records exist
func GetAllGgcmsCategory(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64, c bool) (ml []interface{}, count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsCategory))
	if qs, err = qsInit(&qs, query, sortby, order); err != nil {
		return nil, 0, err
	}
	count = 0
	var l []GgcmsCategory

	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(&v).Elem()
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		if c {
			count, _ = qs.Count()
		}
		beego.Alert(ml)
		return ml, count, nil
	}
	return nil, 0, err
}
func UpdateSortData(items []GgcmsCategory) (err error) {
	o := orm.NewOrm()
	o.Begin()
	p, err := o.Raw("UPDATE `ggcms_category` SET orderid = ?, child = ?, pid = ?, depth = ?, ppath = ? WHERE `id` = ?").Prepare()
	for _, v := range items {
		if _, err = p.Exec(v.Orderid, v.Child, v.Pid, v.Depth, v.Ppath, v.Id); err != nil {
			o.Rollback()
			p.Close()
			return err
		}
	}
	p.Close()
	o.Commit()
	return
}

// UpdateGgcmsCategory updates GgcmsCategory by Id and returns error if
// the record to be updated doesn't exist
func UpdateGgcmsCategoryById(m *GgcmsCategory) (err error) {
	o := orm.NewOrm()
	v := GgcmsCategory{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteGgcmsCategory deletes GgcmsCategory by Id and returns error if
// the record to be deleted doesn't exist
func DeleteGgcmsCategory(id int) (err error) {
	o := orm.NewOrm()
	v := GgcmsCategory{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&GgcmsCategory{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
func MultDeleteGgcmsCategory(query map[string]string, ids []int) (num int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsCategory))
	var strs []string
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return 0, err
	}
	if len(ids) > 0 {
		qs = qs.Filter("id__in", ids)
	}
	return qs.Delete()
}
