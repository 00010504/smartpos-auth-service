package models

type PermissionChild struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

type GetPermissionChildsResponse struct {
	Total uint64             `json:"total"`
	Data  []*PermissionChild `json:"data"`
}
type Permission struct {
	Name               string   `json:"name"`
	PermissionChildIds []string `json:"permission_child_ids"`
}
