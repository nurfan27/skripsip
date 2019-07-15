package app

const (
	STATUS_SUCCESS            = 1
	STATUS_EMPTY_DATA         = 23
	STATUS_ERROR_SYSTEM       = 97
	STATUS_KNOWLEDGE_NOTFOUND = 98
	STATUS_NOT_AUTH           = 99
)

var MESSAGE = map[int]string{
	23: "Maaf, Siakad tidak dapat menampilkan data kamu",
	97: "Terjadi kesalah sistem, mohon coba beberapa saat lagi.",
	98: "Maaf, untuk saat ini pertanyaan anda belum terdapat pada knowledge base kami. \n\n Ada yang lain yang bisa kami bantu.",
	99: "Maaf nomer anda tidak terdaftar pada sistem kami. \n\n Mohon isi data nomer hp anda pada biodata sistem informasi akademik ubhara",
}

type BrivaResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		NomorBriva string `json:"nomor_briva"`
		Nama       string `json:"nama"`
	} `json:"data"`
}

type DospemResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		NamaDosen string `json:"nama_dosen"`
		Nid       string `json:"nid"`
		TlpDosen  string `json:"tlp_dosen"`
	} `json:"data"`
}

type SppResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Tahunajaran string `json:"tahunajaran"`
		TglValidasi string `json:"tgl_validasi"`
		Status      string `json:"status"`
		Semester    string `json:"semester"`
	} `json:"data"`
}

type WhatsappChatRequest struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Event string `json:"event"`
	Text  string `json:"text"`
}

type WhatsappChatResponse struct {
	Autoreply string `json:"autoreply"`
}
