package dao

import "pgutil"
import "bean"
import "fmt"

/**
 * get relationships by userId
 */
func ListRelationship(userId int) []bean.RelationshipCoreInfo {
	con := pgutil.PgConnect()

	stmt, err := con.Prepare("select user_id2, state, type from t_relationship where user_id1 = $1")
	pgutil.CheckErr(err)
	rows, err := stmt.Query(userId)
	pgutil.CheckErr(err)

	var relationships []bean.RelationshipCoreInfo
	for rows.Next() {
		var userId int
		var state string
		var relationshipType string
		err = rows.Scan(&userId, &state, &relationshipType)
		pgutil.CheckErr(err)
		fmt.Println("userId = ", userId, ",state = ", state, ",relationshipType = ", relationshipType)
		relationships = append(relationships, bean.RelationshipCoreInfo{UserId: userId, State: state, RelationshipType: relationshipType})
	}

	pgutil.PgClose(con)

	return relationships
}

func GetRelationshipByUserIdAndOtherUserId(userId int, otherUserId int) *bean.RelationshipCoreInfo {
	con := pgutil.PgConnect()
	stmt, err := con.Prepare("select user_id2, state, type from t_relationship where user_id1 = $1 and user_id2 = $2")
	pgutil.CheckErr(err)
	rows, err := stmt.Query(userId, otherUserId)
	pgutil.CheckErr(err)

	var relationship *bean.RelationshipCoreInfo
	if rows.Next() {
		var userId int
		var state string
		var relationshipType string
		err = rows.Scan(&userId, &state, &relationshipType)
		pgutil.CheckErr(err)
		relationship = &bean.RelationshipCoreInfo{UserId: userId, State: state, RelationshipType: relationshipType}

	}

	pgutil.PgClose(con)

	return relationship
}

// 按照探探的产品设计，候选列表里，如果把此人左划置为不喜欢，以后就不会再出现了，右划置为喜欢也是同理，所以不存在用户1反复设置对用户2的状态，所以这一步只有添加，没有修改
// 可能出现并发的情况有2种
// ----1、用户1登陆了多个终端，并且多个终端同时操作。可能恰好多个终端都看到了用户2，然后划动，此时要在数据库里创建一条 user_id1 = 用户1 ,user_id2 = 用户2 的数据。
// ----由于数据库设置了 (user_id1,user_id2)的唯一约束，只有一个线程可以创建成功，其余均会抛异常
// ----2、用户1喜欢用户2的同时，可能用户2也喜欢了用户1，此时生成2条 liked 记录，就要求2条记录都变为 matched
// ----通过sql维持操作原子性，满足2个exists条件（2个用户彼此之间都是喜欢关系）修改为matched
func UpdateRelationship(userId int, otherUserId int, state string) {

	relationship1 := GetRelationshipByUserIdAndOtherUserId(userId, otherUserId)

	fmt.Println("relationship1=", relationship1)
	con := pgutil.PgConnect()

	if relationship1 == nil {
		stmt, err := con.Prepare("insert into t_relationship(user_id1, user_id2, state, type) values($1,$2,$3,'relationship')")
		pgutil.CheckErr(err)
		res, err := stmt.Exec(userId, otherUserId, state)
		pgutil.CheckErr(err)
		fmt.Println(res.RowsAffected())
	}

	if state == "liked" {
		stmt, err := con.Prepare("update t_relationship set state = 'mathced' where user_id1 in ($1, $2) and user_id2 in ($1, $2)  and exists(select 1 from t_relationship where user_id1=$1 and user_id2=$2 and state='liked') and exists(select 1 from t_relationship where user_id2=$1 and user_id1=$2 and state='liked')")
		pgutil.CheckErr(err)
		stmt.Exec(userId, otherUserId)
	}
	pgutil.PgClose(con)
}
