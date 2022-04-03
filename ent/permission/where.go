// Code generated by entc, DO NOT EDIT.

package permission

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/aintsashqa/roles-and-permissions/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// CreationDate applies equality check predicate on the "creation_date" field. It's identical to CreationDateEQ.
func CreationDate(v time.Time) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreationDate), v))
	})
}

// LastUpdateDate applies equality check predicate on the "last_update_date" field. It's identical to LastUpdateDateEQ.
func LastUpdateDate(v time.Time) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastUpdateDate), v))
	})
}

// MarkAsDeleteDate applies equality check predicate on the "mark_as_delete_date" field. It's identical to MarkAsDeleteDateEQ.
func MarkAsDeleteDate(v time.Time) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMarkAsDeleteDate), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Permission {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Permission(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Permission {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Permission(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// CreationDateEQ applies the EQ predicate on the "creation_date" field.
func CreationDateEQ(v time.Time) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreationDate), v))
	})
}

// CreationDateNEQ applies the NEQ predicate on the "creation_date" field.
func CreationDateNEQ(v time.Time) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreationDate), v))
	})
}

// CreationDateIn applies the In predicate on the "creation_date" field.
func CreationDateIn(vs ...time.Time) predicate.Permission {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Permission(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreationDate), v...))
	})
}

// CreationDateNotIn applies the NotIn predicate on the "creation_date" field.
func CreationDateNotIn(vs ...time.Time) predicate.Permission {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Permission(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreationDate), v...))
	})
}

// CreationDateGT applies the GT predicate on the "creation_date" field.
func CreationDateGT(v time.Time) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreationDate), v))
	})
}

// CreationDateGTE applies the GTE predicate on the "creation_date" field.
func CreationDateGTE(v time.Time) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreationDate), v))
	})
}

// CreationDateLT applies the LT predicate on the "creation_date" field.
func CreationDateLT(v time.Time) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreationDate), v))
	})
}

// CreationDateLTE applies the LTE predicate on the "creation_date" field.
func CreationDateLTE(v time.Time) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreationDate), v))
	})
}

// LastUpdateDateEQ applies the EQ predicate on the "last_update_date" field.
func LastUpdateDateEQ(v time.Time) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastUpdateDate), v))
	})
}

// LastUpdateDateNEQ applies the NEQ predicate on the "last_update_date" field.
func LastUpdateDateNEQ(v time.Time) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastUpdateDate), v))
	})
}

// LastUpdateDateIn applies the In predicate on the "last_update_date" field.
func LastUpdateDateIn(vs ...time.Time) predicate.Permission {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Permission(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLastUpdateDate), v...))
	})
}

// LastUpdateDateNotIn applies the NotIn predicate on the "last_update_date" field.
func LastUpdateDateNotIn(vs ...time.Time) predicate.Permission {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Permission(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLastUpdateDate), v...))
	})
}

// LastUpdateDateGT applies the GT predicate on the "last_update_date" field.
func LastUpdateDateGT(v time.Time) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastUpdateDate), v))
	})
}

// LastUpdateDateGTE applies the GTE predicate on the "last_update_date" field.
func LastUpdateDateGTE(v time.Time) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastUpdateDate), v))
	})
}

// LastUpdateDateLT applies the LT predicate on the "last_update_date" field.
func LastUpdateDateLT(v time.Time) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastUpdateDate), v))
	})
}

// LastUpdateDateLTE applies the LTE predicate on the "last_update_date" field.
func LastUpdateDateLTE(v time.Time) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastUpdateDate), v))
	})
}

// MarkAsDeleteDateEQ applies the EQ predicate on the "mark_as_delete_date" field.
func MarkAsDeleteDateEQ(v time.Time) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMarkAsDeleteDate), v))
	})
}

// MarkAsDeleteDateNEQ applies the NEQ predicate on the "mark_as_delete_date" field.
func MarkAsDeleteDateNEQ(v time.Time) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMarkAsDeleteDate), v))
	})
}

// MarkAsDeleteDateIn applies the In predicate on the "mark_as_delete_date" field.
func MarkAsDeleteDateIn(vs ...time.Time) predicate.Permission {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Permission(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMarkAsDeleteDate), v...))
	})
}

// MarkAsDeleteDateNotIn applies the NotIn predicate on the "mark_as_delete_date" field.
func MarkAsDeleteDateNotIn(vs ...time.Time) predicate.Permission {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Permission(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMarkAsDeleteDate), v...))
	})
}

// MarkAsDeleteDateGT applies the GT predicate on the "mark_as_delete_date" field.
func MarkAsDeleteDateGT(v time.Time) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMarkAsDeleteDate), v))
	})
}

// MarkAsDeleteDateGTE applies the GTE predicate on the "mark_as_delete_date" field.
func MarkAsDeleteDateGTE(v time.Time) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMarkAsDeleteDate), v))
	})
}

// MarkAsDeleteDateLT applies the LT predicate on the "mark_as_delete_date" field.
func MarkAsDeleteDateLT(v time.Time) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMarkAsDeleteDate), v))
	})
}

// MarkAsDeleteDateLTE applies the LTE predicate on the "mark_as_delete_date" field.
func MarkAsDeleteDateLTE(v time.Time) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMarkAsDeleteDate), v))
	})
}

// MarkAsDeleteDateIsNil applies the IsNil predicate on the "mark_as_delete_date" field.
func MarkAsDeleteDateIsNil() predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMarkAsDeleteDate)))
	})
}

// MarkAsDeleteDateNotNil applies the NotNil predicate on the "mark_as_delete_date" field.
func MarkAsDeleteDateNotNil() predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMarkAsDeleteDate)))
	})
}

// HasRoles applies the HasEdge predicate on the "roles" edge.
func HasRoles() predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RolesTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, RolesTable, RolesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRolesWith applies the HasEdge predicate on the "roles" edge with a given conditions (other predicates).
func HasRolesWith(preds ...predicate.Role) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RolesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, RolesTable, RolesPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Permission) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Permission) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Permission) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		p(s.Not())
	})
}
