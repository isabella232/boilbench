package gorps

import (
	"gopkg.in/nullbio/null.v4"

	// Shutup linter
	_ "github.com/lib/pq"
)

// Pilot struct
type Pilot struct {
	ID   int
	Name string
}

// Jet struct
type Jet struct {
	ID         int
	PilotID    int `db:"pilot_id"`
	AirportID  int `db:"airport_id"`
	Name       string
	Color      null.String
	UUID       string
	Identifier string
	Cargo      []byte
	Manifest   []byte
}

// Airport struct
type Airport struct {
	ID   int
	Size null.Int
}

// License struct
type License struct {
	ID      int
	PilotID int `db:"pilot_id"`
}

// Hangar struct
type Hangar struct {
	ID   int
	Name string
}

// Language struct
type Language struct {
	ID       int
	Language string
}
