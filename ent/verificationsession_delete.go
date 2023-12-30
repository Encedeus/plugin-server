// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Encedeus/pluginServer/ent/predicate"
	"github.com/Encedeus/pluginServer/ent/verificationsession"
)

// VerificationSessionDelete is the builder for deleting a VerificationSession entity.
type VerificationSessionDelete struct {
	config
	hooks    []Hook
	mutation *VerificationSessionMutation
}

// Where appends a list predicates to the VerificationSessionDelete builder.
func (vsd *VerificationSessionDelete) Where(ps ...predicate.VerificationSession) *VerificationSessionDelete {
	vsd.mutation.Where(ps...)
	return vsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (vsd *VerificationSessionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, vsd.sqlExec, vsd.mutation, vsd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (vsd *VerificationSessionDelete) ExecX(ctx context.Context) int {
	n, err := vsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (vsd *VerificationSessionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(verificationsession.Table, sqlgraph.NewFieldSpec(verificationsession.FieldID, field.TypeString))
	if ps := vsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, vsd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	vsd.mutation.done = true
	return affected, err
}

// VerificationSessionDeleteOne is the builder for deleting a single VerificationSession entity.
type VerificationSessionDeleteOne struct {
	vsd *VerificationSessionDelete
}

// Where appends a list predicates to the VerificationSessionDelete builder.
func (vsdo *VerificationSessionDeleteOne) Where(ps ...predicate.VerificationSession) *VerificationSessionDeleteOne {
	vsdo.vsd.mutation.Where(ps...)
	return vsdo
}

// Exec executes the deletion query.
func (vsdo *VerificationSessionDeleteOne) Exec(ctx context.Context) error {
	n, err := vsdo.vsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{verificationsession.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (vsdo *VerificationSessionDeleteOne) ExecX(ctx context.Context) {
	if err := vsdo.Exec(ctx); err != nil {
		panic(err)
	}
}