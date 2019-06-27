package dao

import (
	"bytes"
	"errors"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/jmoiron/sqlx"
	"strings"
)

var SetmealDao = setmeal{}

type setmeal struct{}

// 根据【套餐ID】查询【套餐表】表中是否存在相关记录
func (d *setmeal) Exist(setmealid int) (bool, error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from setmeal where setmealid=?", setmealid)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	count := 0
	if rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return false, err
		}
		return count > 0, nil
	}
	return false, nil
}

// 插入单条记录到【套餐表】表中
func (d *setmeal) Insert(m *model.Setmeal) (int64, error) {
	result, err := garglewool.Exec("insert into setmeal(shopid,pkgcode,setmealtype,title,people,isorder,isrefund,isinhouse,isnew,isrecommend,validdays,watchers,validtime,exittime,links,info,adder,addtime,moder,modtime,deleteStatus) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", m.Shopid, m.Pkgcode, m.Setmealtype, m.Title, m.People, m.Isorder, m.Isrefund, m.Isinhouse, m.Isnew, m.Isrecommend, m.Validdays, m.Watchers, m.Validtime, m.Exittime, m.Links, m.Info, m.Adder, m.Addtime, m.Moder, m.Modtime, m.DeleteStatus)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

// 根据【套餐ID】修改【套餐表】表的单条记录
func (d *setmeal) Update(m *model.Setmeal) (bool, error) {
	result, err := garglewool.Exec("update setmeal set shopid=?, pkgcode=?, setmealtype=?, title=?, people=?, isorder=?, isrefund=?, isinhouse=?, isnew=?, isrecommend=?, validdays=?, watchers=?, validtime=?, exittime=?, links=?, info=?, adder=?, addtime=?, moder=?, modtime=?, deleteStatus=? where setmealid=?", m.Shopid, m.Pkgcode, m.Setmealtype, m.Title, m.People, m.Isorder, m.Isrefund, m.Isinhouse, m.Isnew, m.Isrecommend, m.Validdays, m.Watchers, m.Validtime, m.Exittime, m.Links, m.Info, m.Adder, m.Addtime, m.Moder, m.Modtime, m.DeleteStatus, m.Setmealid)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【套餐ID】软删除【套餐表】表中的单条记录
func (d *setmeal) Delete(setmealid int) (bool, error) {
	result, err := garglewool.Exec("update setmeal set deleteStatus=2 where setmealid=?", setmealid)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【套餐ID】数组软删除【套餐表】表中的多条记录
func (d *setmeal) DeleteIn(setmealids []int) (count int64, err error) {
	if len(setmealids) <= 0 {
		return count, errors.New("setmealids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("update setmeal set deleteStatus=2")
	sql_str.WriteString(" where setmealid in(")
	question_mark := strings.Repeat("?,", len(setmealids))
	sql_str.WriteString(question_mark[:len(question_mark)-1])
	sql_str.WriteString(")")
	vals := make([]interface{}, 0, len(setmealids))
	for _, v := range setmealids {
		vals = append(vals, v)
	}
	result, err := garglewool.Exec(sql_str.String(), vals...)
	if err != nil {
		return count, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return count, err
	}
	return affected, nil
}

// 根据【套餐ID】查询【套餐表】表中的单条记录
func (d *setmeal) Get(setmealid int) (setmeal model.Setmeal, err error) {
	rows, err := garglewool.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ setmealid, shopid, pkgcode, setmealtype, title, people, isorder, isrefund, isinhouse, isnew, isrecommend, validdays, watchers, createtime, validtime, exittime, links, info, adder, addtime, moder, modtime, deleteStatus from setmeal where setmealid=?", setmealid)
	if err != nil {
		return setmeal, err
	}
	defer rows.Close()
	setmeals, err := d._RowsToArray(rows)
	if err != nil {
		return setmeal, err
	}
	if len(setmeals) <= 0 {
		return setmeal, err
	}
	return setmeals[0], nil
}

// 根据【套餐ID】数组查询【套餐表】表中的多条记录
func (d *setmeal) GetIn(setmealids []int) (setmeals []model.Setmeal, err error) {
	if len(setmealids) <= 0 {
		return setmeals, errors.New("setmealids is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("select /*+ MAX_EXECUTION_TIME(5000) */ setmealid, shopid, pkgcode, setmealtype, title, people, isorder, isrefund, isinhouse, isnew, isrecommend, validdays, watchers, createtime, validtime, exittime, links, info, adder, addtime, moder, modtime, deleteStatus from ")
	sql_str.WriteString("setmeal")
	sql_str.WriteString(" where setmealid in(")
	param_keys := strings.Repeat("?,", len(setmealids))
	sql_str.WriteString(param_keys[:len(param_keys)-1])
	sql_str.WriteString(")")
	var rows *sqlx.Rows
	vals := make([]interface{}, 0, len(setmealids))
	for _, v := range setmealids {
		vals = append(vals, v)
	}
	rows, err = garglewool.Queryx(sql_str.String(), vals...)
	if err != nil {
		return setmeals, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 查询【套餐表】表总记录数
func (d *setmeal) GetRowCount() (count int, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from setmeal")

	//region 处理deleteStatus
	if conditions > 0 {
		sqlString.Append(" and ")
	} else {
		sqlString.Append(" where ")
	}
	sqlString.Append("deleteStatus = 1")
	//endregion

	//region Query
	rows, err := garglewool.Queryx(sqlString.ToString(), params...)
	if err != nil {
		return -1, err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return -1, err
		}
		return count, nil
	}
	//endregion

	return -1, nil
}

// 根据【店铺ID】查询【套餐表】总记录数
func (d *setmeal) GetRowCountByShopid(shopid int) (count int, err error) {
	rows, err := garglewool.Queryx("select count(0) Count from setmeal where shopid=? and deleteStatus = 1", shopid)
	if err != nil {
		return -1, err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return -1, err
		}
		return count, nil
	}
	return -1, err
}

// 查询【套餐表】列表
func (d *setmeal) GetRowList(pageIndex, pageSize int) (setmeals []model.Setmeal, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ setmealid, shopid, pkgcode, setmealtype, title, people, isorder, isrefund, isinhouse, isnew, isrecommend, validdays, watchers, createtime, validtime, exittime, links, info, adder, addtime, moder, modtime, deleteStatus from setmeal")

	//region 处理deleteStatus
	if conditions > 0 {
		sqlString.Append(" and ")
	} else {
		sqlString.Append(" where ")
	}
	sqlString.Append("deleteStatus = 1")
	//endregion

	//region order by
	sqlString.Append(" order by setmealid desc")
	//endregion

	//region limit
	sqlString.Append(" limit ?,?")
	params = append(params, (pageIndex-1)*pageSize)
	params = append(params, pageSize)
	//endregion

	//region Query
	rows, err := garglewool.Queryx(sqlString.ToString(), params...)
	if err != nil {
		return setmeals, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
	//endregion
}

// 解析【套餐表】表记录
func (d *setmeal) _RowsToArray(rows *sqlx.Rows) (setmeals []model.Setmeal, err error) {
	for rows.Next() {
		m := model.Setmeal{}
		err = rows.Scan(&m.Setmealid, &m.Shopid, &m.Pkgcode, &m.Setmealtype, &m.Title, &m.People, &m.Isorder, &m.Isrefund, &m.Isinhouse, &m.Isnew, &m.Isrecommend, &m.Validdays, &m.Watchers, &m.Createtime, &m.Validtime, &m.Exittime, &m.Links, &m.Info, &m.Adder, &m.Addtime, &m.Moder, &m.Modtime, &m.DeleteStatus)
		if err != nil {
			return setmeals, err
		}
		setmeals = append(setmeals, m)
	}
	return setmeals, err
}
