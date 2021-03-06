package dao

import "github.com/tianxiaoyang2018/go-homework/pgutil"

var hashMap map[string]map[int]string = make(map[string]map[int]string)

func initDictionaryMap() {
	con := pgutil.GetDB()
	rows, err := con.Query("SELECT colname,value,remark FROM t_dictionary")
	pgutil.CheckErr(err)

	for rows.Next() {
		var colname string
		var value int
		var remark string
		err = rows.Scan(&colname, &value, &remark)
		pgutil.CheckErr(err)

		innerHashMap := hashMap[colname]
		if innerHashMap == nil {
			innerHashMap = make(map[int]string)
			hashMap[colname] = innerHashMap
		}
		innerHashMap[value] = remark

	}

}

func GetDictiondayByColumn(column string) map[int]string {
	if hashMap == nil || len(hashMap) == 0 {
		initDictionaryMap()
	}
	return hashMap[column]
}
