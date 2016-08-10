package models

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/astaxie/beego/orm"
)

type GgcmsSites struct {
	Id           int    `orm:"column(id);auto"`
	Site_name    string `orm:"column(site_name);size(30)"`
	Site_domain  string `orm:"column(site_domain);size(255)"`
	Site_descrip string `orm:"column(site_descrip);size(255)"`
	Site_main    int8   `orm:"column(site_main);null"`
}

func (t *GgcmsSites) TableName() string {
	return "ggcms_sites"
}

func init() {
	orm.RegisterModel(new(GgcmsSites))
}

// AddGgcmsSites insert a new GgcmsSites into database and returns
// last inserted Id on success.
func AddGgcmsSites(m *GgcmsSites) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetGgcmsSitesById retrieves GgcmsSites by Id. Returns error if
// Id doesn't exist
func GetGgcmsSitesById(id int) (v *GgcmsSites, err error) {
	o := orm.NewOrm()
	v = &GgcmsSites{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetGgcmsSitesById retrieves GgcmsSites by Id. Returns error if
// Id doesn't exist
func GetGgcmsSitesByDomain(dm string) (v *GgcmsSites, err error) {
	o := orm.NewOrm()

	var l []GgcmsSites
	_, err = o.Raw("SELECT * FROM `ggcms_sites` WHERE site_domain = ? OR site_main = 1 ORDER BY site_main", dm).QueryRows(&l)
	if err != nil {
		return nil, err
	}
	if len(l) == 0 {
		return nil, errors.New("not found")
	}
	return &l[0], nil
}
func GetCountGgcmsSites(query map[string]string) (count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsSites))
	var strs []string
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return 0, err
	}
	return qs.Count()
}
func ExistGgcmsSites(query map[string]string) (exist bool) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsSites))
	var strs []string
	var err error
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return false
	}
	return qs.Exist()
}

// GetAllGgcmsSites retrieves all GgcmsSites matches certain condition. Returns empty list if
// no records exist
func GetAllGgcmsSites(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64, c bool) (ml []interface{}, count int64, err error) {
	o := orm.NewOrm()

	qs := o.QueryTable(new(GgcmsSites))

	if qs, err = qsInit(&qs, query, sortby, order); err != nil {
		return nil, 0, err
	}
	count = 0
	var l []GgcmsSites

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

// UpdateGgcmsSites updates GgcmsSites by Id and returns error if
// the record to be updated doesn't exist
func UpdateGgcmsSitesById(m *GgcmsSites) (err error) {
	o := orm.NewOrm()
	v := GgcmsSites{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteGgcmsSites deletes GgcmsSites by Id and returns error if
// the record to be deleted doesn't exist
func DeleteGgcmsSites(id int) (err error) {
	o := orm.NewOrm()
	v := GgcmsSites{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&GgcmsSites{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func MultDeleteGgcmsSites(query map[string]string, ids []int) (num int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsSites))

	var strs []string
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return 0, err
	}
	if len(ids) > 0 {
		qs = qs.Filter("id__in", ids)
	}
	return qs.Delete()
}
func UpdateConfigGgcmsSites(id int) (err error) {
	mb := GetSystemConfigsBySiteId(-1)
	cfgs := GetSystemConfigsBySiteId(id)
	//删除多余配置
	var ids []int
	for _, v := range cfgs {
		vv := v.(GgcmsSystemConfigs)
		del := true
		for _, m := range mb {
			mm := m.(GgcmsSystemConfigs)
			if mm.Varname == vv.Varname {
				del = false
				break
			}
		}
		if del {
			ids = append(ids, vv.Id)
		}
	}
	if len(ids) > 0 {
		_, err = MultDeleteGgcmsSystemConfigs(make(map[string]string), ids)
		if err != nil {
			return
		}
	}
	//添加新增配置
	o := orm.NewOrm()
	var newcfgs []GgcmsSystemConfigs
	for _, m := range mb {
		mm := m.(GgcmsSystemConfigs)
		add := true
		for _, v := range cfgs {
			vv := v.(GgcmsSystemConfigs)
			if mm.Varname == vv.Varname {
				add = false
				vv.Info = mm.Info
				vv.Groupid = mm.Groupid
				vv.Others = mm.Others
				vv.Sortid = mm.Sortid
				o.Update(&vv, "Info", "Groupid", "Others", "Sortid")
				break
			}
		}
		if add {
			mm.Siteid = id
			newcfgs = append(newcfgs, mm)
		}
	}
	if len(newcfgs) > 0 {
		_, err = o.InsertMulti(100, &newcfgs)
	}
	return
}
