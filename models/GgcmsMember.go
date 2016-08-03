package models

import (
	"fmt"
	"reflect"

	"github.com/astaxie/beego/orm"
)

type GgcmsMember struct {
	Id              int    `orm:"column(id);auto"`
	Mtype           string `orm:"column(mtype);size(255)"`
	Userid          string `orm:"column(userid);size(20)"`
	Pwd             string `orm:"column(pwd);size(32)"`
	Uname           string `orm:"column(uname);size(36)"`
	Sex             int    `orm:"column(sex)"`
	Rank            int    `orm:"column(rank)"`
	Uptime          int64  `orm:"column(uptime)"`
	Exptime         int64  `orm:"column(exptime)"`
	Money           int    `orm:"column(money)"`
	Email           string `orm:"column(email);size(50)"`
	Scores          int    `orm:"column(scores)"`
	Matt            int    `orm:"column(matt)"`
	Spacesta        int    `orm:"column(spacesta)"`
	Face            string `orm:"column(face);size(50)"`
	Safequestion    int    `orm:"column(safequestion)"`
	Safeanswer      string `orm:"column(safeanswer);size(30)"`
	Jointime        int64  `orm:"column(jointime)"`
	Joinip          string `orm:"column(joinip);size(16)"`
	Logintime       int64  `orm:"column(logintime)"`
	Loginip         string `orm:"column(loginip);size(16)"`
	Checkmail       int    `orm:"column(checkmail)"`
	Openid          string `orm:"column(openid);size(100);null"`
	Userlevel       int    `orm:"column(userlevel);null"`
	Marketbind      string `orm:"column(marketbind);null"`
	Nativeplacecode string `orm:"column(nativeplacecode);size(30);null"`
}

func (t *GgcmsMember) TableName() string {
	return "ggcms_member"
}

func init() {
	orm.RegisterModel(new(GgcmsMember))
}

// AddGgcmsMember insert a new GgcmsMember into database and returns
// last inserted Id on success.
func AddGgcmsMember(m *GgcmsMember) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetGgcmsMemberById retrieves GgcmsMember by Id. Returns error if
// Id doesn't exist
func GetGgcmsMemberById(id int) (v *GgcmsMember, err error) {
	o := orm.NewOrm()
	v = &GgcmsMember{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}
func GetCountGgcmsMember(query map[string]string) (count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsMember))
	var strs []string
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return 0, err
	}
	return qs.Count()
}
func ExistGgcmsMember(query map[string]string) (exist bool) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsMember))
	var strs []string
	var err error
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return false
	}
	return qs.Exist()
}

// GetAllGgcmsMember retrieves all GgcmsMember matches certain condition. Returns empty list if
// no records exist
func GetAllGgcmsMember(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64, c bool) (ml []interface{}, count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsMember))
	if qs, err = qsInit(&qs, query, sortby, order); err != nil {
		return nil, 0, err
	}
	count = 0
	var l []GgcmsMember

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

// UpdateGgcmsMember updates GgcmsMember by Id and returns error if
// the record to be updated doesn't exist
func UpdateGgcmsMemberById(m *GgcmsMember) (err error) {
	o := orm.NewOrm()
	v := GgcmsMember{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteGgcmsMember deletes GgcmsMember by Id and returns error if
// the record to be deleted doesn't exist
func DeleteGgcmsMember(id int) (err error) {
	o := orm.NewOrm()
	v := GgcmsMember{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&GgcmsMember{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
func MultDeleteGgcmsMember(query map[string]string, ids []int) (num int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsMember))
	var strs []string
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return 0, err
	}
	if len(ids) > 0 {
		qs = qs.Filter("id__in", ids)
	}
	return qs.Delete()
}
