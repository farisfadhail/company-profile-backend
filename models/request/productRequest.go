package request

type ProductRequest struct {
	ProductCategoryId int    `json:"product_category_id" form:"product_category_id" validate:"required"`
	Title             string `json:"title" form:"title" validate:"required"`
	Slug              string `json:"slug" form:"slug"`
	Material          string `json:"material" form:"material"`
	Type              string `json:"type" form:"type"`
	Static            string `json:"static" form:"static"`
	Dynamic           string `json:"dynamic" form:"dynamic"`
	Racking           string `json:"racking" form:"racking"`
	TokopediaLink     string `json:"tokopedia_link" form:"tokopedia_link"`
	ShopeeLink        string `json:"shopee_link" form:"shopee_link"`
	LazadaLink        string `json:"lazada_link" form:"lazada_link"`
}