package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type GgcmsModulesColumns struct {
	Id       int         `orm:"column(id);auto"`
	Tid      int         `orm:"column(tid)"`
	CName    string      `orm:"column(cName);size(255)"`
	CTitle   string      `orm:"column(cTitle);size(255)"`
	CType    string      `orm:"column(cType);size(50)"`
	CLen     int         `orm:"column(cLen)"`
	CDecimal int8        `orm:"column(cDecimal)"`
	CIndex   int         `orm:"column(cIndex)"`
	Options  string      `orm:"column(options)"`
	Value    interface{} `orm:"-"`
}

func (t *GgcmsModulesColumns) TableName() string {
	return "ggcms_modules_columns"
}

func init() {
	orm.RegisterModel(new(GgcmsModulesColumns))
}

// AddGgcmsModulesColumns insert a new GgcmsModulesColumns into database and returns
// last inserted Id on success.
func AddGgcmsModulesColumns(m *GgcmsModulesColumns) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetGgcmsModulesColumnsById retrieves GgcmsModulesColumns by Id. Returns error if
// Id doesn't exist
func GetGgcmsModulesColumnsById(id int) (v *GgcmsModulesColumns, err error) {
	o := orm.NewOrm()
	v = &GgcmsModulesColumns{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}
func GetCountGgcmsModulesColumns(query map[string]string) (count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsModulesColumns))
	var strs []string
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return 0, err
	}
	return qs.Count()
}
func ExistGgcmsModulesColumns(query map[string]string) (exist bool) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsModulesColumns))
	var strs []string
	var err error
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return false
	}
	return qs.Exist()
}

// GetAllGgcmsModulesColumns retrieves all GgcmsModulesColumns matches certain condition. Returns empty list if
// no records exist
func GetAllGgcmsModulesColumns(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64, c bool) (ml []GgcmsModulesColumns, count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsModulesColumns))
	if qs, err = qsInit(&qs, query, sortby, order); err != nil {
		return nil, 0, err
	}
	count = 0
	var l []GgcmsModulesColumns

	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if c {
			count, _ = qs.Count()
		}
		return l, count, nil
	}
	return nil, 0, err
}

// UpdateGgcmsModulesColumns updates GgcmsModulesColumns by Id and returns error if
// the record to be updated doesn't exist
func UpdateGgcmsModulesColumnsById(m *GgcmsModulesColumns) (err error) {
	o := orm.NewOrm()
	v := GgcmsModulesColumns{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m, "module_name", "module_descrip", "module_version"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteGgcmsModulesColumns deletes GgcmsModulesColumns by Id and returns error if
// the record to be deleted doesn't exist
func DeleteGgcmsModulesColumns(id int) (err error) {
	o := orm.NewOrm()
	v := GgcmsModulesColumns{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&GgcmsModulesColumns{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
func MultDeleteGgcmsModulesColumns(query map[string]string, ids []int) (num int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsModulesColumns))
	var strs []string
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return 0, err
	}
	if len(ids) > 0 {
		qs = qs.Filter("id__in", ids)
	}
	return qs.Delete()
}
