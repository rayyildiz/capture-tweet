// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ContactUsColumns holds the columns for the "contact_us" table.
	ContactUsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 36},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "email", Type: field.TypeString},
		{Name: "full_name", Type: field.TypeString, Nullable: true, Size: 512},
		{Name: "message", Type: field.TypeString, Nullable: true, Size: 4096},
	}
	// ContactUsTable holds the schema information for the "contact_us" table.
	ContactUsTable = &schema.Table{
		Name:       "contact_us",
		Columns:    ContactUsColumns,
		PrimaryKey: []*schema.Column{ContactUsColumns[0]},
	}
	// TweetsColumns holds the columns for the "tweets" table.
	TweetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 32},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "full_text", Type: field.TypeString, Size: 1024},
		{Name: "capture_url", Type: field.TypeString, Nullable: true},
		{Name: "capture_thumb_url", Type: field.TypeString, Nullable: true},
		{Name: "lang", Type: field.TypeString, Size: 4},
		{Name: "favorite_count", Type: field.TypeInt, Default: 0},
		{Name: "retweet_count", Type: field.TypeInt, Default: 0},
		{Name: "resources", Type: field.TypeJSON},
		{Name: "posted_at", Type: field.TypeTime},
		{Name: "author_id", Type: field.TypeString, Nullable: true, Size: 32},
	}
	// TweetsTable holds the schema information for the "tweets" table.
	TweetsTable = &schema.Table{
		Name:       "tweets",
		Columns:    TweetsColumns,
		PrimaryKey: []*schema.Column{TweetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tweets_users_author",
				Columns:    []*schema.Column{TweetsColumns[11]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 32},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "screen_name", Type: field.TypeString},
		{Name: "bio", Type: field.TypeString, Nullable: true},
		{Name: "profile_image_url", Type: field.TypeString, Nullable: true},
		{Name: "registered_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ContactUsTable,
		TweetsTable,
		UsersTable,
	}
)

func init() {
	TweetsTable.ForeignKeys[0].RefTable = UsersTable
}