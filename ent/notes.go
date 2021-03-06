// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/GhvstCode/LR-ENT/ent/notes"
	"github.com/facebook/ent/dialect/sql"
)

// Notes is the model entity for the Notes schema.
type Notes struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "Title" field.
	Title string `json:"Title,omitempty"`
	// Content holds the value of the "Content" field.
	Content string `json:"Content,omitempty"`
	// Private holds the value of the "Private" field.
	Private bool `json:"Private,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Notes) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // Title
		&sql.NullString{}, // Content
		&sql.NullBool{},   // Private
		&sql.NullTime{},   // created_at
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Notes fields.
func (n *Notes) assignValues(values ...interface{}) error {
	if m, n := len(values), len(notes.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	n.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Title", values[0])
	} else if value.Valid {
		n.Title = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Content", values[1])
	} else if value.Valid {
		n.Content = value.String
	}
	if value, ok := values[2].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field Private", values[2])
	} else if value.Valid {
		n.Private = value.Bool
	}
	if value, ok := values[3].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[3])
	} else if value.Valid {
		n.CreatedAt = value.Time
	}
	return nil
}

// Update returns a builder for updating this Notes.
// Note that, you need to call Notes.Unwrap() before calling this method, if this Notes
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *Notes) Update() *NotesUpdateOne {
	return (&NotesClient{config: n.config}).UpdateOne(n)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (n *Notes) Unwrap() *Notes {
	tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("ent: Notes is not a transactional entity")
	}
	n.config.driver = tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *Notes) String() string {
	var builder strings.Builder
	builder.WriteString("Notes(")
	builder.WriteString(fmt.Sprintf("id=%v", n.ID))
	builder.WriteString(", Title=")
	builder.WriteString(n.Title)
	builder.WriteString(", Content=")
	builder.WriteString(n.Content)
	builder.WriteString(", Private=")
	builder.WriteString(fmt.Sprintf("%v", n.Private))
	builder.WriteString(", created_at=")
	builder.WriteString(n.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// NotesSlice is a parsable slice of Notes.
type NotesSlice []*Notes

func (n NotesSlice) config(cfg config) {
	for _i := range n {
		n[_i].config = cfg
	}
}
