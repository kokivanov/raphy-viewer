package core

type GQLRequest struct {
	Query     string      `json:"query"`
	Variables interface{} `json:"variables,omitempty"`
}

type Error struct {
	Message   string     `json:"message"`
	Status    int64      `json:"status"`
	Locations []Location `json:"locations"`
}

type Location struct {
	Line   int64 `json:"line"`
	Column int64 `json:"column"`
}

type ShortResponse struct {
	Errors []Error   `json:"errors"`
	Data   ShortData `json:"data"`
}

type ShortData struct {
	Page ShortPage `json:"Page"`
}

type ShortPage struct {
	Media []ShortMedia `json:"media"`
}

type ShortMedia struct {
	IDMal      int64      `json:"idMal"`
	Title      Title      `json:"title"`
	Status     Status     `json:"status"`
	IsAdult    bool       `json:"isAdult"`
	CoverImage CoverImage `json:"coverImage"`
	Rankings   []Ranking  `json:"rankings"`
}

type CoverImage struct {
	Large      string `json:"large"`
	ExtraLarge string `json:"extraLarge"`
	Medium     string `json:"medium"`
	Color      string `json:"color"`
}

type Title struct {
	English       string `json:"english"`
	UserPreferred string `json:"userPreferred"`
	Native        string `json:"native"`
}

type Status string

type FullDataResponse struct {
	Errors []Error  `json:"errors"`
	Data   FullData `json:"data"`
}

type FullData struct {
	Media FullMedia `json:"Media"`
}

type FullMedia struct {
	IDMal             int64           `json:"idMal"`
	Title             MediaTitle      `json:"title"`
	SiteURL           string          `json:"siteUrl"`
	Reviews           Reviews         `json:"reviews"`
	NextAiringEpisode interface{}     `json:"nextAiringEpisode"`
	IsAdult           bool            `json:"isAdult"`
	Studios           Studios         `json:"studios"`
	Popularity        int64           `json:"popularity"`
	AverageScore      int64           `json:"averageScore"`
	Genres            []string        `json:"genres"`
	Synonyms          []string        `json:"synonyms"`
	BannerImage       string          `json:"bannerImage"`
	CoverImage        MediaCoverImage `json:"coverImage"`
	Rankings          []Ranking       `json:"rankings"`
	Tags              []Tag           `json:"tags"`
	Trailer           Trailer         `json:"trailer"`
	Hashtag           interface{}     `json:"hashtag"`
	IsLicensed        bool            `json:"isLicensed"`
	CountryOfOrigin   string          `json:"countryOfOrigin"`
	Episodes          int64           `json:"episodes"`
	SeasonYear        int64           `json:"seasonYear"`
	Season            string          `json:"season"`
	StartDate         Date            `json:"startDate"`
	Format            string          `json:"format"`
	Status            string          `json:"status"`
	Description       string          `json:"description"`
	Trending          int64           `json:"trending"`
	Relations         Relations       `json:"relations"`
	EndDate           Date            `json:"endDate"`
	Recommendations   Recommendations `json:"recommendations"`
}

type MediaCoverImage struct {
	ExtraLarge string `json:"extraLarge"`
	Large      string `json:"large"`
	Color      string `json:"color"`
}

type Date struct {
	Year  int64 `json:"year"`
	Month int64 `json:"month"`
	Day   int64 `json:"day"`
}

type Ranking struct {
	ID      int64       `json:"id"`
	Rank    int64       `json:"rank"`
	Type    string      `json:"type"`
	Format  string      `json:"format"`
	Year    *int64      `json:"year"`
	Season  interface{} `json:"season"`
	AllTime bool        `json:"allTime"`
	Context string      `json:"context"`
}

type Recommendations struct {
	PageInfo PageInfo              `json:"pageInfo"`
	Edges    []RecommendationsEdge `json:"edges"`
}

type RecommendationsEdge struct {
	Node PurpleNode `json:"node"`
}

type PurpleNode struct {
	MediaRecommendation MediaRecommendationClass `json:"mediaRecommendation"`
}

type MediaRecommendationClass struct {
	IDMal      int64          `json:"idMal"`
	Title      NodeTitle      `json:"title"`
	CoverImage NodeCoverImage `json:"coverImage"`
}

type NodeCoverImage struct {
	Color  *string `json:"color"`
	Medium string  `json:"medium"`
}

type NodeTitle struct {
	English       *string `json:"english"`
	UserPreferred string  `json:"userPreferred"`
}

type PageInfo struct {
	Total       int64 `json:"total"`
	HasNextPage bool  `json:"hasNextPage"`
}

type Relations struct {
	Edges []RelationsEdge `json:"edges"`
}

type RelationsEdge struct {
	RelationType string                   `json:"relationType"`
	Node         MediaRecommendationClass `json:"node"`
}

type Reviews struct {
	PageInfo PageInfo `json:"pageInfo"`
}

type Studios struct {
	Edges []StudiosEdge `json:"edges"`
}

type StudiosEdge struct {
	Node FluffyNode `json:"node"`
}

type FluffyNode struct {
	Name              string `json:"name"`
	SiteURL           string `json:"siteUrl"`
	IsAnimationStudio bool   `json:"isAnimationStudio"`
}

type Tag struct {
	Name    string `json:"name"`
	IsAdult bool   `json:"isAdult"`
}

type MediaTitle struct {
	Romaji        string `json:"romaji"`
	English       string `json:"english"`
	Native        string `json:"native"`
	UserPreferred string `json:"userPreferred"`
}

type Trailer struct {
	Site      string `json:"site"`
	Thumbnail string `json:"thumbnail"`
}
