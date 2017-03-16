package repositories

import (
	"database/sql"
	"fmt"

	"log"

	"github.com/lib/pq"
)

// User .. repository object
type User struct {
	UserID           sql.NullInt64  `db:"user_id"`
	OrgID            sql.NullInt64  `db:"org_id"`
	OauthToken       sql.NullString `db:"oauth_token"`
	OauthTokenSecret sql.NullString `db:"oauth_token_secret"`
	Email            sql.NullString `db:"email"`
	Name             sql.NullString `db:"name"`
}

func GetUsers(field string, value string) []User {
	var retUser = []User{}
	queryField := pq.QuoteIdentifier(field)
	query := fmt.Sprintf("select * from usertable where %s = $1", queryField)
	fmt.Println(query)
	rows, err := DB.Queryx(query, value)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var u = User{}
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		retUser = append(retUser, u)
	}
	defer rows.Close()
	return retUser
}

func GetContactList(userid int) []User {
	var retUser = []User{}
	query := `select user_id, org_id, oauth_token, oauth_token_secret, email, name from usertable
						 join contactlist on
						 contactlist.targetuserid  = usertable.user_id
						 where contactlist.owneruserid = $1`
	rows, err := DB.Queryx(query, userid)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var u = User{}
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		retUser = append(retUser, u)
	}
	defer rows.Close()
	return retUser
}
