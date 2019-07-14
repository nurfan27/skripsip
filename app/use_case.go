package app

import (
	"log"
	nlp "skripsip/nlp"
)

type UseCase struct {
	serviceSiakad   *SiakadService
	serviceWhatsapp *WhatsappService
}

func (uc *UseCase) Handle(request WhatsappChatRequest) {
	log.Println(request)

	// Proses Case Folding
	caseFoldingResult := nlp.CaseFolding(request.Text)
	log.Println("caseFolding Result : ")
	log.Println(caseFoldingResult)

	// Proses Tokenizing
	tokenizingResult := nlp.Tokenizing(caseFoldingResult)
	log.Println("Tokenizing Result : ")
	log.Println(tokenizingResult)

	// Proses Filtering
	filteringResult := nlp.WordList(tokenizingResult)
	log.Println("WordList Result : ")
	log.Println(filteringResult)

	resp := uc.serviceSiakad.CheckPhoneNumber(request.From)
	log.Println(resp)
	return
}

func NewUseCase() *UseCase {
	var usecase UseCase

	usecase.serviceSiakad = NewSiakadService()
	usecase.serviceWhatsapp = NewWhatsappService()
	return &usecase
}
