// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTUser = "t_user"

// TUser mapped from table <t_user>
type TUser struct {
	ID        int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string `gorm:"column:name" json:"name"`
	Password  string `gorm:"column:password" json:"password"`
	Telephone string `gorm:"column:telephone" json:"telephone"`
}

// TableName TUser's table name
func (*TUser) TableName() string {
	return TableNameTUser
}
