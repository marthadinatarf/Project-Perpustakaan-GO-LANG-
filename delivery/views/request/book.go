package request

type InsertBookRequest struct {
	Judul    string `json:"judul"`
	Author   string `json:"author"`
	Penerbit string `json:"penerbit"`
}

type UpdateBookRequest struct {
	Judul    string `json:"judul"`
	Author   string `json:"author"`
	Penerbit string `json:"penerbit"`
}
