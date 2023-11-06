// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/e2b-dev/infra/packages/api/internal/db/ent/internal"
	"github.com/e2b-dev/infra/packages/api/internal/db/ent/predicate"
	"github.com/e2b-dev/infra/packages/api/internal/db/ent/team"
	"github.com/e2b-dev/infra/packages/api/internal/db/ent/tier"
	"github.com/google/uuid"
)

// TierUpdate is the builder for updating Tier entities.
type TierUpdate struct {
	config
	hooks    []Hook
	mutation *TierMutation
}

// Where appends a list predicates to the TierUpdate builder.
func (tu *TierUpdate) Where(ps ...predicate.Tier) *TierUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetVcpu sets the "vcpu" field.
func (tu *TierUpdate) SetVcpu(i int64) *TierUpdate {
	tu.mutation.ResetVcpu()
	tu.mutation.SetVcpu(i)
	return tu
}

// AddVcpu adds i to the "vcpu" field.
func (tu *TierUpdate) AddVcpu(i int64) *TierUpdate {
	tu.mutation.AddVcpu(i)
	return tu
}

// SetRAMMB sets the "ram_mb" field.
func (tu *TierUpdate) SetRAMMB(i int64) *TierUpdate {
	tu.mutation.ResetRAMMB()
	tu.mutation.SetRAMMB(i)
	return tu
}

// AddRAMMB adds i to the "ram_mb" field.
func (tu *TierUpdate) AddRAMMB(i int64) *TierUpdate {
	tu.mutation.AddRAMMB(i)
	return tu
}

// SetDiskMB sets the "disk_mb" field.
func (tu *TierUpdate) SetDiskMB(i int64) *TierUpdate {
	tu.mutation.ResetDiskMB()
	tu.mutation.SetDiskMB(i)
	return tu
}

// AddDiskMB adds i to the "disk_mb" field.
func (tu *TierUpdate) AddDiskMB(i int64) *TierUpdate {
	tu.mutation.AddDiskMB(i)
	return tu
}

// AddTeamIDs adds the "teams" edge to the Team entity by IDs.
func (tu *TierUpdate) AddTeamIDs(ids ...uuid.UUID) *TierUpdate {
	tu.mutation.AddTeamIDs(ids...)
	return tu
}

// AddTeams adds the "teams" edges to the Team entity.
func (tu *TierUpdate) AddTeams(t ...*Team) *TierUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddTeamIDs(ids...)
}

// Mutation returns the TierMutation object of the builder.
func (tu *TierUpdate) Mutation() *TierMutation {
	return tu.mutation
}

// ClearTeams clears all "teams" edges to the Team entity.
func (tu *TierUpdate) ClearTeams() *TierUpdate {
	tu.mutation.ClearTeams()
	return tu
}

// RemoveTeamIDs removes the "teams" edge to Team entities by IDs.
func (tu *TierUpdate) RemoveTeamIDs(ids ...uuid.UUID) *TierUpdate {
	tu.mutation.RemoveTeamIDs(ids...)
	return tu
}

