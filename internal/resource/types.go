// Package resource provides types for resource management.
package resource

import (
	"time"
)

// Name represents a resource name.
type Name string

// Description represents a resource description.
type Description string

// Tag represents a single resource tag.
type Tag string

type CreatedAt time.Time

// Message represents a resource message.
type Message string

// ID represents a unique resource identifier.
type ID string

// Status represents a resource status.
type Status string

// Common status constants.
const (
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
	StatusPending  Status = "pending"
)

// Pattern represents a filter pattern (regex or glob).
type Pattern string

// Limit represents a maximum number of results.
type Limit int

// Offset represents a pagination offset.
type Offset int

// SortField represents a field name to sort by.
type SortField string
