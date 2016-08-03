package models

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type GgcmsArticle struct {
	Id                int       `orm:"column(id);auto"`
	Content           string    `orm:"column(content);null"`
	Title             string    `orm:"column(title);size(255)"`
	Pagetitle         string    `orm:"column(pagetitle);size(255)"`
	Categoryid        int       `orm:"column(categoryid);null"`
	Nkey              string    `orm:"column(nkey);size(255);null"`
	Hits              int       `orm:"column(hits);null"`
	Dateandtime       time.Time `orm:"column(dateandtime);type(datetime);null"`
	DayHits           int       `orm:"column(dayHits);null"`
	WeekHits          int       `orm:"column(weekHits);null"`
	Isimage           int8      `orm:"column(isimage);null"`
	TitleThumbnailImg string    `orm:"column(titleThumbnailImg);size(255);null"`
	TitleImg          string    `orm:"column(titleImg);size(255);null"`
	Memberid          int       `orm:"column(memberid);null"`
	CheckUserid       int       `orm:"column(checkUserid);null"`
	AUrl              string    `orm:"column(aUrl);size(255);null"`
	Author            string    `orm:"column(author);size(255);null"`
	Source            string    `orm:"column(source);size(255);null"`
	SourceUrl         string    `orm:"column(sourceUrl);size(255);null"`
	Memberlimit       int       `orm:"column(memberlimit);null"`
	Keywords          string    `orm:"column(Keywords);size(255);null"`
	Description       string    `orm:"column(Description);size(255);null"`
	Recommendmode     int       `orm:"column(recommendmode);null"`
	Recommendlevel    int       `orm:"column(recommendlevel);null"`
	Tempname          string    `orm:"column(tempname);size(100)"`
	Styledir          string    `orm:"column(styledir);size(100)"`
	Siteid            int       `orm:"column(siteid);null"`
	Mid               int       `orm:"column(mid);null"`
}

func (t *GgcmsArticle) TableName() string {
	return "ggcms_article"
}

func init() {
	orm.RegisterModel(new(GgcmsArticle))
}

// AddGgcmsArticle insert a new GgcmsArticle into database and returns
// last inserted Id on success.
func AddGgcmsArticle(m *GgcmsArticle, pages *ArticlePages, attalist *[]GgcmsArticleAttr) (id int64, err error) {
	o := orm.NewOrm()
	o.Begin()
	m.Pagetitle = pages.Pages[0].Title
	m.Content = pages.Pages[0].Content
	id, err = o.Insert(m)
	if err != nil {
		o.Rollback()
		return
	}
	//多分页
	if len(pages.Pages) > 1 {
		//去掉第一页
		ps := append(pages.Pages[:0], pages.Pages[0+1:]...)
		for i, _ := range ps {
			ps[i].Articleid = int(id)
		}
		if _, err = o.InsertMulti(100, &ps); err != nil {
			o.Rollback()
			return
		}
	}
	//附件
	if len(*attalist) > 0 {
		for _, atta := range *attalist {
			//添加文章id
			atta.Articleid = int(id)
		}
		if _, err = o.InsertMulti(100, attalist); err != nil {
			o.Rollback()
			return
		}
	}
	o.Commit()
	return
}

