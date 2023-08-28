package entities

import "database/sql"

type UserData struct {
	ID      int64          `db:"id"`
	Name    string         `db:"name"`
	Msisdn  string         `db:"msisdn"`
	Address sql.NullString `db:"address"`
}
