package models

import (
	"fmt"
	"reflect"
	"time"

	"github.com/astaxie/beego/orm"
)

type GgcmsSysLogs struct {
	Id            int       `orm:"column(id);auto"`
	Msg           string    `orm:"column(msg);size(255);null"`
	Flg           string    `orm:"column(flg);size(30);null"`
	Logtime       time.Time `orm:"column(logtime);type(datetime);null"`
	Userid        string    `orm:"column(userid);size(30);null"`
	Userip        string    `orm:"column(userip);size(20);null"`
	Logtype       int       `orm:"column(logtype);null"`
	Requestparams string    `orm:"column(requestparams);null"`
}

func (t *GgcmsSysLogs) TableName() string {
	return "ggcms_sys_logs"
}

func init() {
	orm.RegisterModel(new(GgcmsSysLogs))
}

// AddGgcmsSysLogs insert a new GgcmsSysLogs into database and returns
// last inserted Id on success.
func AddGgcmsSysLogs(m *GgcmsSysLogs) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetGgcmsSysLogsById retrieves GgcmsSysLogs by Id. Returns error if
// Id doesn't exist
func GetGgcmsSysLogsById(id int) (v *GgcmsSysLogs, err error) {
	o := orm.NewOrm()
	v = &GgcmsSysLogs{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}
func GetCountGgcmsSysLogs(query map[string]string) (count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsSysLogs))
	var strs []string
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return 0, err
	}
	return qs.Count()
}
func ExistGgcmsSysLogs(query map[string]string) (exist bool) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsSysLogs))
	var strs []string
	var err error
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return false
	}
	return qs.Exist()
}

// GetAllGgcmsSysLogs retrieves all GgcmsSysLogs matches certain condition. Returns empty list if
// no records exist
func GetAllGgcmsSysLogs(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64, c bool) (ml []interface{}, count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsSysLogs))
	if qs, err = qsInit(&qs, query, sortby, order); err != nil {
		return nil, 0, err
	}
	count = 0
	var l []GgcmsSysLogs

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

// UpdateGgcmsSysLogs updates GgcmsSysLogs by Id and returns error if
// the record to be updated doesn't exist
func UpdateGgcmsSysLogsById(m *GgcmsSysLogs) (err error) {
	o := orm.NewOrm()
	v := GgcmsSysLogs{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteGgcmsSysLogs deletes GgcmsSysLogs by Id and returns error if
// the record to be deleted doesn't exist
func DeleteGgcmsSysLogs(id int) (err error) {
	o := orm.NewOrm()
	v := GgcmsSysLogs{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&GgcmsSysLogs{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
func MultDeleteGgcmsSysLogs(query map[string]string, ids []int) (num int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsSysLogs))
	var strs []string
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return 0, err
	}
	if len(ids) > 0 {
		qs = qs.Filter("id__in", ids)
	}
	return qs.Delete()
}
