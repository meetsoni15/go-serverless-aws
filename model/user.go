package model

import "time"

type User struct {
	tableName struct{} `pg:"fx_user,alias:FUR" json:"-"`

	ID        int    `pg:"id,pk" json:"id"`
	FirstName string `pg:"first_name,type:varchar(255)" json:"first_name"`
	LastName  string `pg:"last_name,type:varchar(255)" json:"last_name"`
	Email     string `pg:"email,type:varchar(255)" json:"email"`

	UserRoleId int       `pg:"user_role_id" json:"user_role_id"`
	UserRole   *UserRole `pg:"rel:has-one,fk:user_role_id" json:"user_role,omitempty"`

	CreatedAt time.Time `pg:"created_at,nopk,type:timestamp" json:"created_at"`
	UpdatedAt time.Time `pg:"updated_at,nopk,type:timestamp" json:"updated_at"`
}
