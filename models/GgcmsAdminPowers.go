package models

import (
	"fmt"
	"reflect"

	"github.com/astaxie/beego/orm"
)

type GgcmsAdminPowers struct {
	Id       int    `orm:"column(id);auto"`
	Memberid int    `orm:"column(memberid);null"`
	Typeid   int    `orm:"column(typeid);null"`
	Powers   string `orm:"column(powers);null"`
}

func (t *GgcmsAdminPowers) TableName() string {
	return "ggcms_admin_powers"
}

func init() {
	orm.RegisterModel(new(GgcmsAdminPowers))
}

// AddGgcmsAdminPowers insert a new GgcmsAdminPowers into database and returns
// last inserted Id on success.
func AddGgcmsAdminPowers(m *GgcmsAdminPowers) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetGgcmsAdminPowersById retrieves GgcmsAdminPowers by Id. Returns error if
// Id doesn't exist
func GetGgcmsAdminPowersById(id int) (v *GgcmsAdminPowers, err error) {
	o := orm.NewOrm()
	v = &GgcmsAdminPowers{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}
func GetCountGgcmsAdminPowers(query map[string]string) (count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsAdminPowers))
	var strs []string
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return 0, err
	}
	return qs.Count()
}
func ExistGgcmsAdminPowers(query map[string]string) (exist bool) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsAdminPowers))
	var strs []string
	var err error
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return false
	}
	return qs.Exist()
}

// GetAllGgcmsAdminPowers retrieves all GgcmsAdminPowers matches certain condition. Returns empty list if
// no records exist
func GetAllGgcmsAdminPowers(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64, c bool) (ml []interface{}, count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsAdminPowers))
	if qs, err = qsInit(&qs, query, sortby, order); err != nil {
		return nil, 0, err
	}
	count = 0
	var l []GgcmsAdminPowers

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

// UpdateGgcmsAdminPowers updates GgcmsAdminPowers by Id and returns error if
// the record to be updated doesn't exist
func UpdateGgcmsAdminPowersById(m *GgcmsAdminPowers) (err error) {
	o := orm.NewOrm()
	v := GgcmsAdminPowers{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteGgcmsAdminPowers deletes GgcmsAdminPowers by Id and returns error if
// the record to be deleted doesn't exist
func DeleteGgcmsAdminPowers(id int) (err error) {
	o := orm.NewOrm()
	v := GgcmsAdminPowers{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&GgcmsAdminPowers{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
func MultDeleteGgcmsAdminPowers(query map[string]string, ids []int) (num int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsAdminPowers))
	var strs []string
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return 0, err
	}
	if len(ids) > 0 {
		qs = qs.Filter("id__in", ids)
	}
	return qs.Delete()
}
