//
// This file is responsible for interfacing with the external application and
// offers a simple interface to add, remove, filter and retrieve supers from the
// database
//

package database

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jichall/levpay/src/models"
)

// Supers returns all supers from the database
func Supers() []models.Super {
	var supers []models.Super

	rows, err := query(querySelectSupers, nil)

	if err == nil {
		for rows.Next() {
			s := models.Super{}

			err = rows.Scan(&s.UUID, &s.Name, &s.FullName, &s.Intelligence,
				&s.Power, &s.Occupation, &s.Image)

			if err != nil {
				logger.Warning("failed to scan values, reason %v", err)
			} else {
				supers = append(supers, s)
			}
		}
	}

	return supers
}

// AddSuper adds a super on the database
func AddSuper(super models.Super) {
	execute(queryAddSuper, super.UUID, super.Name, super.FullName,
		super.Intelligence, super.Power, super.Occupation, super.Image)
}

// RemoveSuper removes a super from the database.
func RemoveSuper(param string) {
	execute(queryRemoveSuper, param)
}

// FilterSuper retrieves the super with the specified param, either a name or an
// unique identifier.
func FilterSuper(param string) models.Super {
	super := models.Super{}

	var err error
	var rows pgx.Rows

	if _, err := uuid.Parse(param); err != nil {
		// not an uuid, a name most probably
		rows, err = query(queryFilterNameSuper, param)
	} else {
		rows, err = query(queryFilterUUIDSuper, param)
	}

	if err != nil {
		logger.Error("couldn't filter super, reason %v", err)
	} else {
		for rows.Next() {
			var ID string
			rows.Scan(&ID, &super.UUID, &super.Name, &super.FullName,
				&super.Intelligence, &super.Power, &super.Occupation,
				&super.Image)
		}
	}

	return super
}
