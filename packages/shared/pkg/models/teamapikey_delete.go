// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/internal"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/predicate"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/teamapikey"
)

// TeamAPIKeyDelete is the builder for deleting a TeamAPIKey entity.
type TeamAPIKeyDelete struct {
	config
	hooks    []Hook
	mutation *TeamAPIKeyMutation
}

// Where appends a list predicates to the TeamAPIKeyDelete builder.
func (takd *TeamAPIKeyDelete) Where(ps ...predicate.TeamAPIKey) *TeamAPIKeyDelete {
	takd.mutation.Where(ps...)
	return takd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (takd *TeamAPIKeyDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, takd.sqlExec, takd.mutation, takd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (takd *TeamAPIKeyDelete) ExecX(ctx context.Context) int {
	n, err := takd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (takd *TeamAPIKeyDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(teamapikey.Table, sqlgraph.NewFieldSpec(teamapikey.FieldID, field.TypeString))
	_spec.Node.Schema = takd.schemaConfig.TeamAPIKey
	ctx = internal.NewSchemaConfigContext(ctx, takd.schemaConfig)
	if ps := takd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, takd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	takd.mutation.done = true
	return affected, err
}

// TeamAPIKeyDeleteOne is the builder for deleting a single TeamAPIKey entity.
type TeamAPIKeyDeleteOne struct {
	takd *TeamAPIKeyDelete
}

// Where appends a list predicates to the TeamAPIKeyDelete builder.
func (takdo *TeamAPIKeyDeleteOne) Where(ps ...predicate.TeamAPIKey) *TeamAPIKeyDeleteOne {
	takdo.takd.mutation.Where(ps...)
	return takdo
}

// Exec executes the deletion query.
func (takdo *TeamAPIKeyDeleteOne) Exec(ctx context.Context) error {
	n, err := takdo.takd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{teamapikey.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (takdo *TeamAPIKeyDeleteOne) ExecX(ctx context.Context) {
	if err := takdo.Exec(ctx); err != nil {
		panic(err)
	}
}