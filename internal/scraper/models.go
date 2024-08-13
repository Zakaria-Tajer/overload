package scraper

type Author struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Manga []Manga `json:"manga"`
}
type Manga struct {
	ID     int      `json:"id"`
	Title  string   `json:"title"`
	Status string   `json:"status"`
	Genre  []string `json:"genre"`
	Type   string   `json:"type"`
	Artist string   `json:"artist"`
	Chapters []Chapters `json:"chapters"`
}

type Chapters struct {
	ID    int    `json:"id"`
	Title string `json:"name"`
	Count int    `json:"count"`
	Date  string `json:"date"`
}
