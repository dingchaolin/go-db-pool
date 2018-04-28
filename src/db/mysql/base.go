package mysql

import (
	"errors"
	"database/sql"
)

func BatchInsert(tableName string, rowList []map[string]interface{}) (int64, error) {
	if len(rowList) == 0 {
		return 0, nil
	}
	sql := "insert into `" + tableName + "` ( "
	var argArray []interface{}
	fieldStr := ""
	valueStr := ""
	indexNum := 0

	var fieldList []string //先查出所有字段
	for key, _ := range rowList[0] {
		fieldList = append(fieldList, key)
		if indexNum == 0 {
			fieldStr = fieldStr + "`" + key + "`"
		} else {
			fieldStr = fieldStr + ",`" + key + "`"
		}
		indexNum++
	}

	for rowIndex, row := range rowList {

		if rowIndex > 0 { //只有一个row的时候，插入一个没有逗号
			valueStr += ",("
		} else {
			valueStr += "("
		}

		for filedIndex, key := range fieldList {
			if filedIndex == 0 {
				valueStr += "?"
			} else {
				valueStr += ",?"
			}
			if _, ok := row[key]; ok {
				argArray = append(argArray, row[key])
			} else {
				return 0, errors.New("BatchInsert err " + key + " not exist at index " + string(rowIndex) + " ")
			}
		}
		valueStr += ")"

	}
	sql = sql + fieldStr + ") values " + valueStr + ""
	//fmt.Println("sql==",sql,argArray)
	result, err := MysqlConn.Exec(sql, argArray...)
	if err != nil {
		return 0, err
	}
	rowEffect, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowEffect, nil
}

func Insert(tableName string, rowMap map[string]interface{}) (int64, error) {
	sql := "replace into `" + tableName + "` ( "
	var argArray []interface{}
	fieldStr := ""
	valueStr := ""
	indexNum := 0
	/*
	key-value 必须在一次循环中取出
	 */
	for key, value := range rowMap {
		if indexNum == 0 {
			fieldStr = fieldStr + "`" + key + "`"
			valueStr = valueStr + "?"
		} else {
			fieldStr = fieldStr + ",`" + key + "`"
			valueStr = valueStr + ",?"
		}
		argArray = append(argArray, value)
		indexNum++
	}
	sql = sql + fieldStr + ") values (" + valueStr + ")"
	result, err := MysqlConn.Exec(sql, argArray...)
	if err != nil {
		return 0, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastInsertId, nil
}

func FetchList(tableName string, fields string, whereMap map[string]interface{}) (*sql.Rows, error) {
	indexNum := 0
	var argArray []interface{}
	whereStr := ""
	for key, value := range whereMap {
		if indexNum == 0 {
			whereStr = whereStr + " where `" + key + "` = ? "
		} else {
			whereStr = whereStr + "And `" + key + "` = ? "
		}
		argArray = append(argArray, value)
		indexNum++
	}
	sql := "select " + fields + " from " + tableName + whereStr
	rows, err := MysqlConn.Query(sql, argArray...)
	if err != nil {
		return nil, err
	}
	return rows, err
}

func Update(tableName string, fieldsMap map[string]interface{}, whereMap map[string]interface{}) (int64, error) {

	indexNum := 0
	var argArray []interface{}
	setStr := ""

	for key, value := range fieldsMap {
		if indexNum == 0 {
			setStr = setStr + " `" + key + "` = ? "
		} else {
			setStr = setStr + ", `" + key + "` = ? "
		}
		argArray = append(argArray, value)
		indexNum++
	}
	indexNum = 0
	whereStr := ""
	for key, value := range whereMap {
		if indexNum == 0 {
			whereStr = whereStr + " where `" + key + "` = ? "
		} else {
			whereStr = whereStr + "And `" + key + "` = ? "
		}
		argArray = append(argArray, value)
		indexNum++
	}

	sql := "update " + tableName + " set " + setStr + whereStr
	_, err := MysqlConn.Exec(sql, argArray...)
	if err != nil {
		return 0, err
	}
	return 1, nil
}

func Delete(tableName string, whereMap map[string]interface{}) (bool, error) {

	indexNum := 0
	var argArray []interface{}
	whereStr := ""
	for key, value := range whereMap {
		if indexNum == 0 {
			whereStr = whereStr + " where `" + key + "` = ? "
		} else {
			whereStr = whereStr + "And `" + key + "` = ? "
		}
		argArray = append(argArray, value)
		indexNum++
	}

	sql := "delete from " + tableName + whereStr
	_, err := MysqlConn.Exec(sql, argArray...)
	if err != nil {
		return false, err
	}
	return true, nil
}

func Count(tableName string, whereMap map[string]interface{}) (int64, error) {
	rows, err := FetchList(tableName, "count(1) as num", whereMap)
	if err != nil {
		return 0, err
	}
	num := int64(0)
	for rows.Next() {
		rows.Scan(&num)
		break;
	}
	return num, nil
}
