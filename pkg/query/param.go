package query

type Param struct {
	SortBy string
	Sort   string
	Limit  uint
	Page   uint
}

func DefaultParam() *Param {
	return &Param{
		SortBy: "created_at",
		Sort:   "asc",
		Limit:  10,
		Page:   1,
	}
}
