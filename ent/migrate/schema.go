// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// PluginsColumns holds the columns for the "plugins" table.
	PluginsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "owner_id", Type: field.TypeUUID},
		{Name: "source_id", Type: field.TypeInt},
	}
	// PluginsTable holds the schema information for the "plugins" table.
	PluginsTable = &schema.Table{
		Name:       "plugins",
		Columns:    PluginsColumns,
		PrimaryKey: []*schema.Column{PluginsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "plugins_users_owner",
				Columns:    []*schema.Column{PluginsColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "plugins_sources_source",
				Columns:    []*schema.Column{PluginsColumns[3]},
				RefColumns: []*schema.Column{SourcesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// SourcesColumns holds the columns for the "sources" table.
	SourcesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "repository", Type: field.TypeString},
	}
	// SourcesTable holds the schema information for the "sources" table.
	SourcesTable = &schema.Table{
		Name:       "sources",
		Columns:    SourcesColumns,
		PrimaryKey: []*schema.Column{SourcesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "auth_updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "email", Type: field.TypeString, Unique: true, Size: 32},
		{Name: "password", Type: field.TypeString},
		{Name: "name", Type: field.TypeString, Unique: true, Size: 32},
		{Name: "email_verified", Type: field.TypeBool, Default: false},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		PluginsTable,
		SourcesTable,
		UsersTable,
	}
)

func init() {
	PluginsTable.ForeignKeys[0].RefTable = UsersTable
	PluginsTable.ForeignKeys[1].RefTable = SourcesTable
}
