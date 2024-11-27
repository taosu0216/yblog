// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ArticlesColumns holds the columns for the "articles" table.
	ArticlesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "title", Type: field.TypeString, Unique: true},
		{Name: "desc", Type: field.TypeString},
		{Name: "category", Type: field.TypeString},
		{Name: "tags", Type: field.TypeString},
		{Name: "url", Type: field.TypeString},
		{Name: "create_time", Type: field.TypeString, Default: "2024-11-27 11:55:19"},
		{Name: "is_show", Type: field.TypeBool, Default: true},
		{Name: "content", Type: field.TypeString, Size: 2147483647, Default: ""},
	}
	// ArticlesTable holds the schema information for the "articles" table.
	ArticlesTable = &schema.Table{
		Name:       "articles",
		Columns:    ArticlesColumns,
		PrimaryKey: []*schema.Column{ArticlesColumns[0]},
	}
	// FriendsColumns holds the columns for the "friends" table.
	FriendsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "title", Type: field.TypeString, Default: ""},
		{Name: "desc", Type: field.TypeString, Default: ""},
		{Name: "link", Type: field.TypeString, Default: ""},
		{Name: "avatar", Type: field.TypeString, Default: ""},
		{Name: "create_time", Type: field.TypeTime},
	}
	// FriendsTable holds the schema information for the "friends" table.
	FriendsTable = &schema.Table{
		Name:       "friends",
		Columns:    FriendsColumns,
		PrimaryKey: []*schema.Column{FriendsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "is_root", Type: field.TypeBool},
		{Name: "create_time", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ArticlesTable,
		FriendsTable,
		UsersTable,
	}
)

func init() {
}
