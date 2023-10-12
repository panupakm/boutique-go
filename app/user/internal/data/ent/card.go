// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/panupakm/boutique-go/app/user/internal/data/ent/card"
	"github.com/panupakm/boutique-go/app/user/internal/data/ent/user"
)

// Card is the model entity for the Card schema.
type Card struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Number holds the value of the "number" field.
	Number string `json:"number,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Ccv holds the value of the "ccv" field.
	Ccv int `json:"ccv,omitempty"`
	// ExpirationYear holds the value of the "expiration_year" field.
	ExpirationYear int `json:"expiration_year,omitempty"`
	// ExpirationMonth holds the value of the "expiration_month" field.
	ExpirationMonth int `json:"expiration_month,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CardQuery when eager-loading is set.
	Edges        CardEdges `json:"edges"`
	user_cards   *uuid.UUID
	selectValues sql.SelectValues
}

// CardEdges holds the relations/edges for other nodes in the graph.
type CardEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CardEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Card) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case card.FieldCcv, card.FieldExpirationYear, card.FieldExpirationMonth:
			values[i] = new(sql.NullInt64)
		case card.FieldNumber, card.FieldName:
			values[i] = new(sql.NullString)
		case card.FieldID:
			values[i] = new(uuid.UUID)
		case card.ForeignKeys[0]: // user_cards
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Card fields.
func (c *Card) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case card.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				c.ID = *value
			}
		case card.FieldNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field number", values[i])
			} else if value.Valid {
				c.Number = value.String
			}
		case card.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case card.FieldCcv:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ccv", values[i])
			} else if value.Valid {
				c.Ccv = int(value.Int64)
			}
		case card.FieldExpirationYear:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field expiration_year", values[i])
			} else if value.Valid {
				c.ExpirationYear = int(value.Int64)
			}
		case card.FieldExpirationMonth:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field expiration_month", values[i])
			} else if value.Valid {
				c.ExpirationMonth = int(value.Int64)
			}
		case card.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_cards", values[i])
			} else if value.Valid {
				c.user_cards = new(uuid.UUID)
				*c.user_cards = *value.S.(*uuid.UUID)
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Card.
// This includes values selected through modifiers, order, etc.
func (c *Card) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Card entity.
func (c *Card) QueryOwner() *UserQuery {
	return NewCardClient(c.config).QueryOwner(c)
}

// Update returns a builder for updating this Card.
// Note that you need to call Card.Unwrap() before calling this method if this Card
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Card) Update() *CardUpdateOne {
	return NewCardClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Card entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Card) Unwrap() *Card {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Card is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Card) String() string {
	var builder strings.Builder
	builder.WriteString("Card(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("number=")
	builder.WriteString(c.Number)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("ccv=")
	builder.WriteString(fmt.Sprintf("%v", c.Ccv))
	builder.WriteString(", ")
	builder.WriteString("expiration_year=")
	builder.WriteString(fmt.Sprintf("%v", c.ExpirationYear))
	builder.WriteString(", ")
	builder.WriteString("expiration_month=")
	builder.WriteString(fmt.Sprintf("%v", c.ExpirationMonth))
	builder.WriteByte(')')
	return builder.String()
}

// Cards is a parsable slice of Card.
type Cards []*Card
