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
	Racks     []Rack `gorm:"foreignKey:SectionID" json:"racks"`
	Company   Company
}

func (Section) TableName() string {
	return "segments_section"
}

type Rack struct {
	ID        uint   `gorm:"primarykey;not null" json:"id"`
	Name      string `gorm:"type:varchar(30);not null;uniqueIndex:idx_rack_name_section_uniq" json:"name"`
	SectionID uint   `gorm:"foreignkey:SectionID;references:ID;not null;uniqueIndex:idx_rack_name_section_uniq" json:"section"`
	Section   Section
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

func GetSection(sectionSlug string) Section {
	var section Section
	DB.Preload("Racks").Where(&Section{Slug: sectionSlug}).First(&section)
	return section
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
				"sum(segments_segment.square) as square_sum").
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

func GetAdminSections() []Section {
	var sections []Section
	DB.Preload("Company").Find(&sections)
	return sections
}

func GetAdminRacks() []Rack {
	var racks []Rack
	DB.Preload("Section").Find(&racks)
	return racks
}

func GetRack(companySlug, sectionSlug string, rackId int) Rack {
	var rack Rack
	DB.Joins("join segments_section on segments_rack.section_id = segments_section.id").
		Joins("join segments_company on segments_section.company_id = segments_company.id").
		Where("segments_company.slug = ? AND segments_section.slug = ? AND segments_rack.id = ?", companySlug, sectionSlug, uint(rackId)).
		First(&rack)
	return rack
}
