package models

import (
	"fmt"
	"reflect"

	"github.com/astaxie/beego/orm"
)

type GgcmsPowers struct {
	Id        int    `orm:"column(id);auto"`
	Powername string `orm:"column(powername);size(255);null"`
	Powertype int    `orm:"column(powertype);null"`
	Icon      string `orm:"column(icon);size(255);null"`
	Menuid    string `orm:"column(menuid);size(255);null"`
	Menuurl   string `orm:"column(menuurl);size(255);null"`
	Funccode  string `orm:"column(funccode);size(255);null"`
	Pid       int    `orm:"column(pid);null"`
	Orderid   int    `orm:"column(orderid);null"`
}

func (t *GgcmsPowers) TableName() string {
	return "ggcms_powers"
}

func init() {
	orm.RegisterModel(new(GgcmsPowers))
}

// AddGgcmsPowers insert a new GgcmsPowers into database and returns
// last inserted Id on success.
func AddGgcmsPowers(m *GgcmsPowers) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetGgcmsPowersById retrieves GgcmsPowers by Id. Returns error if
// Id doesn't exist
func GetGgcmsPowersById(id int) (v *GgcmsPowers, err error) {
	o := orm.NewOrm()
	v = &GgcmsPowers{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}
func GetCountGgcmsPowers(query map[string]string) (count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsPowers))
	var strs []string
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return 0, err
	}
	return qs.Count()
}
func ExistGgcmsPowers(query map[string]string) (exist bool) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsPowers))
	var strs []string
	var err error
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return false
	}
	return qs.Exist()
}

// GetAllGgcmsPowers retrieves all GgcmsPowers matches certain condition. Returns empty list if
// no records exist
func GetAllGgcmsPowers(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64, c bool) (ml []interface{}, count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsPowers))
	if qs, err = qsInit(&qs, query, sortby, order); err != nil {
		return nil, 0, err
	}
	count = 0
	var l []GgcmsPowers

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

// UpdateGgcmsPowers updates GgcmsPowers by Id and returns error if
// the record to be updated doesn't exist
func UpdateGgcmsPowersById(m *GgcmsPowers) (err error) {
	o := orm.NewOrm()
	v := GgcmsPowers{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteGgcmsPowers deletes GgcmsPowers by Id and returns error if
// the record to be deleted doesn't exist
func DeleteGgcmsPowers(id int) (err error) {
	o := orm.NewOrm()
	v := GgcmsPowers{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&GgcmsPowers{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
func MultDeleteGgcmsPowers(query map[string]string, ids []int) (num int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsPowers))
	var strs []string
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return 0, err
	}
	if len(ids) > 0 {
		qs = qs.Filter("id__in", ids)
	}
	return qs.Delete()
}
