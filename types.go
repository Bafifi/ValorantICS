package main

import "time"

type ValorantEsportsResponse struct {
	Data Data `json:"data"`
}
type DisplayPriority struct {
	Typename string `json:"__typename"`
	Position int    `json:"position"`
	Status   string `json:"status"`
}
type League struct {
	Typename        string          `json:"__typename"`
	DisplayPriority DisplayPriority `json:"displayPriority"`
	ID              string          `json:"id"`
	Image           string          `json:"image"`
	Name            string          `json:"name"`
	Slug            string          `json:"slug"`
}
type Games struct {
	Typename string `json:"__typename"`
	ID       string `json:"id"`
	Number   int    `json:"number"`
	State    string `json:"state"`
	Vods     []any  `json:"vods"`
	Recaps   []any  `json:"recaps"`
}
type Strategy struct {
	Typename string `json:"__typename"`
	Count    int    `json:"count"`
	Type     string `json:"type"`
}
type Match struct {
	Typename string   `json:"__typename"`
	Flags    []any    `json:"flags"`
	Games    []Games  `json:"games"`
	ID       string   `json:"id"`
	State    string   `json:"state"`
	Strategy Strategy `json:"strategy"`
	Type     string   `json:"type"`
}
type Result struct {
	Typename string `json:"__typename"`
	GameWins int    `json:"gameWins"`
	Outcome  any    `json:"outcome"`
}
type MatchTeams struct {
	Typename   string `json:"__typename"`
	Code       string `json:"code"`
	ID         string `json:"id"`
	Image      string `json:"image"`
	LightImage any    `json:"lightImage"`
	Name       string `json:"name"`
	Result     Result `json:"result"`
}
type Tournament struct {
	Typename string `json:"__typename"`
	ID       string `json:"id"`
	Name     string `json:"name"`
}
type Events struct {
	Typename   string       `json:"__typename"`
	BlockName  string       `json:"blockName"`
	ID         string       `json:"id"`
	League     League       `json:"league"`
	Match      Match        `json:"match"`
	MatchTeams []MatchTeams `json:"matchTeams"`
	StartTime  time.Time    `json:"startTime"`
	State      string       `json:"state"`
	Streams    []any        `json:"streams"`
	Tournament Tournament   `json:"tournament"`
	Type       string       `json:"type"`
}
type Pages struct {
	Typename string `json:"__typename"`
	Newer    string `json:"newer"`
	Older    any    `json:"older"`
}
type Esports struct {
	Typename string   `json:"__typename"`
	Events   []Events `json:"events"`
	Pages    Pages    `json:"pages"`
}
type Data struct {
	Typename string  `json:"__typename"`
	Esports  Esports `json:"esports"`
}

type EventSummary struct {
	TeamNames  []string
	BestOf     int
	LeagueSlug string
	StartTime  time.Time
}
