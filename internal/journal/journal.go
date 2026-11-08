package journal

import (
	"fmt"
	"strings"
	"time"
)

type Entry struct {
	Hash string
	Author string
	Email string
	Date time.Time
	Subject string
	Body string
}

func ParseLog(output string) ([]Entry, error) {
	if strings.TrimSpace(output) == "" { return []Entry{}, nil }
	records := strings.Split(output, "\x1e")
	entries := make([]Entry, 0, len(records))
	for _, record := range records {
		record = strings.TrimSpace(record)
		if record == "" { continue }
		fields := strings.SplitN(record, "\x1f", 6)
		if len(fields) != 6 { return nil, fmt.Errorf("unexpected Git log record with %d fields", len(fields)) }
		date, err := time.Parse(time.RFC3339, strings.TrimSpace(fields[3]))
		if err != nil { return nil, fmt.Errorf("parse commit date %q: %w", fields[3], err) }
		subject := strings.TrimSpace(fields[4])
		if subject == "" { subject = "(no commit message)" }
		entries = append(entries, Entry{
			Hash: strings.TrimSpace(fields[0]), Author: strings.TrimSpace(fields[1]),
			Email: strings.TrimSpace(fields[2]), Date: date, Subject: subject,
			Body: strings.TrimSpace(fields[5]),
		})
	}
	return entries, nil
}

func GroupByDate(entries []Entry, location *time.Location) map[string][]Entry {
	groups := make(map[string][]Entry)
	for _, entry := range entries {
		date := entry.Date.In(location).Format("2006-01-02")
		groups[date] = append(groups[date], entry)
	}
	return groups
}
