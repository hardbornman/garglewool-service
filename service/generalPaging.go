package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ha666/golibs"
	"github.com/hardbornman/garglewool-service/dao"
	"github.com/hardbornman/garglewool-service/model"
	"github.com/xwb1989/sqlparser"
	"strings"
)

// 通用分页查询
func GeneralPaging(sql string, pageIndex, pageSize int) (list []map[string]interface{}, total int, err error) {
	//region 验证请求路径
	validCallPath()
	//endregion

	//region 解析sql语句
	lowerSqlString := strings.ToLower(sql)
	stmt, err := sqlparser.Parse(sql)
	if err != nil {
		return list, total, err
	}
	var sqlStructure model.SelectSQLStruct
	err = json.Unmarshal(golibs.StringToSliceByte(golibs.ToJson(stmt)), &sqlStructure)
	if err != nil {
		return list, total, err
	}
	//endregion

	//region 验证sql语句
	{

		//region 判断是否查询语句
		if len(sqlStructure.SelectExprs) <= 0 || len(sqlStructure.From) <= 0 {
			return list, total, errors.New("非法sql，非查询语句")
		}
		//endregion

		//region 判断是否包含from
		if !strings.Contains(lowerSqlString, "from") {
			return list, total, errors.New("非法sql，没有找到from")
		}
		//endregion

		//region 判断是否包含limit
		if golibs.Length(sqlStructure.Limit.Offset.Val) > 0 || golibs.Length(sqlStructure.Limit.Rowcount.Val) > 0 {
			return list, total, errors.New("非法sql，不能包含limit")
		}
		//endregion

		//region 判断是否包含or
		if strings.Contains(lowerSqlString, " or ") {
			return list, total, errors.New("非法sql，不能包含or")
		}
		//endregion

		//region 判断是否包含pwd
		if strings.Contains(lowerSqlString, "pwd") {
			return list, total, errors.New("非法sql，不能包含pwd")
		}
		//endregion

		//region 判断是否包含group by或者having
		if strings.Contains(lowerSqlString, "group by") || strings.Contains(lowerSqlString, "having") {
			return list, total, errors.New("非法sql，不能包含分组信息")
		}
		//endregion

	}
	//endregion

	//region 查询总数
	{
		start := strings.Index(lowerSqlString, "from")
		data := golibs.NewStringBuilder()
		data.Append("select count(0) as 'Count' ")
		start1 := strings.Index(lowerSqlString, " order by")
		if start1 > -1 {
			data.Append(golibs.SubString(sql, start, -start+start1))
		} else {
			data.Append(golibs.SubString(sql, start, len(sql)-start))
		}
		total, err = dao.GeneralPagingDao.GeneralPagingTotal(data.ToString())
		if err != nil {
			return list, total, errors.New("查询总数出错:" + err.Error())
		}
		if total <= 0 {
			return list, total, nil
		}
	}
	//endregion

	//region 查询列表
	sql = fmt.Sprintf("%s limit %d,%d", sql, (pageIndex-1)*pageSize, pageSize)
	list, err = dao.GeneralPagingDao.GeneralPagingList(fmt.Sprintf(sql))
	return list, total, err
	//endregion

}
