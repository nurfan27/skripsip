package app

import (
	"log"
	"strings"
)

type Repository struct {
	serviceSiakad   *SiakadService
	serviceWhatsapp *WhatsappService
	model           *Model
	phoneNumber     string
}

func (r *Repository) SetPhoneNumber(value string) {
	r.phoneNumber = value
}

func (r *Repository) GetPhoneNumber() string {
	return r.phoneNumber
}

func (r *Repository) Handle(params []string) string {
	if len(params) < 1 {
		return MESSAGE[STATUS_KNOWLEDGE_NOTFOUND]
	}

	// generate query
	q := r.generateQuery(params)

	// get sentence
	intent := r.getIntents(q)
	log.Println("intent : ")
	log.Println(intent)

	switch intent.IntentID {
	case 1:
		return r.findSpp()
	case 2:
		return r.findBriva()
	case 3:
		return r.fetchTranskip()
	case 4:
		return r.findDospem()
	case 6:
		return r.fetchKhs()
	default:
		return MESSAGE[STATUS_KNOWLEDGE_NOTFOUND]
	}

}

func (r *Repository) getIntents(query string) IntentDataset {

	intent := r.model.FindSentence(query)

	return intent
}

func (r *Repository) generateQuery(params []string) string {
	where := " where 1=1 "

	for _, param := range params {
		param = strings.TrimSpace(param)
		where += " AND sentence like \"%" + param + "%\" "
	}

	return "select * from intent_datasets " + where + " LIMIT 1"
}

func (r *Repository) CheckPhoneNumber(phone string) bool {
	resp := r.serviceSiakad.CheckPhoneNumber(phone)

	if resp.Status != 1 {
		return false
	}

	return true
}

func NewRepository() *Repository {
	var repo Repository

	repo.serviceSiakad = NewSiakadService()
	repo.serviceWhatsapp = NewWhatsappService()
	repo.model = NewModel()
	return &repo
}
