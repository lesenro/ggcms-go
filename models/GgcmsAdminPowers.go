package models

import (
	"reflect"

	"github.com/astaxie/beego/orm"
)

type GgcmsAdminPowers struct {
	Id       int    `orm:"column(id);auto"`
	Usertype string `orm:"column(usertype);size(20)"`
	Powerid  int    `orm:"column(powerid);null"`
}

func (t *GgcmsAdminPowers) TableName() string {
	return "ggcms_admin_powers"
}

type AdminPowers struct {
	GgcmsPowers
	Usertype string `orm:"column(usertype);size(20)"`
}

func (t *AdminPowers) TableName() string {
	return "v_admin_power"
}
func init() {
	orm.RegisterModel(new(GgcmsAdminPowers))
	orm.RegisterModel(new(AdminPowers))
}

func (this *AdminPowers) GetAll(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64, c bool) (ml []interface{}, count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(AdminPowers))
	if qs, err = qsInit(&qs, query, sortby, order); err != nil {
		return nil, 0, err
	}
	count = 0
	var l []AdminPowers

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
func UpdateGgcmsAdminPowersById(ids *[]int, usertype string) (err error) {
	o := orm.NewOrm()
	o.Begin()
	var pws []GgcmsAdminPowers
	for _, val := range *ids {
		pws = append(pws, GgcmsAdminPowers{Usertype: usertype, Powerid: val})
	}
	_, err = o.Raw("Delete From `ggcms_admin_powers` WHERE `usertype` = ?").SetArgs(usertype).Exec()
	if err != nil {
		o.Rollback()
		return
	}
	_, err = o.InsertMulti(100, pws)
	if err != nil {
		o.Rollback()
		return
	}
	o.Commit()
	return
}
