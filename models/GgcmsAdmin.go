package models

import (
	"fmt"
	"reflect"

	"github.com/astaxie/beego/orm"
)

type GgcmsAdmin struct {
	Id        int    `orm:"column(id);auto"`
	Usertype  int    `orm:"column(usertype);null"`
	Userid    string `orm:"column(userid);size(30)"`
	Pwd       string `orm:"column(pwd);size(32)"`
	Uname     string `orm:"column(uname);size(20)"`
	Tname     string `orm:"column(tname);size(30)"`
	Email     string `orm:"column(email);size(30)"`
	Typeid    string `orm:"column(typeid);null"`
	Logintime int64  `orm:"column(logintime)"`
	Loginip   string `orm:"column(loginip);size(20)"`
}

func (t *GgcmsAdmin) TableName() string {
	return "ggcms_admin"
}

func init() {
	orm.RegisterModel(new(GgcmsAdmin))
}

// AddGgcmsAdmin insert a new GgcmsAdmin into database and returns
// last inserted Id on success.
func AddGgcmsAdmin(m *GgcmsAdmin) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetGgcmsAdminById retrieves GgcmsAdmin by Id. Returns error if
// Id doesn't exist
func GetGgcmsAdminById(id int) (v *GgcmsAdmin, err error) {
	o := orm.NewOrm()
	v = &GgcmsAdmin{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}
func GetCountGgcmsAdmin(query map[string]string) (count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsAdmin))
	var strs []string
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return 0, err
	}
	return qs.Count()
}
func ExistGgcmsAdmin(query map[string]string) (exist bool) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsAdmin))
	var strs []string
	var err error
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return false
	}
	return qs.Exist()
}

// GetAllGgcmsAdmin retrieves all GgcmsAdmin matches certain condition. Returns empty list if
// no records exist
func GetAllGgcmsAdmin(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64, c bool) (ml []interface{}, count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsAdmin))
	if qs, err = qsInit(&qs, query, sortby, order); err != nil {
		return nil, 0, err
	}
	count = 0
	var l []GgcmsAdmin

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

// UpdateGgcmsAdmin updates GgcmsAdmin by Id and returns error if
// the record to be updated doesn't exist
func UpdateGgcmsAdminById(m *GgcmsAdmin) (err error) {
	o := orm.NewOrm()
	v := GgcmsAdmin{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteGgcmsAdmin deletes GgcmsAdmin by Id and returns error if
// the record to be deleted doesn't exist
func DeleteGgcmsAdmin(id int) (err error) {
	o := orm.NewOrm()
	v := GgcmsAdmin{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&GgcmsAdmin{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
func MultDeleteGgcmsAdmin(query map[string]string, ids []int) (num int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsAdmin))
	var strs []string
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return 0, err
	}
	if len(ids) > 0 {
		qs = qs.Filter("id__in", ids)
	}
	return qs.Delete()
}
