package models

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type GgcmsSystemConfigs struct {
	Id      int    `orm:"column(id);auto"`
	Siteid  int    `orm:"column(siteid);null"`
	Varname string `orm:"column(varname);size(20);null"`
	Info    string `orm:"column(info);size(100);null"`
	Groupid int16  `orm:"column(groupid);null"`
	Value   string `orm:"column(value);null"`
	Others  string `orm:"column(others);null"`
	Sortid  int    `orm:"column(sortid);null"`
}

func (t *GgcmsSystemConfigs) TableName() string {
	return "ggcms_sysconfig"
}

func init() {
	orm.RegisterModel(new(GgcmsSystemConfigs))
}

// AddGgcmsSystemConfigs insert a new GgcmsSystemConfigs into database and returns
// last inserted Id on success.
func AddGgcmsSystemConfigs(m *GgcmsSystemConfigs) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetGgcmsSystemConfigsById retrieves GgcmsSystemConfigs by Id. Returns error if
// Id doesn't exist
func GetGgcmsSystemConfigsById(id int) (v *GgcmsSystemConfigs, err error) {
	o := orm.NewOrm()
	v = &GgcmsSystemConfigs{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}
func GetCountGgcmsSystemConfigs(query map[string]string) (count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsSystemConfigs))
	var strs []string
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return 0, err
	}
	return qs.Count()
}
func ExistGgcmsSystemConfigs(query map[string]string) (exist bool) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsSystemConfigs))
	var strs []string
	var err error
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return false
	}
	return qs.Exist()
}

// GetAllGgcmsSystemConfigs retrieves all GgcmsSystemConfigs matches certain condition. Returns empty list if
// no records exist
func GetAllGgcmsSystemConfigs(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64, c bool) (ml []interface{}, count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsSystemConfigs))
	if qs, err = qsInit(&qs, query, sortby, order); err != nil {
		return nil, 0, err
	}
	count = 0
	var l []GgcmsSystemConfigs

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

// UpdateGgcmsSystemConfigs updates GgcmsSystemConfigs by Id and returns error if
// the record to be updated doesn't exist
func UpdateGgcmsSystemConfigsById(m *GgcmsSystemConfigs) (err error) {
	o := orm.NewOrm()
	v := GgcmsSystemConfigs{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m, "value"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}
func MultUpdateGgcmsSystemConfig(v map[string]GgcmsSystemConfigs) (err error) {
	o := orm.NewOrm()
	o.Begin()
	//改记录
	if len(v) > 0 {
		p, err := o.Raw("UPDATE `ggcms_sysconfig` SET Value = ? WHERE `id` = ?").Prepare()
		for _, val := range v {
			if _, err = p.Exec(val.Value, val.Id); err != nil {
				o.Rollback()
				p.Close()
				return err
			}
		}
		p.Close()
	}
	o.Commit()
	return
}

// DeleteGgcmsSystemConfigs deletes GgcmsSystemConfigs by Id and returns error if
// the record to be deleted doesn't exist
func DeleteGgcmsSystemConfigs(id int) (err error) {
	o := orm.NewOrm()
	v := GgcmsSystemConfigs{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&GgcmsSystemConfigs{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
func MultDeleteGgcmsSystemConfigs(query map[string]string, ids []int) (num int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsSystemConfigs))
	var strs []string
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return 0, err
	}
	if len(ids) > 0 {
		qs = qs.Filter("id__in", ids)
	}
	return qs.Delete()
}

func GetSystemConfigsBySiteId(sid int) []interface{} {
	var strs []string
	sortstr := []string{"sortid"}
	orderstr := []string{"asc"}
	qmap := make(map[string]string)
	qmap["siteid"] = strconv.Itoa(sid)
	ml, _, _ := GetAllGgcmsSystemConfigs(qmap, strs, sortstr, orderstr, 0, 0, false)
	return ml
}
