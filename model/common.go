package model

type DataState int64

const (
	STATE_EXIST   DataState = 1
	STATE_DELETED DataState = 2
)

// ContentType
const (
	TextPlain   = "text/plain" //content
	ImageUrl    = "image/url"
	ImageJpg    = "image/jpg"
	ImageJpeg   = "image/jpeg"
	ImagePng    = "image/png"
	ImageGif    = "image/gif"
	ImageWebp   = "image/webp"
	AudioWav    = "audio/wav"
	Transaction = "transaction"
	BsvProtocol = "bsvProtocol"
)
