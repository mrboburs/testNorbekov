package model

type TablePost struct {
	ID          int    `json:"-"`
	PostTitle   string `json:"post_title"`
	PostBody    string `json:"post_body"`
	PostTitleRu string `json:"post_title_ru"`
	PostBodyRu  string `json:"post_body_ru"`
	Date        string `json:"date" `
	Format      string `json:"format" `
	PostImgUrl  string `json:"post_img_url"`
	Price       string `json:"price" `
	Duration    string `json:"duration" `
}

type TableFull struct {
	ID        int    `json:"id" db:"id"`
	PostTitle string `json:"post_title" db:"post_title"`

	PostBodyRu  string `json:"post_body_ru" db:"post_body_ru"`
	PostTitleRu string `json:"post_title_ru" db:"post_title_ru"`
	PostImgPath string `json:"post_img_path" db:"post_img_path"`
	PostImgUrl  string `json:"post_img_url" db:"post_img_url"`
	PostBody    string `json:"post_body" db:"post_body"`
	Date        string `json:"date" db:"date"`
	Format      string `json:"format" db:"format"`
	PostDate    string `json:"post_date" db:"post_date"`
	Price       string `json:"price" db:"price"`
	Duration    string `json:"duration" db:"duration"`
}
type allTable struct {
	AllHome []TableFull
}

type CourseFull struct {
	Title   string `json:"title" `
	TitleRu string `json:"title_ru"`
	Body    string `json:"body" `

	BodyRu   string `json:"body_ru"`
	Price    string `json:"price" `
	Duration string `json:"duration" `
	Term     string `json:"term" `
	Format   string `json:"format" `
	Date     string `json:"date" `
}

type CourseFull1 struct {
	ID       int    `json:"id" db:"id"`
	Title    string `json:"title" db:"title"`
	Body     string `json:"body" db:"body"`
	BodyRu   string `json:"body_ru" db:"body_ru"`
	TitleRu  string `json:"title_ru" db:"title_ru"`
	Price    string `json:"price" db:"price"`
	Duration string `json:"duration" db:"duration"`
	Term     string `json:"term" db:"term"`
	Format   string `json:"format" db:"format"`
	Date     string `json:"date" db:"created_at"`
}
