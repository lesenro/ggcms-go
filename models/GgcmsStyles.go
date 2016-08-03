package models

import (
	"fmt"
	"reflect"

	"github.com/astaxie/beego/orm"
)

type GgcmsStyles struct {
	Id            int    `orm:"column(id);auto"`
	Style_name    string `orm:"column(style_name);size(255);null"`
	Style_folder  string `orm:"column(style_folder);size(255);null"`
	Style_descrip string `orm:"column(style_descrip);size(255);null"`
}

func (t *GgcmsStyles) TableName() string {
	return "ggcms_styles"
}

func init() {
	orm.RegisterModel(new(GgcmsStyles))
}

// AddGgcmsStyles insert a new GgcmsStyles into database and returns
// last inserted Id on success.
func AddGgcmsStyles(m *GgcmsStyles) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetGgcmsStylesById retrieves GgcmsStyles by Id. Returns error if
// Id doesn't exist
func GetGgcmsStylesById(id int) (v *GgcmsStyles, err error) {
	o := orm.NewOrm()
	v = &GgcmsStyles{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}
func GetCountGgcmsStyles(query map[string]string) (count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsStyles))
	var strs []string
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return 0, err
	}
	return qs.Count()
}
func ExistGgcmsStyles(query map[string]string) (exist bool) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsStyles))
	var strs []string
	var err error
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return false
	}
	return qs.Exist()
}

// GetAllGgcmsStyles retrieves all GgcmsStyles matches certain condition. Returns empty list if
// no records exist
func GetAllGgcmsStyles(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64, c bool) (ml []interface{}, count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsStyles))
	if qs, err = qsInit(&qs, query, sortby, order); err != nil {
		return nil, 0, err
	}
	count = 0
	var l []GgcmsStyles

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

// UpdateGgcmsStyles updates GgcmsStyles by Id and returns error if
// the record to be updated doesn't exist
func UpdateGgcmsStylesById(m *GgcmsStyles) (err error) {
	o := orm.NewOrm()
	v := GgcmsStyles{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteGgcmsStyles deletes GgcmsStyles by Id and returns error if
// the record to be deleted doesn't exist
func DeleteGgcmsStyles(id int) (err error) {
	o := orm.NewOrm()
	v := GgcmsStyles{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&GgcmsStyles{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
func MultDeleteGgcmsStyles(query map[string]string, ids []int) (num int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsStyles))
	var strs []string
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return 0, err
	}
	if len(ids) > 0 {
		qs = qs.Filter("id__in", ids)
	}
	return qs.Delete()
}
