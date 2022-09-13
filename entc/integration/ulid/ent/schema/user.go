package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/ent/ent/entc/integration/ulid/ent/schema/mixin"
)

type OccupancyPricing struct {
	Price int
}

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		// issue: 2761
		field.JSON("occupancy_pricing", map[string]OccupancyPricing{}),
	}
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{},
	}
}
