package app

import (
	"qasircore"
)

type Model struct {
	Db         *qasircore.Db
	repository *Repository
}

func (m *Model) FindSentence(query string) IntentDataset {
	var data IntentDataset

	db := m.Db.GetConnection()
	db.LogMode(true)
	db.Raw(query).Scan(&data)
	return data
}

func NewModel() *Model {
	var model Model
	configs := Env.GetConfig()

	model.Db = qasircore.Database(configs)

	return &model
}
