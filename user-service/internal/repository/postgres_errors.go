package repository

import (
	"github.com/lib/pq"
)

func isUniqueViolation(err error) bool {
	if pqError, ok := err.(*pq.Error); ok {
		if pqError.Code == "23505" { // unique_violation
			return true
		}
	}

	return false
}
