package constants

const (
	MayTapiBaseURL = "https://api.maytapi.com"
)

// API Path
const (
	SendMessagePath = "/api/%s/%s/sendMessage"
)

// Header
const (
	MaytapiAuthHeader = "x-maytapi-key"
)

// Action type
const (
	ACKType     = "ack"
	MessageType = "message"
)

// Message type
const (
	TextType = "text"
)

// Special key to detect intent
const (
	NotRegisteredKey = "NOT_REGISTERED"
)

const (
	HiIntroductionKey = "hi"
	NotFoundKey       = "not-found"
)

var IntroductionChat = map[string][]string{
	HiIntroductionKey: {"Wassap", "Yo", "Hai", "Gass", "Mulai", "Start", "Coba", "Woi", "Testing", "Tes", "Hi", "Halo", "Selamat"},
	NotFoundKey:       {"Not Found"},
}

var MenusChat = map[string][]string{
	"informasi-pengurusan-skck":                   {"1", "Informasi Pengurusan SKCK", "SKCK"},
	"informasi-pembuatan-ktp":                     {"2", "Informasi Pembuatan KTP", "KTP"},
	"informasi-pengurusan-pindah-kk-nikah":        {"3", "Informasi Pengurusan Pindah KK Karena Nikah", "KK", "Nikah"},
	"informasi-pengurusan-surat-hibah-tanah":      {"4", "Informasi Pengurusan Surat Hibah Tanah", "Hibah", "Tanah"},
	"informasi-pengurusan-surat-keterangan-lahir": {"5", "Informasi Pengurusan Surat Keterangan Lahir", "Lahir"},
	"informasi-data-diri":                         {"6", "Informasi Data Diri", "Data Diri"},
}
