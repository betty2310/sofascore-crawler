package types

type StatisticsItem struct {
	Name string `json:"name"`
	Home string `json:"home"`
	Away string `json:"away"`
}

type Group struct {
	GroupName       string           `json:"groupName"`
	StatisticsItems []StatisticsItem `json:"statisticsItems"`
}

type PeriodStatistics struct {
	Period string  `json:"period"`
	Groups []Group `json:"groups"`
}

type StatisticsData struct {
	Statistics []PeriodStatistics `json:"statistics"`
}
