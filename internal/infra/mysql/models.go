// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"database/sql"
	"encoding/json"
)

type AccessibilityFeature struct {
	ReviewID sql.NullString
	Feature  sql.NullString
}

type Favorite struct {
	PlaceID sql.NullString
	UserID  sql.NullString
}

type Place struct {
	ID               string
	GooglePlaceID    sql.NullString
	Name             sql.NullString
	FormattedAddress sql.NullString
	Lat              sql.NullFloat64
	Lng              sql.NullFloat64
	Icon             sql.NullString
	Types            json.RawMessage
	OpeningPeriods   json.RawMessage
	Photos           json.RawMessage
	Rating           sql.NullFloat64
	CreatedAt        sql.NullTime
	UpdatedAt        sql.NullTime
}

type Review struct {
	ID        string
	PlaceID   sql.NullString
	UserID    sql.NullString
	Text      sql.NullString
	Images    json.RawMessage
	Rating    sql.NullFloat64
	Reactions json.RawMessage
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}

type User struct {
	ID        string
	Email     string
	Avatar    sql.NullString
	Username  string
	Password  string
	Points    int32
	Missions  json.RawMessage
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}
