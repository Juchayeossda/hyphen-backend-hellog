// Code generated by ent, DO NOT EDIT.

package user

import (
	"hyphen-backend-hellog/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// IDFromUserMs applies equality check predicate on the "id_from_user_ms" field. It's identical to IDFromUserMsEQ.
func IDFromUserMs(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldIDFromUserMs, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldName, v))
}

// ProfileImage applies equality check predicate on the "profile_image" field. It's identical to ProfileImageEQ.
func ProfileImage(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldProfileImage, v))
}

// JoinedAt applies equality check predicate on the "joined_at" field. It's identical to JoinedAtEQ.
func JoinedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldJoinedAt, v))
}

// IDFromUserMsEQ applies the EQ predicate on the "id_from_user_ms" field.
func IDFromUserMsEQ(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldIDFromUserMs, v))
}

// IDFromUserMsNEQ applies the NEQ predicate on the "id_from_user_ms" field.
func IDFromUserMsNEQ(v int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldIDFromUserMs, v))
}

// IDFromUserMsIn applies the In predicate on the "id_from_user_ms" field.
func IDFromUserMsIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldIDFromUserMs, vs...))
}

// IDFromUserMsNotIn applies the NotIn predicate on the "id_from_user_ms" field.
func IDFromUserMsNotIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldIDFromUserMs, vs...))
}

// IDFromUserMsGT applies the GT predicate on the "id_from_user_ms" field.
func IDFromUserMsGT(v int) predicate.User {
	return predicate.User(sql.FieldGT(FieldIDFromUserMs, v))
}

// IDFromUserMsGTE applies the GTE predicate on the "id_from_user_ms" field.
func IDFromUserMsGTE(v int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldIDFromUserMs, v))
}

// IDFromUserMsLT applies the LT predicate on the "id_from_user_ms" field.
func IDFromUserMsLT(v int) predicate.User {
	return predicate.User(sql.FieldLT(FieldIDFromUserMs, v))
}

// IDFromUserMsLTE applies the LTE predicate on the "id_from_user_ms" field.
func IDFromUserMsLTE(v int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldIDFromUserMs, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldName, v))
}

// ProfileImageEQ applies the EQ predicate on the "profile_image" field.
func ProfileImageEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldProfileImage, v))
}

// ProfileImageNEQ applies the NEQ predicate on the "profile_image" field.
func ProfileImageNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldProfileImage, v))
}

// ProfileImageIn applies the In predicate on the "profile_image" field.
func ProfileImageIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldProfileImage, vs...))
}

// ProfileImageNotIn applies the NotIn predicate on the "profile_image" field.
func ProfileImageNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldProfileImage, vs...))
}

// ProfileImageGT applies the GT predicate on the "profile_image" field.
func ProfileImageGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldProfileImage, v))
}

// ProfileImageGTE applies the GTE predicate on the "profile_image" field.
func ProfileImageGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldProfileImage, v))
}

// ProfileImageLT applies the LT predicate on the "profile_image" field.
func ProfileImageLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldProfileImage, v))
}

// ProfileImageLTE applies the LTE predicate on the "profile_image" field.
func ProfileImageLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldProfileImage, v))
}

// ProfileImageContains applies the Contains predicate on the "profile_image" field.
func ProfileImageContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldProfileImage, v))
}

// ProfileImageHasPrefix applies the HasPrefix predicate on the "profile_image" field.
func ProfileImageHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldProfileImage, v))
}

// ProfileImageHasSuffix applies the HasSuffix predicate on the "profile_image" field.
func ProfileImageHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldProfileImage, v))
}

// ProfileImageEqualFold applies the EqualFold predicate on the "profile_image" field.
func ProfileImageEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldProfileImage, v))
}

// ProfileImageContainsFold applies the ContainsFold predicate on the "profile_image" field.
func ProfileImageContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldProfileImage, v))
}

// JoinedAtEQ applies the EQ predicate on the "joined_at" field.
func JoinedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldJoinedAt, v))
}

// JoinedAtNEQ applies the NEQ predicate on the "joined_at" field.
func JoinedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldJoinedAt, v))
}

// JoinedAtIn applies the In predicate on the "joined_at" field.
func JoinedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldJoinedAt, vs...))
}

// JoinedAtNotIn applies the NotIn predicate on the "joined_at" field.
func JoinedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldJoinedAt, vs...))
}

// JoinedAtGT applies the GT predicate on the "joined_at" field.
func JoinedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldJoinedAt, v))
}

// JoinedAtGTE applies the GTE predicate on the "joined_at" field.
func JoinedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldJoinedAt, v))
}

// JoinedAtLT applies the LT predicate on the "joined_at" field.
func JoinedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldJoinedAt, v))
}

// JoinedAtLTE applies the LTE predicate on the "joined_at" field.
func JoinedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldJoinedAt, v))
}

// HasPosts applies the HasEdge predicate on the "posts" edge.
func HasPosts() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PostsTable, PostsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPostsWith applies the HasEdge predicate on the "posts" edge with a given conditions (other predicates).
func HasPostsWith(preds ...predicate.Post) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newPostsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasComments applies the HasEdge predicate on the "comments" edge.
func HasComments() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CommentsTable, CommentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCommentsWith applies the HasEdge predicate on the "comments" edge with a given conditions (other predicates).
func HasCommentsWith(preds ...predicate.Comment) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newCommentsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}
