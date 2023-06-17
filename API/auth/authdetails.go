package auth

// "encoding/json"

type (
	AuthDetails struct {
		Username   string `json:"username"`
		Cookie     string
		UserAgent  string
		Email      string
		Password   string
		Hashed     bool
		Support2fa bool
		Active     bool
	}

	LegacyAuthDetails struct {
		Username   string
		Cookie     string
		UserAgent  string
		Email      string
		Password   string
		Hashed     bool
		Support2fa bool
		Active     bool
		AuthID     string
		Session    string
		AuthHash   string
		XBC        string
	}

	GenericResponse struct {
		Message string `json:"message"`
		Code    string `json:"code"`
		Status  bool   `json:"status"`
	}
)
