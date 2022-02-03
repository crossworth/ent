// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BlobsColumns holds the columns for the "blobs" table.
	BlobsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true, Default: "uuid_generate_v4()"},
		{Name: "uuid", Type: field.TypeUUID, Unique: true},
		{Name: "count", Type: field.TypeInt, Default: 0},
		{Name: "blob_parent", Type: field.TypeUUID, Unique: true, Nullable: true},
	}
	// BlobsTable holds the schema information for the "blobs" table.
	BlobsTable = &schema.Table{
		Name:       "blobs",
		Columns:    BlobsColumns,
		PrimaryKey: []*schema.Column{BlobsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "blobs_blobs_parent",
				Columns:    []*schema.Column{BlobsColumns[3]},
				RefColumns: []*schema.Column{BlobsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CarsColumns holds the columns for the "cars" table.
	CarsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "before_id", Type: field.TypeFloat64, Nullable: true},
		{Name: "after_id", Type: field.TypeFloat64, Nullable: true},
		{Name: "model", Type: field.TypeString},
		{Name: "pet_cars", Type: field.TypeString, Nullable: true, Size: 36},
	}
	// CarsTable holds the schema information for the "cars" table.
	CarsTable = &schema.Table{
		Name:       "cars",
		Columns:    CarsColumns,
		PrimaryKey: []*schema.Column{CarsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "cars_pets_cars",
				Columns:    []*schema.Column{CarsColumns[4]},
				RefColumns: []*schema.Column{PetsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DevicesColumns holds the columns for the "devices" table.
	DevicesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeBytes, Size: 64},
		{Name: "device_active_session", Type: field.TypeBytes, Nullable: true, Size: 64},
	}
	// DevicesTable holds the schema information for the "devices" table.
	DevicesTable = &schema.Table{
		Name:       "devices",
		Columns:    DevicesColumns,
		PrimaryKey: []*schema.Column{DevicesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "devices_sessions_active_session",
				Columns:    []*schema.Column{DevicesColumns[1]},
				RefColumns: []*schema.Column{SessionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DocsColumns holds the columns for the "docs" table.
	DocsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 36},
		{Name: "text", Type: field.TypeString, Nullable: true},
		{Name: "doc_children", Type: field.TypeString, Nullable: true, Size: 36},
	}
	// DocsTable holds the schema information for the "docs" table.
	DocsTable = &schema.Table{
		Name:       "docs",
		Columns:    DocsColumns,
		PrimaryKey: []*schema.Column{DocsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "docs_docs_children",
				Columns:    []*schema.Column{DocsColumns[2]},
				RefColumns: []*schema.Column{DocsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
	}
	// MixinIdsColumns holds the columns for the "mixin_ids" table.
	MixinIdsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "some_field", Type: field.TypeString},
		{Name: "mixin_field", Type: field.TypeString},
	}
	// MixinIdsTable holds the schema information for the "mixin_ids" table.
	MixinIdsTable = &schema.Table{
		Name:       "mixin_ids",
		Columns:    MixinIdsColumns,
		PrimaryKey: []*schema.Column{MixinIdsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "mixinid_id",
				Unique:  false,
				Columns: []*schema.Column{MixinIdsColumns[0]},
			},
			{
				Name:    "mixinid_id_some_field",
				Unique:  false,
				Columns: []*schema.Column{MixinIdsColumns[0], MixinIdsColumns[1]},
			},
			{
				Name:    "mixinid_id_mixin_field",
				Unique:  false,
				Columns: []*schema.Column{MixinIdsColumns[0], MixinIdsColumns[2]},
			},
			{
				Name:    "mixinid_id_mixin_field_some_field",
				Unique:  false,
				Columns: []*schema.Column{MixinIdsColumns[0], MixinIdsColumns[2], MixinIdsColumns[1]},
			},
		},
	}
	// NotesColumns holds the columns for the "notes" table.
	NotesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 36},
		{Name: "text", Type: field.TypeString, Nullable: true},
		{Name: "note_children", Type: field.TypeString, Nullable: true, Size: 36},
	}
	// NotesTable holds the schema information for the "notes" table.
	NotesTable = &schema.Table{
		Name:       "notes",
		Columns:    NotesColumns,
		PrimaryKey: []*schema.Column{NotesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "notes_notes_children",
				Columns:    []*schema.Column{NotesColumns[2]},
				RefColumns: []*schema.Column{NotesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OthersColumns holds the columns for the "others" table.
	OthersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeOther, Unique: true, SchemaType: map[string]string{"mysql": "bigint", "postgres": "bigint", "sqlite3": "integer"}},
	}
	// OthersTable holds the schema information for the "others" table.
	OthersTable = &schema.Table{
		Name:       "others",
		Columns:    OthersColumns,
		PrimaryKey: []*schema.Column{OthersColumns[0]},
	}
	// PetsColumns holds the columns for the "pets" table.
	PetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 36},
		{Name: "pet_best_friend", Type: field.TypeString, Unique: true, Nullable: true, Size: 36},
		{Name: "user_pets", Type: field.TypeInt, Nullable: true},
	}
	// PetsTable holds the schema information for the "pets" table.
	PetsTable = &schema.Table{
		Name:       "pets",
		Columns:    PetsColumns,
		PrimaryKey: []*schema.Column{PetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "pets_pets_best_friend",
				Columns:    []*schema.Column{PetsColumns[1]},
				RefColumns: []*schema.Column{PetsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "pets_users_pets",
				Columns:    []*schema.Column{PetsColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SessionsColumns holds the columns for the "sessions" table.
	SessionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeBytes, Size: 64},
		{Name: "device_sessions", Type: field.TypeBytes, Nullable: true, Size: 64},
	}
	// SessionsTable holds the schema information for the "sessions" table.
	SessionsTable = &schema.Table{
		Name:       "sessions",
		Columns:    SessionsColumns,
		PrimaryKey: []*schema.Column{SessionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sessions_devices_sessions",
				Columns:    []*schema.Column{SessionsColumns[1]},
				RefColumns: []*schema.Column{DevicesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeInt, Increment: true},
		{Name: "user_children", Type: field.TypeInt, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_users_children",
				Columns:    []*schema.Column{UsersColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BlobLinksColumns holds the columns for the "blob_links" table.
	BlobLinksColumns = []*schema.Column{
		{Name: "blob_id", Type: field.TypeUUID},
		{Name: "link_id", Type: field.TypeUUID},
	}
	// BlobLinksTable holds the schema information for the "blob_links" table.
	BlobLinksTable = &schema.Table{
		Name:       "blob_links",
		Columns:    BlobLinksColumns,
		PrimaryKey: []*schema.Column{BlobLinksColumns[0], BlobLinksColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "blob_links_blob_id",
				Columns:    []*schema.Column{BlobLinksColumns[0]},
				RefColumns: []*schema.Column{BlobsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "blob_links_link_id",
				Columns:    []*schema.Column{BlobLinksColumns[1]},
				RefColumns: []*schema.Column{BlobsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// GroupUsersColumns holds the columns for the "group_users" table.
	GroupUsersColumns = []*schema.Column{
		{Name: "group_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// GroupUsersTable holds the schema information for the "group_users" table.
	GroupUsersTable = &schema.Table{
		Name:       "group_users",
		Columns:    GroupUsersColumns,
		PrimaryKey: []*schema.Column{GroupUsersColumns[0], GroupUsersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "group_users_group_id",
				Columns:    []*schema.Column{GroupUsersColumns[0]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "group_users_user_id",
				Columns:    []*schema.Column{GroupUsersColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// PetFriendsColumns holds the columns for the "pet_friends" table.
	PetFriendsColumns = []*schema.Column{
		{Name: "pet_id", Type: field.TypeString, Size: 36},
		{Name: "friend_id", Type: field.TypeString, Size: 36},
	}
	// PetFriendsTable holds the schema information for the "pet_friends" table.
	PetFriendsTable = &schema.Table{
		Name:       "pet_friends",
		Columns:    PetFriendsColumns,
		PrimaryKey: []*schema.Column{PetFriendsColumns[0], PetFriendsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "pet_friends_pet_id",
				Columns:    []*schema.Column{PetFriendsColumns[0]},
				RefColumns: []*schema.Column{PetsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "pet_friends_friend_id",
				Columns:    []*schema.Column{PetFriendsColumns[1]},
				RefColumns: []*schema.Column{PetsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BlobsTable,
		CarsTable,
		DevicesTable,
		DocsTable,
		GroupsTable,
		MixinIdsTable,
		NotesTable,
		OthersTable,
		PetsTable,
		SessionsTable,
		UsersTable,
		BlobLinksTable,
		GroupUsersTable,
		PetFriendsTable,
	}
)

func init() {
	BlobsTable.ForeignKeys[0].RefTable = BlobsTable
	CarsTable.ForeignKeys[0].RefTable = PetsTable
	DevicesTable.ForeignKeys[0].RefTable = SessionsTable
	DocsTable.ForeignKeys[0].RefTable = DocsTable
	NotesTable.ForeignKeys[0].RefTable = NotesTable
	PetsTable.ForeignKeys[0].RefTable = PetsTable
	PetsTable.ForeignKeys[1].RefTable = UsersTable
	SessionsTable.ForeignKeys[0].RefTable = DevicesTable
	UsersTable.ForeignKeys[0].RefTable = UsersTable
	BlobLinksTable.ForeignKeys[0].RefTable = BlobsTable
	BlobLinksTable.ForeignKeys[1].RefTable = BlobsTable
	GroupUsersTable.ForeignKeys[0].RefTable = GroupsTable
	GroupUsersTable.ForeignKeys[1].RefTable = UsersTable
	PetFriendsTable.ForeignKeys[0].RefTable = PetsTable
	PetFriendsTable.ForeignKeys[1].RefTable = PetsTable
}
