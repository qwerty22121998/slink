package dtos

type ShortLink struct {
	Name      string `json:"name"`
	URL       string `json:"url"`
	Short     string `json:"short"`
	ExpiredAt int64  `json:"expired_at"`
}

type LinkShortenerResponse struct {
	Message string     `json:"message"`
	Data    *ShortLink `json:"data,omitempty"`
}
