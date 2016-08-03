package models

import (
	"fmt"
	"reflect"
	"strconv"

	"strings"

	"github.com/astaxie/beego/orm"
)

type GgcmsModulesPostData struct {
	GgcmsModules
	Columns []GgcmsModulesColumns
}
type GgcmsModules struct {
	Id             int    `orm:"column(id);auto"`
	Module_name    string `orm:"column(module_name);size(255)"`
	Module_descrip string `orm:"column(module_descrip);size(255)"`
	Module_version string `orm:"column(module_version);size(255)"`
	Table_name     string `orm:"column(table_name);size(50)"`
	View_name      string `orm:"column(view_name);size(50)"`
}

func (t *GgcmsModules) TableName() string {
	return "ggcms_modules"
}

func init() {
	orm.RegisterModel(new(GgcmsModules))
}
func addColumnsSql(col GgcmsModulesColumns, id, i int) (sql string) {
	switch col.CType {
	case "varchar":
		sql = " `" + col.CName + "` varchar(255) DEFAULT '' COMMENT '" + col.CTitle + "', "
	case "int":
		sql = " `" + col.CName + "` int(11) DEFAULT '0' COMMENT '" + col.CTitle + "', "
	case "text":
		sql = " `" + col.CName + "` text COMMENT '" + col.CTitle + "', "
	case "datetime":
		sql = " `" + col.CName + "` datetime DEFAULT NULL COMMENT '" + col.CTitle + "', "
	case "decimal":
		sql = " `" + col.CName + "` decimal(10," + strconv.Itoa(int(col.CDecimal)) + ") DEFAULT '0' COMMENT '" + col.CTitle + "', "
	}
	return
}

// AddGgcmsModules insert a new GgcmsModules into database and returns
// last inserted Id on success.
func AddGgcmsModules(m *GgcmsModulesPostData) (id int64, err error) {
	o := orm.NewOrm()
	o.Raw("DROP TABLE IF EXISTS `" + m.Table_name + "`").Exec()
	o.Raw("DROP VIEW IF EXISTS `" + m.View_name + "`").Exec()
	o.Begin()
	if id, err = o.Insert(&m.GgcmsModules); err != nil {
		o.Rollback()
		return
	}
	var cols []GgcmsModulesColumns
	var sqltmp []string
	for i, col := range m.Columns {
		col.Tid = int(id)
		col.CName = "cn_" + strconv.Itoa(int(id)) + "_" + strconv.Itoa(i)
		col.CIndex = i
		cols = append(cols, col)
		sqltmp = append(sqltmp, addColumnsSql(col, int(id), i))
	}

	if _, err = o.InsertMulti(100, &cols); err != nil {
		o.Rollback()
		return
	}
	sql := "CREATE TABLE `" + m.Table_name + "` ( " +
		"`tid` int(11) NOT NULL AUTO_INCREMENT," +
		"`aid` int(11) DEFAULT '0' ," +
		"`sid` int(11) DEFAULT '0' ," +
		strings.Join(sqltmp, "") +
		"PRIMARY KEY (`tid`)" +
		") DEFAULT CHARSET=utf8"
	if _, err = o.Raw(sql).Exec(); err != nil {
		o.Rollback()
		return
	}
	sql = "CREATE VIEW `" + m.View_name + "`AS " +
		"SELECT a.*, b.* FROM ggcms_article a " +
		"INNER JOIN " + m.Table_name + " b ON a.id = b.aid ;"
	if _, err = o.Raw(sql).Exec(); err != nil {
		o.Rollback()
		return
	}
	o.Commit()
	return
}

