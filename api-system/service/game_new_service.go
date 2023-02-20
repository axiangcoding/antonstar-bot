package service

import (
	"github.com/axiangcoding/antonstar-bot/data/dal"
	"github.com/axiangcoding/antonstar-bot/data/table"
	"github.com/axiangcoding/antonstar-bot/logging"
)

func FindGameNewByLink(link string) (*table.GameNew, error) {
	take, err := dal.Q.GameNew.Where(dal.GameNew.Link.Eq(link)).Take()
	if err != nil {
		return nil, err
	}
	return take, err
}

func MustFindGameNewByLink(link string) *table.GameNew {
	f, err := FindGameNewByLink(link)
	if err != nil {
		logging.L().Warn("find gameNew error", logging.Error(err), logging.Any("link", link))
	}
	return f
}

func MustSaveGameNew(gn *table.GameNew) {
	err := dal.Q.GameNew.Save(gn)
	if err != nil {
		logging.L().Warn("save gameNew error", logging.Error(err))
	}
}
