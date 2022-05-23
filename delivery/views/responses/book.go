package responses

import "net/http"

type BookResponse struct {
	ID       int    `json:"id"`
	Judul    string `json:"judul"`
	Author   string `json:"author"`
	Penerbit string `json:"penerbit"`
}

func InsertBookSuccess(data BookResponse) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "berhasil insert data buku",
		"status":  true,
		"data":    data,
	}
}

func SelectBookSuccess(data BookResponse) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "berhasil menampilkan buku",
		"status":  true,
		"data":    data,
	}
}