// UpdateGgcmsModules updates GgcmsModules by Id and returns error if
// the record to be updated doesn't exist
func UpdateGgcmsModulesById(m *GgcmsModulesPostData) (err error) {
	o := orm.NewOrm()
	// ascertain id exists in the database
	v := GgcmsModules{Id: m.Id}
	if err = o.Read(&v); err == nil {
		var strs []string
		query := make(map[string]string)
		query["tid"] = strconv.Itoa(m.Id)
		//edit
		oldCols, _, _ := GetAllGgcmsModulesColumns(query, strs, strs, strs, 0, 0, false)
		maxid := 0
		//最大table,id号
		for _, col := range oldCols {
			if col.CIndex > maxid {
				maxid = col.CIndex
			}
		}
		var cols []GgcmsModulesColumns
		var sqltmp []string
		var editcols []GgcmsModulesColumns
		//检查列的修改和添加
		for _, col := range m.Columns {
			if col.Id > 0 {
				for _, ocol := range oldCols {
					if col.CName == ocol.CName {
						if col.CType != ocol.CType || col.CDecimal != ocol.CDecimal {
							sqltmp = append(sqltmp, " MODIFY COLUMN "+addColumnsSql(col, m.Id, col.CIndex))
							editcols = append(editcols, col)
						} else if col.CDecimal != ocol.CDecimal || col.Options != ocol.Options {
							editcols = append(editcols, col)
						}
					}
				}
			} else {
				maxid++
				col.Tid = int(m.Id)
				col.CName = "cn_" + strconv.Itoa(int(m.Id)) + "_" + strconv.Itoa(maxid)
				col.CIndex = maxid
				cols = append(cols, col)
				sqltmp = append(sqltmp, " ADD COLUMN "+addColumnsSql(col, m.Id, maxid))
			}
		}
		o.Begin()
		//改表
		if len(sqltmp) > 0 {
			sql := "ALTER TABLE `" + m.Table_name + "` "
			sql = sql + strings.Join(sqltmp, "")
			sql = strings.Trim(sql, " ")
			sql = strings.TrimRight(sql, ",")
			if _, err = o.Raw(sql).Exec(); err != nil {
				o.Rollback()
				return
			}
			//改视图
			if _, err = o.Raw("DROP VIEW IF EXISTS `" + m.View_name + "`").Exec(); err != nil {
				o.Rollback()
				return
			}
			sql = "CREATE VIEW `" + m.View_name + "`AS " +
				"SELECT a.*, b.* FROM ggcms_article a " +
				"INNER JOIN " + m.Table_name + " b ON a.id = b.aid ;"
			if _, err = o.Raw(sql).Exec(); err != nil {
				o.Rollback()
				return
			}
		}
		//改记录
		if len(editcols) > 0 {
			p, err := o.Raw("UPDATE `ggcms_modules_columns` SET cTitle = ?, cType = ?, cDecimal = ?, Options = ? WHERE `id` = ?").Prepare()
			for _, col := range editcols {
				if _, err = p.Exec(col.CTitle, col.CType, col.CDecimal, col.Options, col.Id); err != nil {
					o.Rollback()
					p.Close()
					return err
				}
			}

			p.Close()
		}
		//加记录
		if len(cols) > 0 {
			if _, err = o.InsertMulti(100, &cols); err != nil {
				o.Rollback()
				return
			}
		}
		//删除记录
		var delids []int
		for _, ocol := range oldCols {
			isdel := true
			for _, col := range m.Columns {
				if col.CName == ocol.CName {
					isdel = false
				}
			}
			if isdel {
				delids = append(delids, ocol.Id)
				sqltmp = append(sqltmp, " DROP COLUMN `"+ocol.CName+"` , ")
			}
		}
		//删除记录
		if len(delids) > 0 {
			var query map[string]string
			if _, err = MultDeleteGgcmsModulesColumns(query, delids); err != nil {
				o.Rollback()
				return
			}
		}

		if num, err := o.Update(&m.GgcmsModules, "module_name", "module_descrip", "module_version"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		} else {
			o.Rollback()
			return err
		}
		o.Commit()
	}

	return
}

// GetGgcmsModulesById retrieves GgcmsModules by Id. Returns error if
// Id doesn't exist
func GetGgcmsModulesById(id int) (v *GgcmsModules, err error) {
	o := orm.NewOrm()
	v = &GgcmsModules{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}
func GetGgcmsModulesInfo(aid int, tab_name string) (maps *[]orm.Params, err error) {
	o := orm.NewOrm()
	var mmaps []orm.Params
	_, err = o.Raw("SELECT * FROM `"+tab_name+"` WHERE aid = ?", aid).Values(&mmaps)
	if err == nil {
		return &mmaps, nil
	}
	return nil, err
}
func GetCountGgcmsModules(query map[string]string) (count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsModules))
	var strs []string
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return 0, err
	}
	return qs.Count()
}
func ExistGgcmsModules(query map[string]string) (exist bool) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsModules))
	var strs []string
	var err error
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return false
	}
	return qs.Exist()
}

// GetAllGgcmsModules retrieves all GgcmsModules matches certain condition. Returns empty list if
// no records exist
func GetAllGgcmsModules(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64, c bool) (ml []interface{}, count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsModules))
	if qs, err = qsInit(&qs, query, sortby, order); err != nil {
		return nil, 0, err
	}
	count = 0
	var l []GgcmsModules
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

// DeleteGgcmsModules deletes GgcmsModules by Id and returns error if
// the record to be deleted doesn't exist
func DeleteGgcmsModules(id int) (err error) {
	o := orm.NewOrm()
	v := GgcmsModules{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		o.Begin()
		if _, err = o.Raw("DROP TABLE IF EXISTS `" + v.Table_name + "`").Exec(); err != nil {
			o.Rollback()
			return
		}
		if _, err = o.Raw("DROP VIEW IF EXISTS `" + v.View_name + "`").Exec(); err != nil {
			o.Rollback()
			return
		}
		if _, err = o.Raw("DELETE FROM `ggcms_modules_columns` WHERE `tid` = ?", id).Exec(); err != nil {
			o.Rollback()
			return
		}
		if num, err = o.Delete(&GgcmsModules{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		} else {
			o.Rollback()
			return
		}
		o.Commit()
	}
	return
}
func MultDeleteGgcmsModules(query map[string]string, ids []int) (num int64, err error) {
	for i := 0; i < len(ids); i++ {
		if err = DeleteGgcmsModules(ids[i]); err != nil {
			break
		}
		num++
	}
	return
}
