package models

type Company struct {
	ID       uint   `gorm:"primarykey;not null" json:"id"`
	Name     string `gorm:"type:varchar(30);not null;unique" json:"name"`
	Slug     string `gorm:"type:varchar(50);not null" json:"slug"`
	ImageUrl string `gorm:"type:varchar(255);not null" json:"image_url"`
}

func (Company) TableName() string {
	return "segments_company"
}

type Section struct {
	ID        uint   `gorm:"primarykey;not null" json:"id"`
	Name      string `gorm:"type:varchar(30);not null;uniqueIndex:idx_section_name_company_uniq" json:"name"`
	Slug      string `gorm:"type:varchar(50);not null" json:"slug"`
	CompanyID uint   `gorm:"foreignkey:CompanyID;not null;uniqueIndex:idx_section_name_company_uniq" json:"company"`
}

func (Section) TableName() string {
	return "segments_section"
}

type SectionExcludedColors struct {
	ID          uint `gorm:"primarykey;not null" json:"id"`
	SectionID   uint `gorm:"foreignkey:SectionID;not null" json:"section"`
	ColortypeID uint `gorm:"foreignkey:ColorTypeID;not null" json:"color_type"`
}

func (SectionExcludedColors) TableName() string {
	return "segments_section_excluded_colors"
}

type Rack struct {
	ID        uint   `gorm:"primarykey;not null" json:"id"`
	Name      string `gorm:"type:varchar(30);not null;uniqueIndex:idx_rack_name_section_uniq" json:"name"`
	SectionID uint   `gorm:"foreignkey:SectionID;not null;uniqueIndex:idx_rack_name_section_uniq" json:"section"`
}

func (Rack) TableName() string {
	return "segments_rack"
}
