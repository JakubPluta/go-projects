package newsapi

type Options struct {
	Country  string // ISO 3166-1 code of country e.g pl,de, us
	Category string // business entertainment general health science sports technology
	Domains  string //  A comma-separated string of domains
	From     string //  date from ISO 8601 format
	To       string // date to ISO 8601 format
	Language string // ar de en es fr he it nl no pt ru se ud zh
	SortBy   string // sort order: relevancy, popularity, publishedAt.
	Sources  string // news sources
	PageSize int    // number of results per page
	Page     int    // page number
	Q        string // query phrase
}
