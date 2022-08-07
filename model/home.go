package model

// import "database/sql"

// import "database/sql"

// import "database/sql"

// "database/sql"
// "github.com/lib/pq"

type HomePost struct {
	ID          int    `json:"-"`
	PostTitle   string `json:"post_title"`
	PostTitleRu string `json:"post_title_ru"`
	PostImgUrl  string `json:"post_img_url"`
	PostBody    string `json:"post_body"`
	PostBodyRu  string `json:"post_body_ru"`
}

type HomeFull struct {
	ID          int    `json:"id" db:"id"`
	PostTitle   string `json:"post_title" db:"post_title"`
	PostBodyRu  string `json:"post_body_ru" db:"post_body_ru"`
	PostTitleRu string `json:"post_title_ru" db:"post_title_ru"`
	PostImgPath string `json:"post_img_path" db:"post_img_path"`
	PostImgUrl  string `json:"post_img_url" db:"post_img_url"`
	PostBody    string `json:"post_body" db:"post_body"`
	PostDate    string `json:"post_date" db:"post_date"`
}

type allHome struct {
	AllHome []HomeFull
}
