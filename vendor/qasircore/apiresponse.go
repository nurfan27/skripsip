package qasircore

const (
	STATUS_SUCCESS                    = 1
	STATUS_ERR_INVALID_SUBDOMAIN      = 2
	STATUS_ERR_INVALID_LOGIN          = 3
	STATUS_ERR_INVALID_PARAMETER      = 4
	STATUS_ERR_INACTIVE_ACCOUNT       = 6
	STATUS_ERR_INVALID_SESSION        = 7
	STATUS_ERR_INVALID_PERMISSION     = 8
	STATUS_ERR_INVALID_ACCESS         = 9
	STATUS_ERR_DUPLICATE_LOGIN        = 10
	STATUS_ERR_INVALID_USER           = 11
	STATUS_ERR_INVALID_DATE           = 12
	STATUS_ERR_INVALID_OUTLET         = 13
	STATUS_ERR_INVALID_AREA           = 14
	STATUS_ERR_INACTIVE_SUBDOMAIN     = 15
	STATUS_ERR_PROCESS_SUBDOMAIN      = 16
	STATUS_ERR_SUSPENDED_SUBDOMAIN    = 17
	STATUS_ERR_INVALID_TRANSACTION    = 21
	STATUS_ERR_DUPLICATE_TRANSACTION  = 22
	STATUS_DATA_IS_EMPTY              = 23
	STATUS_ERR_USER_DELETED           = 24
	STATUS_ERR_USER_DELETE_IN_OUTLET  = 25
	STATUS_ERR_USER_DELETE_VALiDATION = 26
	STATUS_ERR_VALIDATION             = 40
	STATUS_ERR_NOT_FOUND              = 60
	STATUS_ERR_INVALID_METHOD         = 66
	STATUS_ERR_PROCESS                = 88
	STATUS_ERR_INVALID_KEY            = 90
	STATUS_ERR_SYSTEM                 = 99
)

var STATUS_MESSAGE = map[int]string{
	1:  "Berhasil",
	2:  "Nama Toko kamu salah",
	3:  "Username atau password kamu salah",
	4:  "Parameter tidak valid",
	6:  "Akun tidak aktif",
	7:  "Invalid session",
	8:  "Akses tidak dibolehkan atau kamu tidak memiliki akses",
	9:  "Akses tidak valid karena token tidak cocok atau login sudah kadaluarsa",
	10: "Akses tidak valid karena akun kamu digunakan untuk login di device lain.",
	11: "User tidak valid",
	12: "Tanggal transaksi tidak valid",
	13: "Outlet tidak valid",
	14: "Area tidak valid atau area tidak ditemukan",
	15: "Toko belum diaktifkan",
	16: "Toko masih dalam proses pembuatan",
	17: "Toko kamu diblokir, silahkan hubungi Qasir melalui email hello@qasir.id",
	21: "Data transaksi tidak valid atau tidak sesuai",
	22: "Data transaksi duplikat",
	23: "Data tidak ditemukan",
	24: "User ini telah dihapus, mohon hubungi atasan kamu.",
	25: "Kamu tidak terdaftar di outlet ini",
	26: "User ini telah dihapus, mohon hubungi atasan kamu.",
	40: "Validasi tidak cocok",
	57: "Deposit kamu tidak mencukupi",
	60: "Resource not found",
	66: "Bad Request",
	88: "Proses gagal",
	90: "API key tidak valid",
	99: "Internal system error",
}

type Response struct {
	HttpStatus     int
	StatusResponse int
	Data           map[string]interface{}
}

type ResponseData struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Token   map[string]interface{} `json:"token"`
	Data    map[string]interface{} `json:"data"`
}

type ApiResponse struct {
	status        int
	messageStatus map[int]string
	message       string
	data          map[string]interface{}
	token         map[string]interface{}
}

func (api *ApiResponse) SetMessageStatus(messageStatus map[int]string) {
	api.messageStatus = messageStatus
}

func (api *ApiResponse) SetStatus(status int) {
	api.status = status

	if api.messageStatus == nil {
		api.SetMessage(STATUS_MESSAGE[api.status])
	}
}

func (api *ApiResponse) SetMessage(message string) {
	api.message = message
}

func (api *ApiResponse) SetData(data map[string]interface{}) {
	api.data = data
}

func (api *ApiResponse) SetToken(tokendata map[string]interface{}) {
	api.token = tokendata
}

func (api *ApiResponse) ToResponse() ResponseData {
	var res ResponseData

	res.Status = api.status
	res.Message = api.message
	res.Data = api.data
	res.Token = api.token

	api.data = nil

	return res
}
