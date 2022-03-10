package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/riviatechs/mt940_server/models"
)

func Search(ctx context.Context, input string) ([]*models.Confirmation, error) {
	db := Db.Where("searchable_part_b_name @@ websearch_to_tsquery(?) OR searchable_part_b_account @@ websearch_to_tsquery(?) ", input, input)

	var confirmations []*models.Confirmation

	db.Order("date_time desc").Find(&confirmations)

	if confirmations == nil {
		return nil, fmt.Errorf("failed to get statements")
	}

	return confirmations, nil
}

func SearchWithFilters(ctx context.Context, input *models.FilterInput) ([]*models.Confirmation, error) {
	db := Db

	if input.Tt != nil {
		tt := strings.ToUpper(*input.Tt)
		db = db.Where("mark LIKE ?", tt)
	}

	db = db.Where("searchable_part_b_name @@ websearch_to_tsquery(?) OR searchable_part_b_account @@ websearch_to_tsquery(?) ", *input.Search, *input.Search)

	var confirmations []*models.Confirmation

	db.Order("date_time desc").Find(&confirmations)

	if confirmations == nil {
		return nil, fmt.Errorf("failed to get statements")
	}

	return confirmations, nil
}
