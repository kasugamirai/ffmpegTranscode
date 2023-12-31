// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"freefrom.space/videoTransform/ent/predicate"
	"freefrom.space/videoTransform/ent/video"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeVideo = "Video"
)

// VideoMutation represents an operation that mutates the Video nodes in the graph.
type VideoMutation struct {
	config
	op            Op
	typ           string
	id            *int
	origin_url    *string
	convert_url   *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Video, error)
	predicates    []predicate.Video
}

var _ ent.Mutation = (*VideoMutation)(nil)

// videoOption allows management of the mutation configuration using functional options.
type videoOption func(*VideoMutation)

// newVideoMutation creates new mutation for the Video entity.
func newVideoMutation(c config, op Op, opts ...videoOption) *VideoMutation {
	m := &VideoMutation{
		config:        c,
		op:            op,
		typ:           TypeVideo,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withVideoID sets the ID field of the mutation.
func withVideoID(id int) videoOption {
	return func(m *VideoMutation) {
		var (
			err   error
			once  sync.Once
			value *Video
		)
		m.oldValue = func(ctx context.Context) (*Video, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Video.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withVideo sets the old Video of the mutation.
func withVideo(node *Video) videoOption {
	return func(m *VideoMutation) {
		m.oldValue = func(context.Context) (*Video, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m VideoMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m VideoMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *VideoMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *VideoMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Video.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetOriginURL sets the "origin_url" field.
func (m *VideoMutation) SetOriginURL(s string) {
	m.origin_url = &s
}

// OriginURL returns the value of the "origin_url" field in the mutation.
func (m *VideoMutation) OriginURL() (r string, exists bool) {
	v := m.origin_url
	if v == nil {
		return
	}
	return *v, true
}

// OldOriginURL returns the old "origin_url" field's value of the Video entity.
// If the Video object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *VideoMutation) OldOriginURL(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldOriginURL is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldOriginURL requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldOriginURL: %w", err)
	}
	return oldValue.OriginURL, nil
}

// ResetOriginURL resets all changes to the "origin_url" field.
func (m *VideoMutation) ResetOriginURL() {
	m.origin_url = nil
}

// SetConvertURL sets the "convert_url" field.
func (m *VideoMutation) SetConvertURL(s string) {
	m.convert_url = &s
}

// ConvertURL returns the value of the "convert_url" field in the mutation.
func (m *VideoMutation) ConvertURL() (r string, exists bool) {
	v := m.convert_url
	if v == nil {
		return
	}
	return *v, true
}

// OldConvertURL returns the old "convert_url" field's value of the Video entity.
// If the Video object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *VideoMutation) OldConvertURL(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldConvertURL is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldConvertURL requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldConvertURL: %w", err)
	}
	return oldValue.ConvertURL, nil
}

// ClearConvertURL clears the value of the "convert_url" field.
func (m *VideoMutation) ClearConvertURL() {
	m.convert_url = nil
	m.clearedFields[video.FieldConvertURL] = struct{}{}
}

// ConvertURLCleared returns if the "convert_url" field was cleared in this mutation.
func (m *VideoMutation) ConvertURLCleared() bool {
	_, ok := m.clearedFields[video.FieldConvertURL]
	return ok
}

// ResetConvertURL resets all changes to the "convert_url" field.
func (m *VideoMutation) ResetConvertURL() {
	m.convert_url = nil
	delete(m.clearedFields, video.FieldConvertURL)
}

// Where appends a list predicates to the VideoMutation builder.
func (m *VideoMutation) Where(ps ...predicate.Video) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the VideoMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *VideoMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Video, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *VideoMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *VideoMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Video).
func (m *VideoMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *VideoMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.origin_url != nil {
		fields = append(fields, video.FieldOriginURL)
	}
	if m.convert_url != nil {
		fields = append(fields, video.FieldConvertURL)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *VideoMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case video.FieldOriginURL:
		return m.OriginURL()
	case video.FieldConvertURL:
		return m.ConvertURL()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *VideoMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case video.FieldOriginURL:
		return m.OldOriginURL(ctx)
	case video.FieldConvertURL:
		return m.OldConvertURL(ctx)
	}
	return nil, fmt.Errorf("unknown Video field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *VideoMutation) SetField(name string, value ent.Value) error {
	switch name {
	case video.FieldOriginURL:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetOriginURL(v)
		return nil
	case video.FieldConvertURL:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetConvertURL(v)
		return nil
	}
	return fmt.Errorf("unknown Video field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *VideoMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *VideoMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *VideoMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Video numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *VideoMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(video.FieldConvertURL) {
		fields = append(fields, video.FieldConvertURL)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *VideoMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *VideoMutation) ClearField(name string) error {
	switch name {
	case video.FieldConvertURL:
		m.ClearConvertURL()
		return nil
	}
	return fmt.Errorf("unknown Video nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *VideoMutation) ResetField(name string) error {
	switch name {
	case video.FieldOriginURL:
		m.ResetOriginURL()
		return nil
	case video.FieldConvertURL:
		m.ResetConvertURL()
		return nil
	}
	return fmt.Errorf("unknown Video field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *VideoMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *VideoMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *VideoMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *VideoMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *VideoMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *VideoMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *VideoMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Video unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *VideoMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Video edge %s", name)
}
