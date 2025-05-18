package dto 

import "time"

type DoctorListResponse struct {
	Id        int       `json:"id"`
	Uid       string    `json:"uid,omitempty"`
	NameEn    string    `json:"name_en,omitempty"`
	NameBn    string    `json:"name_bn,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	Email     string    `json:"email,omitempty"`
	Address   string    `json:"address,omitempty"`
	Designation string    `json:"designation,omitempty"`
	DepartmentId int       `json:"department_id,omitempty"`
	DepartmentName string    `json:"department_name,omitempty"`
	BranchId  int       `json:"branch_id,omitempty"`
	BranchName string    `json:"branch_name,omitempty"`
	Price     float64   `json:"price,omitempty"`
	Status    int8      `json:"status"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Meta      *string   `json:"meta,omitempty"`
	Remarks   *string   `json:"remarks,omitempty"`
	// Additional fields for pagination
	TotalCount int `json:"total_count,omitempty"`
	CurrentPage int `json:"current_page,omitempty"`
	TotalPages int `json:"total_pages,omitempty"`
	PerPage int `json:"per_page,omitempty"`
	// Additional fields for sorting
	SortBy string `json:"sort_by,omitempty"`
	SortOrder string `json:"sort_order,omitempty"`
	// Additional fields for filtering
	FilterBy string `json:"filter_by,omitempty"`
	FilterValue string `json:"filter_value,omitempty"`
	// Additional fields for searching
	SearchQuery string `json:"search_query,omitempty"`
	// Additional fields for grouping
	GroupBy string `json:"group_by,omitempty"`
	// Additional fields for aggregation
	AggregateBy string `json:"aggregate_by,omitempty"`
	// Additional fields for custom metadata
	CustomMeta map[string]interface{} `json:"custom_meta,omitempty"`
	// Additional fields for related entities
	RelatedEntities []string `json:"related_entities,omitempty"`
	// Additional fields for localization
	Locale string `json:"locale,omitempty"`
	// Additional fields for versioning					
}