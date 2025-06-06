package models

type SearchForm struct {
	ColorType   string `form:"color_type"`
	Color       string `form:"color"`
	Width       int    `form:"width"`
	Height      int    `form:"height"`
	OrderNumber string `form:"order_number"`
	Deleted     string `form:"deleted"`
	Page        int    `form:"page"`
}

func (s SearchForm) GetPage() int {
	if s.Page <= 0 {
		return 1
	}
	return s.Page
}

type AddForm struct {
	Width     int    `form:"width" binding:"required"`
	Height    int    `form:"height" binding:"required"`
	Color     string `form:"color" binding:"required"`
	ColorType string `form:"color_type" binding:"required"`
	RackID    int    `form:"rack_id" binding:"required"`
}

type MoveForm struct {
	Rack int `form:"rack" binding:"required"`
}

type RemoveForm struct {
	OrderNumber string `form:"order_number"`
	Defect      string `form:"defect"`
	Description string `form:"description"`
}

type PrintForm struct {
	PrintRack int `form:"print_rack"`
}

type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type ColorForm struct {
	Name      string `form:"name" binding:"required"`
	Slug      string `form:"slug" binding:"required,slug"`
	ColorType string `form:"color_type" binding:"required"`
}
