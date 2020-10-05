package syshelper

// SearchQuery struct
type SearchQuery struct {
	QueryString []interface{}
}

// AppendSearch function
func (sq *SearchQuery) AppendSearch(str string) {
	if str == "" {
		sq.QueryString = append(sq.QueryString, str)
	} else {
		sq.QueryString = append(sq.QueryString, "%"+str+"%")
	}
}
