package dao

import "bean"
import "pgutil"
import "fmt"

func ListUser() []bean.UserCoreInfo {

	con := pgutil.PgConnect()

	rows, err := con.Query("SELECT id,name,type FROM t_user")
	pgutil.CheckErr(err)

	var users []bean.UserCoreInfo

	for rows.Next() {
		var id int
		var name string
		var userType string
		err = rows.Scan(&id, &name, &userType)
		pgutil.CheckErr(err)
		users = append(users, bean.UserCoreInfo{Id: id, Name: name, UserType: userType})
	}

	pgutil.PgClose(con)
	return users
}

func GetTUser(id int) bean.UserCoreInfo {
	con := pgutil.PgConnect()

	stmt, err := con.Prepare("select id, name, type from t_user where id = $1")
	rows, err := stmt.Query(id)

	if rows.Next() {
		var id int
		var name string
		var userType string
		err = rows.Scan(&id, &name, &userType)
		pgutil.CheckErr(err)
		fmt.Println("id = ", id, ",name = ", name, ",type = ", userType)
		var user bean.UserCoreInfo = bean.UserCoreInfo{Id: id, Name: name, UserType: userType}
		return user
	}
	return bean.UserCoreInfo{}
}

func GetTUserByName(name string) bean.UserCoreInfo {
	con := pgutil.PgConnect()

	stmt, err := con.Prepare("select id, name, type from t_user where name = $1")
	rows, err := stmt.Query(name)

	for rows.Next() {
		var id int
		var name string
		var userType string
		err = rows.Scan(&id, &name, &userType)
		pgutil.CheckErr(err)
		var user bean.UserCoreInfo = bean.UserCoreInfo{Id: id, Name: name, UserType: userType}
		return user
	}
	return bean.UserCoreInfo{}
}

func InsertUser(user bean.UserCoreInfo) {
	con := pgutil.PgConnect()

	stmt, err := con.Prepare("INSERT INTO t_user(name, type) VALUES($1, 'user')")
	pgutil.CheckErr(err)

	stmt.Exec(user.Name)

	pgutil.CheckErr(err)

	pgutil.PgClose(con)

}