// RemoveTeams removes "teams" edges to Team entities.
func (tu *TierUpdate) RemoveTeams(t ...*Team) *TierUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveTeamIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TierUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TierUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TierUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TierUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TierUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(tier.Table, tier.Columns, sqlgraph.NewFieldSpec(tier.FieldID, field.TypeString))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Vcpu(); ok {
		_spec.SetField(tier.FieldVcpu, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.AddedVcpu(); ok {
		_spec.AddField(tier.FieldVcpu, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.RAMMB(); ok {
		_spec.SetField(tier.FieldRAMMB, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.AddedRAMMB(); ok {
		_spec.AddField(tier.FieldRAMMB, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.DiskMB(); ok {
		_spec.SetField(tier.FieldDiskMB, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.AddedDiskMB(); ok {
		_spec.AddField(tier.FieldDiskMB, field.TypeInt64, value)
	}
	if tu.mutation.TeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tier.TeamsTable,
			Columns: []string{tier.TeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = tu.schemaConfig.Team
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedTeamsIDs(); len(nodes) > 0 && !tu.mutation.TeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tier.TeamsTable,
			Columns: []string{tier.TeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = tu.schemaConfig.Team
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.TeamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tier.TeamsTable,
			Columns: []string{tier.TeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = tu.schemaConfig.Team
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = tu.schemaConfig.Tier
	ctx = internal.NewSchemaConfigContext(ctx, tu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tier.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TierUpdateOne is the builder for updating a single Tier entity.
type TierUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TierMutation
}

// SetVcpu sets the "vcpu" field.
func (tuo *TierUpdateOne) SetVcpu(i int64) *TierUpdateOne {
	tuo.mutation.ResetVcpu()
	tuo.mutation.SetVcpu(i)
	return tuo
}

// AddVcpu adds i to the "vcpu" field.
func (tuo *TierUpdateOne) AddVcpu(i int64) *TierUpdateOne {
	tuo.mutation.AddVcpu(i)
	return tuo
}

// SetRAMMB sets the "ram_mb" field.
func (tuo *TierUpdateOne) SetRAMMB(i int64) *TierUpdateOne {
	tuo.mutation.ResetRAMMB()
	tuo.mutation.SetRAMMB(i)
	return tuo
}

// AddRAMMB adds i to the "ram_mb" field.
func (tuo *TierUpdateOne) AddRAMMB(i int64) *TierUpdateOne {
	tuo.mutation.AddRAMMB(i)
	return tuo
}

// SetDiskMB sets the "disk_mb" field.
func (tuo *TierUpdateOne) SetDiskMB(i int64) *TierUpdateOne {
	tuo.mutation.ResetDiskMB()
	tuo.mutation.SetDiskMB(i)
	return tuo
}

// AddDiskMB adds i to the "disk_mb" field.
func (tuo *TierUpdateOne) AddDiskMB(i int64) *TierUpdateOne {
	tuo.mutation.AddDiskMB(i)
	return tuo
}

// AddTeamIDs adds the "teams" edge to the Team entity by IDs.
func (tuo *TierUpdateOne) AddTeamIDs(ids ...uuid.UUID) *TierUpdateOne {
	tuo.mutation.AddTeamIDs(ids...)
	return tuo
}

// AddTeams adds the "teams" edges to the Team entity.
func (tuo *TierUpdateOne) AddTeams(t ...*Team) *TierUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddTeamIDs(ids...)
}

// Mutation returns the TierMutation object of the builder.
func (tuo *TierUpdateOne) Mutation() *TierMutation {
	return tuo.mutation
}

// ClearTeams clears all "teams" edges to the Team entity.
func (tuo *TierUpdateOne) ClearTeams() *TierUpdateOne {
	tuo.mutation.ClearTeams()
	return tuo
}

// RemoveTeamIDs removes the "teams" edge to Team entities by IDs.
func (tuo *TierUpdateOne) RemoveTeamIDs(ids ...uuid.UUID) *TierUpdateOne {
	tuo.mutation.RemoveTeamIDs(ids...)
	return tuo
}

// RemoveTeams removes "teams" edges to Team entities.
func (tuo *TierUpdateOne) RemoveTeams(t ...*Team) *TierUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveTeamIDs(ids...)
}

// Where appends a list predicates to the TierUpdate builder.
func (tuo *TierUpdateOne) Where(ps ...predicate.Tier) *TierUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TierUpdateOne) Select(field string, fields ...string) *TierUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Tier entity.
func (tuo *TierUpdateOne) Save(ctx context.Context) (*Tier, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TierUpdateOne) SaveX(ctx context.Context) *Tier {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TierUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TierUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TierUpdateOne) sqlSave(ctx context.Context) (_node *Tier, err error) {
	_spec := sqlgraph.NewUpdateSpec(tier.Table, tier.Columns, sqlgraph.NewFieldSpec(tier.FieldID, field.TypeString))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Tier.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tier.FieldID)
		for _, f := range fields {
			if !tier.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tier.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Vcpu(); ok {
		_spec.SetField(tier.FieldVcpu, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.AddedVcpu(); ok {
		_spec.AddField(tier.FieldVcpu, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.RAMMB(); ok {
		_spec.SetField(tier.FieldRAMMB, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.AddedRAMMB(); ok {
		_spec.AddField(tier.FieldRAMMB, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.DiskMB(); ok {
		_spec.SetField(tier.FieldDiskMB, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.AddedDiskMB(); ok {
		_spec.AddField(tier.FieldDiskMB, field.TypeInt64, value)
	}
	if tuo.mutation.TeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tier.TeamsTable,
			Columns: []string{tier.TeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = tuo.schemaConfig.Team
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedTeamsIDs(); len(nodes) > 0 && !tuo.mutation.TeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tier.TeamsTable,
			Columns: []string{tier.TeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = tuo.schemaConfig.Team
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.TeamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tier.TeamsTable,
			Columns: []string{tier.TeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = tuo.schemaConfig.Team
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = tuo.schemaConfig.Tier
	ctx = internal.NewSchemaConfigContext(ctx, tuo.schemaConfig)
	_node = &Tier{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tier.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
