package models

// type Quote struct {
// 	Content        string   `json:"content"`
// 	Author_details Author   `json:"author_details"`
// 	Tags           []string `json:"tags"`
// }

type Quote struct {
	Content        string   `json:"content"`
	Author_details Author   `json:"author_details"`
	Tags           []string `json:"tags"`
}

type Author struct {
	Author_name   string `json:"author_name"`
	Born_date     string `json:"born_date"`
	Born_Location string `json:"born_location"`
	Desciption    string `json:"description"`
}
