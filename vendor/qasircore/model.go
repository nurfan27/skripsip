package qasircore

type Model struct {
	Db *Db
}

func (m *Model) SetMerchantID(merchantID string) {
	gormDb := m.Db.GetConnection()

	gormDb = gormDb.Where("merchant_id = ?", merchantID)

	m.Db.SetConnection(gormDb)
}

func (m *Model) GetDb() *Db {
	return m.Db
}

func (m *Model) SetDb(Db *Db) {
	m.Db = Db
}
