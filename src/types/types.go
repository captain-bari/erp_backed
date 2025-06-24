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

	//For getting Materials
	Material struct {
		ID         int      `json:"id,omitempty"`
		Name       string   `json:"name,omitempty"`
		Type       string   `json:"type,omitempty"`
		Properties []string `json:"userID,omitempty"`
		Processes  []string `json:"processes,omitempty"`
		Purchage   []string `json:"purchase,omitempty"`
		Sell       []string `json:"sell,omitempty"`
	}

	GetMaterialsResp struct {
		Materials []Material `json:"materials,omitempty"`
	}
)
