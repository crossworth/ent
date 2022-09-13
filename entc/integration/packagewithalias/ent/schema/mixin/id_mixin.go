package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/ent/ent/entc/integration/pacakgewithalias/pkg/v2"
)

type IDMixin struct {
	mixin.Schema
}

func (IDMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").GoType(pkg.SomeInt(0)),
	}
}
