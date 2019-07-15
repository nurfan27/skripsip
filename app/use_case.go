package app

import (
	"log"
	nlp "skripsip/nlp"
)

type UseCase struct {
	repository *Repository
}

func (uc *UseCase) Handle(request WhatsappChatRequest) WhatsappChatResponse {
	var response WhatsappChatResponse

	request.From = RemoveCountryCode(request.From)

	log.Println("request : ")
	log.Println(request)

	auth := uc.repository.CheckPhoneNumber(request.From)

	log.Println(auth)

	if !auth {
		response.Autoreply = MESSAGE[STATUS_NOT_AUTH]
		return response
	}

	// Proses Case Folding
	caseFoldingResult := nlp.CaseFolding(request.Text)
	// log.Println("caseFolding Result : ")
	// log.Println(caseFoldingResult)

	// Proses Tokenizing
	tokenizingResult := nlp.Tokenizing(caseFoldingResult)
	// log.Println("Tokenizing Result : ")
	// log.Println(tokenizingResult)

	// Proses Filtering
	filteringResult := nlp.WordList(tokenizingResult)
	// log.Println("WordList Result : ")
	// log.Println(filteringResult)

	uc.repository.SetPhoneNumber(request.From)
	answer := uc.repository.Handle(filteringResult)

	response.Autoreply = answer

	return response
}

func NewUseCase() *UseCase {
	var usecase UseCase

	usecase.repository = NewRepository()

	return &usecase
}