// UpdateGgcmsArticle updates GgcmsArticle by Id and returns error if
// the record to be updated doesn't exist
func UpdateGgcmsArticleById(m *GgcmsArticle, pages *ArticlePages, ids []interface{}, attalist *[]GgcmsArticleAttr, modulesInfo *map[string]interface{}) (err error) {
	o := orm.NewOrm()
	v := GgcmsArticle{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		o.Begin()
		m.Pagetitle = pages.Pages[0].Title
		m.Content = pages.Pages[0].Content
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		} else {
			o.Rollback()
			return
		}
		//多分页
		if len(pages.Pages) > 1 {
			//去掉第一页
			ps := append(pages.Pages[:0], pages.Pages[0+1:]...)
			//页码变动
			for i, _ := range ps {
				pinfo := ps[i]
				if pinfo.OldSortid > 0 && pinfo.OldSortid != pinfo.Sortid {
					_, err = o.Raw("UPDATE `ggcms_article_pages` SET `sortid` = ? WHERE `articleid` = ? and `sortid` = ?").SetArgs(pinfo.Sortid, m.Id, pinfo.OldSortid).Exec()
					if err != nil {
						o.Rollback()
						return
					}
				}
			}
			for i, _ := range ps {
				pinfo := ps[i]
				if pinfo.Id > 0 { //有修改
					_, err = o.Update(&pinfo)
					if err != nil {
						o.Rollback()
						return
					}
				} else if pinfo.OldSortid == 0 { //新添加
					_, err = o.Insert(&pinfo)
					if err != nil {
						o.Rollback()
						return
					}
				}
			}
			//删除多余分页
			maxpage := ps[len(ps)-1].Sortid
			_, err = o.Raw("Delete From `ggcms_article_pages` WHERE `articleid` = ? and `sortid` > ?").SetArgs(m.Id, maxpage).Exec()
			if err != nil {
				o.Rollback()
				return
			}
			//要删除的分页
			if len(ids) > 0 {
				var sids []string
				for i := 0; i < len(ids); i++ {
					sids = append(sids, "?")
				}
				_, err = o.Raw("Delete From `ggcms_article_pages` WHERE `id` in  (" + strings.Join(sids, ",") + ")").SetArgs(ids...).Exec()
				if err != nil {
					o.Rollback()
					return
				}
			}
		} else {
			_, err = o.Raw("Delete From `ggcms_article_pages` WHERE `articleid` = ?").SetArgs(m.Id).Exec()
			if err != nil {
				o.Rollback()
				return
			}
		}
		//附件
		if len(*attalist) > 0 {
			addattalist := make([]GgcmsArticleAttr, 0)
			delids := make([]interface{}, 0)
			var sids []string

			for _, atta := range *attalist {
				//添加
				if atta.Id == 0 {
					//添加文章id
					atta.Articleid = m.Id
					addattalist = append(addattalist, atta)
				} else if atta.Id < 0 {
					//删除
					delids = append(delids, -atta.Id)
					sids = append(sids, "?")
				}
			}
			//添加
			if len(addattalist) > 0 {
				if _, err = o.InsertMulti(100, addattalist); err != nil {
					o.Rollback()
					return
				}
			}
			//删除
			if len(sids) > 0 {
				_, err = o.Raw("Delete From `ggcms_article_attr` WHERE `id` in  (" + strings.Join(sids, ",") + ")").SetArgs(delids...).Exec()
				if err != nil {
					o.Rollback()
					return
				}
			}
		}
		//模型信息
		if m.Mid > 0 {
			//模型信息有变动
			if v.Mid > 0 && m.Mid != v.Mid {
				var mod *GgcmsModules
				mod, err = GetGgcmsModulesById(v.Mid)
				if err == nil {
					//删除旧模型数据
					_, err = o.Raw("Delete From `" + mod.Table_name + "` WHERE `aid` = ?").SetArgs(m.Id).Exec()
					if err != nil {
						o.Rollback()
						return
					}
				}

			}
			var modinfo *GgcmsModules
			modinfo, err = GetGgcmsModulesById(m.Mid)
			if err != nil {
				o.Rollback()
				return
			}
			var mm []orm.Params
			_, err = o.Raw("select count(*) as mcount from `" + modinfo.Table_name + "` where aid = ?").SetArgs(m.Id).Values(&mm)
			mcount, _ := strconv.Atoi(mm[0]["mcount"].(string))
			if mcount == 0 {
				//添加
				var qmark, cols []string
				var vals []interface{}
				(*modulesInfo)["aid"] = m.Id
				for key, val := range *modulesInfo {
					qmark = append(qmark, "?")
					cols = append(cols, "`"+key+"`")
					vals = append(vals, val)
				}
				_, err = o.Raw("INSERT INTO `" + modinfo.Table_name + "` (" + strings.Join(cols, ",") + ") VALUES (" + strings.Join(qmark, ",") + ")").SetArgs(vals...).Exec()
				if err != nil {
					o.Rollback()
					return
				}
			} else {
				//修改
				var cols []string
				var vals []interface{}
				for key, val := range *modulesInfo {
					cols = append(cols, "`"+key+"` = ?")
					vals = append(vals, val)
					beego.Debug(key, val)
				}
				vals = append(vals, m.Id)
				_, err = o.Raw("UPDATE `" + modinfo.Table_name + "` SET " + strings.Join(cols, ",") + " WHERE aid = ?").SetArgs(vals...).Exec()
				if err != nil {
					o.Rollback()
					return
				}
			}
		}
		o.Commit()
	}
	return
}

// GetGgcmsArticleById retrieves GgcmsArticle by Id. Returns error if
// Id doesn't exist
func GetGgcmsArticleById(id int) (v *GgcmsArticle, err error) {
	o := orm.NewOrm()
	v = &GgcmsArticle{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}
func GetCountGgcmsArticle(query map[string]string) (count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsArticle))
	var strs []string
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return 0, err
	}
	return qs.Count()
}
func ExistGgcmsArticle(query map[string]string) (exist bool) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsArticle))
	var strs []string
	var err error
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return false
	}
	return qs.Exist()
}

// GetAllGgcmsArticle retrieves all GgcmsArticle matches certain condition. Returns empty list if
// no records exist
func GetAllGgcmsArticle(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64, c bool) (ml []interface{}, count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsArticle))
	if qs, err = qsInit(&qs, query, sortby, order); err != nil {
		return nil, 0, err
	}
	count = 0
	var l []GgcmsArticle

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

// DeleteGgcmsArticle deletes GgcmsArticle by Id and returns error if
// the record to be deleted doesn't exist
func DeleteGgcmsArticle(id int) (err error) {
	o := orm.NewOrm()
	v := GgcmsArticle{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&GgcmsArticle{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
func MultDeleteGgcmsArticle(query map[string]string, ids []int) (num int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GgcmsArticle))
	var strs []string
	if qs, err = qsInit(&qs, query, strs, strs); err != nil {
		return 0, err
	}
	if len(ids) > 0 {
		qs = qs.Filter("id__in", ids)
	}
	return qs.Delete()
}
