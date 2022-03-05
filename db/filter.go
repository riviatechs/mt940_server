package db

import (
	"context"
	"fmt"

	"github.com/riviatechs/mt940_server/models"
)

func StatementsFiltered(ctx context.Context, input *models.FilterInput) ([]*models.ConfGroup, error) {
	if input == nil {
		return Statements(ctx)
	}

	db := Db
	if input.Currency != nil {
		db = db.Where("currency LIKE ?", *input.Currency)
	}

	if input.Tt != nil {
		db = db.Where("mark LIKE ?", *input.Tt)
	}

	if input.Period != nil {
		if input.Period.Date != nil {
			db = db.Where("date_time >= ?::date AND date_time <(?::date + '1 day'::interval)", *input.Period.Date, *input.Period.Date)

		} else {
			if input.Period.Start != nil && input.Period.End != nil {
				db = db.Where("date_time >= ?::date AND date_time <(?::date)", input.Period.Start, input.Period.End)
			}
		}
	}

	if input.AmountRange != nil {
		if input.AmountRange.Amount != nil {

			db = db.Where("ceil(amount) = ceil((?::NUMERIC))", *input.AmountRange.Amount)
		} else {
			if input.AmountRange.Lower != nil && input.AmountRange.Upper != nil {

				db = db.Where("ceil(amount) >= ceil((?::NUMERIC)) AND amount <= ceil((?::NUMERIC))", *input.AmountRange.Lower, *input.AmountRange.Upper)
			}
		}
	}

	var confirmations []*models.Confirmation

	db.Order("date_time desc").Find(&confirmations)

	if confirmations == nil {
		return nil, fmt.Errorf("failed to get statements")
	}

	return GroupStmtsByDate(confirmations)
}
