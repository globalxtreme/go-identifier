package data

var (
	Access AccessIdentifierData
)

type AccessIdentifierData struct {
	Roles       *map[string]interface{} `json:"roles"`
	Permissions *map[string]interface{} `json:"permissions"`
}
