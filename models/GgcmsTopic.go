package models

import (
	"fmt"
	"reflect"
	"time"

	"github.com/astaxie/beego/orm"
)

type GgcmsTopic struct {
	Id          int       `orm:"column(id);auto"`
	Topic       string    `orm:"column(topic);size(255);null"`
	Dateandtime time.Time `orm:"column(dateandtime);type(datetime);null"`
	Articlenum  int       `orm:"column(articlenum);null"`
	Categoryid  int       `orm:"column(categoryid);null"`
	Content     string    `orm:"column(content);null"`
	Pagesize    int       `orm:"column(pagesize);null"`
	Tempname    string    `orm:"column(tempname);size(100)"`
	Styledir    string    `orm:"column(styledir);size(100)"`
	Subtopic    string    `orm:"column(subtopic);null"`
	Title       string    `orm:"column(title);size(255);null"`
	Description string    `orm:"column(Description);size(255);null"`
	Keywords    string    `orm:"column(Keywords);size(255);null"`
	Logo        string    `orm:"column(logo);size(255);null"`
	Turl        string    `orm:"column(turl);size(255);null"`
	Extattrib   string    `orm:"column(extattrib);size(255);null"`
}

func (t *GgcmsTopic) TableName() string {
	return "ggcms_topic"
}

func init() {
	orm.RegisterModel(new(GgcmsTopic))
}

// AddGgcmsTopic insert a new GgcmsTopic into database and returns
// last inserted Id on success.
func AddGgcmsTopic(m *GgcmsTopic) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetGgcmsTopicById retrieves GgcmsTopic by Id. Returns error if
// Id doesn't exist
func GetGgcmsTopicById(id int) (v *GgcmsTopic, err error) {
	o := orm.NewOrm()
	v = &GgcmsTopic{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}
func GetCountGgcmsTopic(query map[string]string) (count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsTopic))
	var strs []string
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return 0, err
	}
	return qs.Count()
}
func ExistGgcmsTopic(query map[string]string) (exist bool) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsTopic))
	var strs []string
	var err error
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return false
	}
	return qs.Exist()
}

// GetAllGgcmsTopic retrieves all GgcmsTopic matches certain condition. Returns empty list if
// no records exist
func GetAllGgcmsTopic(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64, c bool) (ml []interface{}, count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsTopic))
	if qs, err = qsInit(&qs, query, sortby, order); err != nil {
		return nil, 0, err
	}
	count = 0
	var l []GgcmsTopic

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

// UpdateGgcmsTopic updates GgcmsTopic by Id and returns error if
// the record to be updated doesn't exist
func UpdateGgcmsTopicById(m *GgcmsTopic) (err error) {
	o := orm.NewOrm()
	v := GgcmsTopic{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteGgcmsTopic deletes GgcmsTopic by Id and returns error if
// the record to be deleted doesn't exist
func DeleteGgcmsTopic(id int) (err error) {
	o := orm.NewOrm()
	v := GgcmsTopic{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&GgcmsTopic{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
func MultDeleteGgcmsTopic(query map[string]string, ids []int) (num int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsTopic))
	var strs []string
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return 0, err
	}
	if len(ids) > 0 {
		qs = qs.Filter("id__in", ids)
	}
	return qs.Delete()
}
