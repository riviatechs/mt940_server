package db

import (
	"context"

	"github.com/riviatechs/mt940_server/models"
	"github.com/riviatechs/mt940_server/util"
)

func GetStmtLineGroupedByDate(ctx context.Context) ([]*models.SlGroups, error) {
	var sl []*models.Sl
	Db.Find(&sl)

	return GroupStmtLinesByDate(sl)
}

func GroupStmtLinesByDate(sl []*models.Sl) ([]*models.SlGroups, error) {
	sls := make(map[string][]*models.Sl)

	for _, v := range sl {
		if l, ok := sls[v.ValueDate.Format(util.TimeFormat)]; ok {
			s := append(l, v)
			sls[v.ValueDate.Format(util.TimeFormat)] = s
		} else {
			vv := sls[v.ValueDate.Format(util.TimeFormat)]
			vv = append(vv, v)
			sls[v.ValueDate.Format(util.TimeFormat)] = vv
		}
	}

	var slg []*models.SlGroups

	for k, v := range sls {
		slgb := models.SlGroups{ValueDate: k, Sls: v}
		slg = append(slg, &slgb)
	}

	return slg, nil
}

func GetStmtLinesFilterByPeriod(ctx context.Context, amount models.Amount) ([]*models.SlGroups, error) {
	var sl []*models.Sl
	Db.Where("amount BETWEEN ? AND ?", amount.Lower, amount.Upper).Find(&sl)

	return GroupStmtLinesByDate(sl)
}
