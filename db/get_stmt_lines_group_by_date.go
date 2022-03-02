package db

import (
	"context"

	"github.com/riviatechs/mt940_server/models"
)

func GetStmtLineGroupedByDate(ctx context.Context) ([]*models.SlGroups, error) {
	var sl []*models.Sl
	Db.Find(&sl)

	sls := make(map[string][]*models.Sl)
	for _, v := range sl {
		if l, ok := sls[*v.ValueDate]; ok {
			s := append(l, v)
			sls[*v.ValueDate] = s
		} else {
			var vv []*models.Sl = sls[*v.ValueDate]
			vv = append(vv, v)
			sls[*v.ValueDate] = vv
		}
	}

	var slg []*models.SlGroups

	for k, v := range sls {
		slgb := models.SlGroups{ValueDate: k, Sls: v}
		slg = append(slg, &slgb)
	}

	return slg, nil
}
