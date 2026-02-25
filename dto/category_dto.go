package dto

// `form` tegi** — `json` o'rniga `form` ishlatdik, chunki image yuklashda
// `multipart/form-data` formatida keladi, `application/json`
type CreateCategoryRequest struct {
	Name        string `form:"name" binding:"required,min=2,max=100"`
	Description string `form:"description"`
	Color       string `form:"color"`
}

// omitempty -> barchasi o'zgarishi shart emas ya'ni barcha fieldni o'zgartirish majburiy emas
type UpdateCategoryRequest struct {
	Name        string `form:"name" binding:"omitempty,min=2,max=100"`
	Description string `form:"description"`
	Color       string `form:"color"`
}

type CategoryFilter struct {
	Search  string `form:"search"`   // name bo'yicha qidirish
	SortBy  string `form:"sort_by"`  // name, created_at
	SortDir string `form:"sort_dir"` // asc, desc
}

type CategoryResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
	Image       string `json:"image"`
}
