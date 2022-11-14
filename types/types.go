package types

const (
	// User roles ENUM stored in DB
	Role_Admin = "ADMIN"
	Role_Read  = "READ"
	Role_Add   = "ADD"
	Role_Edit  = "EDIT"
)

type (
	// For Login API to validate user and fetch role & name
	LoginReq struct {
		UserID   string `json:"userID,omitempty"`
		UserHash string `json:"userHash,omitempty"`
	}
	LoginResp struct {
		UserRole string `json:"userRole,omitempty"`
		UserName string `json:"userName,omitempty"`
		Tokken   string `json:"tokken,omitempty"`
	}
)
