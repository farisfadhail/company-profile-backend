package request

type ProductRequest struct {
	ProductCategoryId int    `json:"product_category_id" validate:"required"`
	Title             string `json:"title" validate:"required"`
	Slug              string `json:"slug"`
	Material          string `json:"material"`
	Type              string `json:"type"`
	Static            string `json:"static"`
	Dynamic           string `json:"dynamic"`
	Racking           string `json:"racking"`
	TokopediaLink     string `json:"tokopedia_link"`
	ShopeeLink        string `json:"shopee_link"`
	LazadaLink        string `json:"lazada_link"`
}