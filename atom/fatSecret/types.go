package atom_fatSecret

import (
	"net/http"
	"sync"
	"time"
)

type fatSecretTokenResponseModel struct {
	Access_Token string `json:"access_token"`
	Expires_In   int    `json:"expires_in"`
}

type tokenCacheModel struct {
	Token      string
	ExpiryTime time.Time
	mutex      sync.Mutex
	httpClient *http.Client
}
