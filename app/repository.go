package app

import (
	"fmt"
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

	switch intent.IntentID {
	case 1:
		return ""
	case 2:
		return r.findBriva()
	case 3:
		return ""
	case 4:
		return r.findDospem()
	default:
		return MESSAGE[STATUS_KNOWLEDGE_NOTFOUND]
	}

}

func (r *Repository) findDospem() string {
	resp := r.serviceSiakad.GetDospem(r.phoneNumber)

	if resp.Status != 1 {
		return MESSAGE[STATUS_ERROR_SYSTEM]
	}

	answer := fmt.Sprintf("Data dosen pembimbing akademik anda : \n\nNama :  %s \n\nNID :  %s \n\nNo.Tlp : %s \n", resp.Data.NamaDosen, resp.Data.Nid, resp.Data.TlpDosen)

	return answer
}

func (r *Repository) findBriva() string {

	resp := r.serviceSiakad.GetBriva(r.phoneNumber)

	if resp.Status != 1 {
		return MESSAGE[STATUS_ERROR_SYSTEM]
	}

	answer := fmt.Sprintf("nomer briva %s adalah %s", resp.Data.Nama, resp.Data.NomorBriva)

	return answer
}

func (r *Repository) getIntents(query string) IntentDataset {

	intent := r.model.FindSentence(query)

	return intent
}

func (r *Repository) generateQuery(params []string) string {
	where := " where 1=1 "

	for _, param := range params {
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
