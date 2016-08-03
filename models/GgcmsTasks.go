package models

import (
	"fmt"
	"reflect"

	"github.com/astaxie/beego/orm"
)

type GgcmsTasks struct {
	Id       int    `orm:"column(id);auto"`
	Taskname string `orm:"column(taskname);size(255);null"`
	Tasktype string `orm:"column(tasktype);size(100);null"`
	Taskset  string `orm:"column(taskset);null"`
	State    int    `orm:"column(state);null"`
	Onoff    int    `orm:"column(onoff);null"`
	Plantype string `orm:"column(plantype);size(100);null"`
	Nexttime int64  `orm:"column(nexttime);null"`
	Planset  string `orm:"column(planset);null"`
}

func (t *GgcmsTasks) TableName() string {
	return "ggcms_tasks"
}

func init() {
	orm.RegisterModel(new(GgcmsTasks))
}

// AddGgcmsTasks insert a new GgcmsTasks into database and returns
// last inserted Id on success.
func AddGgcmsTasks(m *GgcmsTasks) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetGgcmsTasksById retrieves GgcmsTasks by Id. Returns error if
// Id doesn't exist
func GetGgcmsTasksById(id int) (v *GgcmsTasks, err error) {
	o := orm.NewOrm()
	v = &GgcmsTasks{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}
func GetCountGgcmsTasks(query map[string]string) (count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsTasks))
	var strs []string
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return 0, err
	}
	return qs.Count()
}
func ExistGgcmsTasks(query map[string]string) (exist bool) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsTasks))
	var strs []string
	var err error
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return false
	}
	return qs.Exist()
}

// GetAllGgcmsTasks retrieves all GgcmsTasks matches certain condition. Returns empty list if
// no records exist
func GetAllGgcmsTasks(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64, c bool) (ml []interface{}, count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsTasks))
	if qs, err = qsInit(&qs, query, sortby, order); err != nil {
		return nil, 0, err
	}
	count = 0
	var l []GgcmsTasks

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

// UpdateGgcmsTasks updates GgcmsTasks by Id and returns error if
// the record to be updated doesn't exist
func UpdateGgcmsTasksById(m *GgcmsTasks) (err error) {
	o := orm.NewOrm()
	v := GgcmsTasks{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteGgcmsTasks deletes GgcmsTasks by Id and returns error if
// the record to be deleted doesn't exist
func DeleteGgcmsTasks(id int) (err error) {
	o := orm.NewOrm()
	v := GgcmsTasks{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&GgcmsTasks{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
func MultDeleteGgcmsTasks(query map[string]string, ids []int) (num int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsTasks))
	var strs []string
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return 0, err
	}
	if len(ids) > 0 {
		qs = qs.Filter("id__in", ids)
	}
	return qs.Delete()
}
