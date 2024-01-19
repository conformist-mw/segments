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
	CompanyID uint   `gorm:"foreignkey:CompanyID;references:ID;not null;uniqueIndex:idx_section_name_company_uniq" json:"company"`
}

func (Section) TableName() string {
	return "segments_section"
}

type Rack struct {
	ID        uint   `gorm:"primarykey;not null" json:"id"`
	Name      string `gorm:"type:varchar(30);not null;uniqueIndex:idx_rack_name_section_uniq" json:"name"`
	SectionID uint   `gorm:"foreignkey:SectionID;references:ID;not null;uniqueIndex:idx_rack_name_section_uniq" json:"section"`
}

func (Rack) TableName() string {
	return "segments_rack"
}

func GetCompanies() []Company {
	var companies []Company
	DB.Find(&companies)
	return companies
}

func GetCompany(companySlug string) Company {
	var company Company
	DB.Where(&Company{Slug: companySlug}).First(&company)
	return company
}

type SectionWithAmount struct {
	Section
	SegmentsCount int
	SquareSum     float64
	RacksCount    int
}

func GetSections(companySlug string) []SectionWithAmount {
	var sections []SectionWithAmount

	DB.Table(new(Section).TableName()).
		Select(
			"segments_section.*, "+
				"count(distinct segments_rack.id) as racks_count, "+
				"count(distinct segments_segment.id) as segments_count, "+
				"sum(distinct segments_segment.square) as square_sum").
		Joins(
			"join segments_company "+
				"on segments_company.id = segments_section.company_id "+
				"AND segments_company.slug = ?",
			companySlug).
		Joins("left join segments_rack on segments_rack.section_id = segments_section.id").
		Joins("left join segments_segment on segments_segment.rack_id = segments_rack.id").
		Group("segments_section.id").
		Scan(&sections)

	return sections
}
