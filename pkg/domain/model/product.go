package model

type ProductCreateReq struct {
	Name             string                 `json:"name" validate:"required"`
	Description      *string                `json:"description,omitempty"`
	ShortDescription *string                `json:"short_description,omitempty"`
	Sku              string                 `json:"sku" validate:"required"`
	Price            int                    `json:"price" validate:"required"`
	OriginalPrice    *int                   `json:"original_price,omitempty"`
	CostPrice        *int                   `json:"cost_price,omitempty"`
	Stock            int                    `json:"stock" validate:"required"`
	MinStock         *int                   `json:"min_stock,omitempty"`
	CategoryID       *int                   `json:"category_id,omitempty"`
	Brand            *string                `json:"brand,omitempty"`
	Unit             *string                `json:"unit,omitempty"`
	Weight           *float64               `json:"weight,omitempty"`
	Volume           *float64               `json:"volume,omitempty"`
	Images           []string               `json:"images,omitempty"`
	Attributes       map[string]interface{} `json:"attributes,omitempty"`
	Tags             []string               `json:"tags,omitempty"`
	Active           *bool                  `json:"active,omitempty"`
	Featured         *bool                  `json:"featured,omitempty"`
	Digital          *bool                  `json:"digital,omitempty"`
	MetaTitle        *string                `json:"meta_title,omitempty"`
	MetaDescription  *string                `json:"meta_description,omitempty"`
	MetaKeywords     *string                `json:"meta_keywords,omitempty"`
	SortOrder        *int                   `json:"sort_order,omitempty"`
}

type ProductBatchUpdateReq struct {
	IDs    []int  `json:"ids" validate:"required"`
	Active *bool  `json:"active,omitempty"`
}

type ProductBatchDeleteReq struct {
	IDs []int `json:"ids" validate:"required"`
}

type ProductUpdateReq struct {
	Name             *string                `json:"name,omitempty"`
	Description      *string                `json:"description,omitempty"`
	ShortDescription *string                `json:"short_description,omitempty"`
	Sku              *string                `json:"sku,omitempty"`
	Price            *int                   `json:"price,omitempty"`
	OriginalPrice    *int                   `json:"original_price,omitempty"`
	CostPrice        *int                   `json:"cost_price,omitempty"`
	Stock            *int                   `json:"stock,omitempty"`
	MinStock         *int                   `json:"min_stock,omitempty"`
	Sales            *int                   `json:"sales,omitempty"`
	CategoryID       *int                   `json:"category_id,omitempty"`
	Brand            *string                `json:"brand,omitempty"`
	Unit             *string                `json:"unit,omitempty"`
	Weight           *float64               `json:"weight,omitempty"`
	Volume           *float64               `json:"volume,omitempty"`
	Images           []string               `json:"images,omitempty"`
	Attributes       map[string]interface{} `json:"attributes,omitempty"`
	Tags             []string               `json:"tags,omitempty"`
	Active           *bool                  `json:"active,omitempty"`
	Featured         *bool                  `json:"featured,omitempty"`
	Digital          *bool                  `json:"digital,omitempty"`
	MetaTitle        *string                `json:"meta_title,omitempty"`
	MetaDescription  *string                `json:"meta_description,omitempty"`
	MetaKeywords     *string                `json:"meta_keywords,omitempty"`
	SortOrder        *int                   `json:"sort_order,omitempty"`
}
