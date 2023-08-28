package entities

import "database/sql"

type BannedData struct {
	ID                int64        `db:"id"`
	Name              string       `db:"user_name"`
	Msisdn            string       `db:"user_msisdn"`
	BannedStatusID    int64        `db:"banned_status_id"`
	BannedStatus      string       `db:"banned_status"`
	BannedStatusLabel string       `db:"banned_status_label"`
	Counter           int          `db:"counter"`
	ValidUntil        sql.NullTime `db:"valid_until"`
}
