package app

import "github.com/mitchellh/mapstructure"

type Mapping struct {
}

func (m *Mapping) WhatsappChatMapping(chat map[string]interface{}) WhatsappChatRequest {
	var result WhatsappChatRequest

	mapstructure.Decode(chat, &result)
	return result
}

func NewMapping() *Mapping {
	var mapping Mapping
	return &mapping
}
