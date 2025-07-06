package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	ical "github.com/arran4/golang-ical"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	now := time.Now().UTC()

	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	endOfDay := startOfDay.Add(168 * time.Hour)

	// Build GraphQL query
	variables := map[string]any{
		"hl":             "en-US",
		"sport":          "val",
		"eventDateStart": startOfDay.Format("2006-01-02T15:04:05.000Z"),
		"eventDateEnd":   endOfDay.Format("2006-01-02T15:04:05.000Z"),
		"eventState":     []string{"unstarted"},
		"eventType":      "all",
		"pageSize":       1000,
	}
	variablesJson, _ := json.Marshal(variables)

	extensions := map[string]any{
		"persistedQuery": map[string]any{
			"version":    1,
			"sha256Hash": "7246add6f577cf30b304e651bf9e25fc6a41fe49aeafb0754c16b5778060fc0a",
		},
	}
	extensionsJson, _ := json.Marshal(extensions)

	queryUrl := fmt.Sprintf("https://valorantesports.com/api/gql?operationName=homeEvents&variables=%s&extensions=%s",
		urlQueryEscape(string(variablesJson)),
		urlQueryEscape(string(extensionsJson)),
	)

	req, _ := http.NewRequest("GET", queryUrl, nil)
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "en-US,en;q=0.5")
	req.Header.Add("apollographql-client-name", "Esports Web")
	req.Header.Add("apollographql-client-version", "8eecb20")
	req.Header.Add("content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var data ValorantEsportsResponse
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Failed to parse response:", err)
		return
	}

	// Group events by region
	grouped := make(map[string][]Events)
	for _, event := range data.Data.Esports.Events {
		region, ok := getRegionFromSlug(event.League.Slug)
		if !ok {
			continue
		}
		if len(event.MatchTeams) < 2 {
			continue
		}
		grouped[region] = append(grouped[region], event)
	}

	// List of all regions to always create ICS files for
	allRegions := []string{"emea", "americas", "pacific", "china", "international"}

	// Create ICS for each region, even if empty
	for _, region := range allRegions {
		events := grouped[region]
		err := writeICS(region, events)
		if err != nil {
			fmt.Printf("Failed to write ICS for region %s: %v\n", region, err)
		} else {
			fmt.Printf("Created valorant_%s.ics with %d matches\n", region, len(events))
		}
	}

	err2 := generateIndexHTML("output")
	if err2 != nil {
		fmt.Println("Failed to write index.html:", err2)
	}
}

func getRegionFromSlug(slug string) (string, bool) {
	if slug == "vct_masters" || slug == "champions" {
		return "international", true
	}
	if strings.HasPrefix(slug, "vct_") {
		return strings.TrimPrefix(slug, "vct_"), true
	}
	if strings.HasPrefix(slug, "last_chance_qualifier_") {
		return strings.TrimPrefix(slug, "last_chance_qualifier_"), true
	}
	return "", false
}

func writeICS(region string, events []Events) error {
	cal := ical.NewCalendar()
	cal.SetMethod(ical.MethodRequest)
	titleCaser := cases.Title(language.Und)
	cal.SetProductId(fmt.Sprintf("-//Valorant Esports Calendar - %s//EN", titleCaser.String(region)))

	for _, event := range events {
		team1 := event.MatchTeams[0].Name
		team2 := event.MatchTeams[1].Name
		bestOf := event.Match.Strategy.Count
		startTime := event.StartTime

		title := fmt.Sprintf("%s Vs %s (BO%d)", team1, team2, bestOf)
		description := fmt.Sprintf("League: %s", event.League.Slug)

		duration := 2 * time.Hour
		if bestOf == 5 {
			duration = 4 * time.Hour
		}
		endTime := startTime.Add(duration)

		eventId := fmt.Sprintf("%s-%s-%s", team1, team2, startTime.Format("20060102150405"))

		evt := cal.AddEvent(eventId)
		evt.SetSummary(title)
		evt.SetDescription(description)
		evt.SetDtStampTime(time.Now().UTC())
		evt.SetStartAt(startTime.UTC())
		evt.SetEndAt(endTime.UTC())
		evt.SetCreatedTime(time.Now().UTC())
		evt.SetModifiedAt(time.Now().UTC())
	}

	fileName := fmt.Sprintf("valorant_%s.ics", region)
	f, err := os.Create(fmt.Sprintf("output/%s", fileName))
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write([]byte(cal.Serialize()))
	return err
}

func urlQueryEscape(s string) string {
	return (&url.URL{Path: s}).EscapedPath()
}

func generateIndexHTML(outputDir string) error {
	indexPath := filepath.Join(outputDir, "index.html")
	f, err := os.Create(indexPath)
	if err != nil {
		return err
	}
	defer f.Close()

	// Start HTML
	f.WriteString("<!DOCTYPE html>\n<html>\n<head><meta charset=\"UTF-8\"><title>Calendars</title></head><body>\n")
	f.WriteString("<h1>Available Calendars</h1>\n<ul>\n")

	// List all .ics files in outputDir
	err = filepath.WalkDir(outputDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && filepath.Ext(path) == ".ics" {
			relPath, _ := filepath.Rel(outputDir, path)
			f.WriteString(fmt.Sprintf("<li><a href=\"%s\">%s</a></li>\n", relPath, relPath))
		}
		return nil
	})

	if err != nil {
		return err
	}

	// End HTML
	f.WriteString("</ul>\n</body>\n</html>\n")
	return nil
}
