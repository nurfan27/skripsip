package app

type BrivaResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		NomorBriva string `json:"nomor_briva"`
		Nama       string `json:"nama"`
	} `json:"data"`
}

type WhatsappChatRequest struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Event string `json:"event"`
	Text  string `json:"text"`
}
