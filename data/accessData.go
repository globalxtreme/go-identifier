package data

type AccessIdentifierData struct {
	Roles       *map[string]interface{} `json:"roles"`
	Permissions *map[string]interface{} `json:"permissions"`
}
