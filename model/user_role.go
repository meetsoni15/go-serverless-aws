package model

import "time"

type UserRole struct {
	tableName struct{} `pg:"fx_user_role,alias:FUR" json:"-"`

	ID        int       `pg:"id,pk" json:"id"`
	Name      string    `pg:"role_name,type:varchar(255)" json:"role_name"`
	CreatedAt time.Time `pg:"created_at,nopk,type:timestamp" json:"created_at"`
	UpdatedAt time.Time `pg:"updated_at,nopk,type:timestamp" json:"updated_at"`
}
