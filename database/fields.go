package database

import (
	"encoding/json"
	"time"
)

// Numeric types
type TinyInt int8
type SmallInt int16
type MediumInt int32
type Int int32
type BigInt int64
type Decimal string // Consider using a dedicated decimal library
type Float float32
type Double float64

// Date and time types
type Date time.Time
type Time time.Time
type Year int
type DateTime time.Time
type Timestamp time.Time

// String types
type Char string
type Varchar string
type Text string
type Blob []byte
type Enum string
type Set string

// Spatial types (simplified)
type Point struct {
	X, Y float64
}
type LineString []Point
type Polygon [][]Point

// JSON type
type JSON json.RawMessage

// todo try to make this definition more readable
type Field[T Int | TinyInt | SmallInt | MediumInt | BigInt | Decimal |
	Float | Double | Date | Time | Year | DateTime | Timestamp | Char | Varchar | Text | Blob | Enum | Set | Point | LineString | Polygon] struct {
	Name string
	Type T
}
