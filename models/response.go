package models

const (
	STATUS_OVER_QUERY_LIMIT = "OVER_QUERY_LIMIT"
	STATUS_OK               = "OK"
)

type PlaceTextSearchResponse struct {
	NextPageToken string
	Results       []*Result
	Status        string
}

type Result struct {
	FormattedaAddress string `json:"formatted_address"`
	Icon              string
	Name              string
}
