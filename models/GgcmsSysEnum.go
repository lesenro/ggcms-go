package models

import (
	"fmt"
	"reflect"

	"github.com/astaxie/beego/orm"
)

type GgcmsSysEnum struct {
	Id       int    `orm:"column(id);auto"`
	Ename    string `orm:"column(ename);size(30)"`
	Evalue   string `orm:"column(evalue);size(255)"`
	Egroup   string `orm:"column(egroup);size(20)"`
	Pvalue   string `orm:"column(pvalue);size(255)"`
	Orderid  uint   `orm:"column(orderid)"`
	Issystem int    `orm:"column(issystem)"`
}

func (t *GgcmsSysEnum) TableName() string {
	return "ggcms_sys_enum"
}

func init() {
	orm.RegisterModel(new(GgcmsSysEnum))
}

// AddGgcmsSysEnum insert a new GgcmsSysEnum into database and returns
// last inserted Id on success.
func AddGgcmsSysEnum(m *GgcmsSysEnum) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetGgcmsSysEnumById retrieves GgcmsSysEnum by Id. Returns error if
// Id doesn't exist
func GetGgcmsSysEnumById(id int) (v *GgcmsSysEnum, err error) {
	o := orm.NewOrm()
	v = &GgcmsSysEnum{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}
func GetCountGgcmsSysEnum(query map[string]string) (count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsSysEnum))
	var strs []string
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return 0, err
	}
	return qs.Count()
}
func ExistGgcmsSysEnum(query map[string]string) (exist bool) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsSysEnum))
	var strs []string
	var err error
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return false
	}
	return qs.Exist()
}
func GetGgcmsSysEnumGroups() (groups []string, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsSysEnum))
	var l []GgcmsSysEnum
	_, err = qs.GroupBy("egroup").All(&l, "egroup")
	for _, v := range l {
		groups = append(groups, v.Egroup)
	}
	return
}

// GetAllGgcmsSysEnum retrieves all GgcmsSysEnum matches certain condition. Returns empty list if
// no records exist
func GetAllGgcmsSysEnum(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64, c bool) (ml []interface{}, count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsSysEnum))
	if qs, err = qsInit(&qs, query, sortby, order); err != nil {
		return nil, 0, err
	}
	count = 0
	var l []GgcmsSysEnum

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

// UpdateGgcmsSysEnum updates GgcmsSysEnum by Id and returns error if
// the record to be updated doesn't exist
func UpdateGgcmsSysEnumById(m *GgcmsSysEnum) (err error) {
	o := orm.NewOrm()
	v := GgcmsSysEnum{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteGgcmsSysEnum deletes GgcmsSysEnum by Id and returns error if
// the record to be deleted doesn't exist
func DeleteGgcmsSysEnum(id int) (err error) {
	o := orm.NewOrm()
	v := GgcmsSysEnum{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&GgcmsSysEnum{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
func MultDeleteGgcmsSysEnum(query map[string]string, ids []int) (num int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsSysEnum))
	var strs []string
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return 0, err
	}
	if len(ids) > 0 {
		qs = qs.Filter("id__in", ids)
	}
	return qs.Delete()
}
