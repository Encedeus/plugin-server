package dto

type CreatePluginDTO struct {
	Name         string   `json:"name"`
	Desc         string   `json:"description"`
	Repo         string   `json:"repo"`
	Homepage     string   `json:"homepage"`
	Contributors []string `json:"contributors"`
}
type UpdatePluginDTO struct {
	Name         string   `json:"name"`
	Desc         string   `json:"description"`
	Repo         string   `json:"repo"`
	Homepage     string   `json:"homepage"`
	Contributors []string `json:"contributors"`
}
