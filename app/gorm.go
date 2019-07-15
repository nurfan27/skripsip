package app

type IntentDataset struct {
	ID       int    `gorm:"column:id"`
	IntentID int    `gorm:"column:intent_id"`
	Sentence string `gorm:"column:sentence"`
}
