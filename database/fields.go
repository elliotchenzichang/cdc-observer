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
type Json json.RawMessage

// todo try to make this definition more readable
type Field struct {
	Name       string
	Type       string
	PrimaryKey bool
	NullAble   bool
}

const (
	TINY_INT   = "TINYINT"
	SMALL_INT  = "SMALLINT"
	MEDIUM_INT = "MEDIUMINT"
	INT        = "INT"
	BIG_INT    = "BIGINT"
	DECIMAL    = "DECIMAL"
	FLOAT      = "FLOAT"
	DOUBLE     = "DOUBLE"
	DTAE       = "DATE"
	TIME       = "TIME"
	YEAR       = "YEAR"
	DATETIME   = "DATETIME"
	TIMESTAMP  = "TIMESTAMP"
	CHAR       = "CHAR"
	VARCHAR    = "VARCHAR"
	TEXT       = "TEXT"
	BLOB       = "BLOB"
	ENUM       = "ENUM"
	SET        = "SET"
	POINT      = "POINT"
	LINESTRING = "LINESTRING"
	PLOYGON    = "PLOYGON"
	JSON       = "JSON"
)
