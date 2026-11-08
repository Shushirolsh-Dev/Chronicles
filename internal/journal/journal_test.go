package journal

import (
	"testing"
	"time"
)

func TestParseLog(t *testing.T) {
	raw := "abc123\x1fAda Lovelace\x1fada@example.com\x1f2026-11-08T09:30:00+01:00\x1fAdd journal output\x1fA longer description\n\x1e"
	entries, err := ParseLog(raw)
	if err != nil { t.Fatalf("ParseLog() error = %v", err) }
	if len(entries) != 1 { t.Fatalf("got %d entries, want 1", len(entries)) }
	if entries[0].Subject != "Add journal output" { t.Errorf("subject = %q", entries[0].Subject) }
	_, offset := entries[0].Date.Zone()
	if offset != 3600 { t.Errorf("date offset = %d, want 3600", offset) }
}

func TestParseEmptyLog(t *testing.T) {
	entries, err := ParseLog("")
	if err != nil { t.Fatalf("ParseLog() error = %v", err) }
	if len(entries) != 0 { t.Fatalf("got %d entries, want 0", len(entries)) }
}

func TestGroupByDateUsesRequestedLocation(t *testing.T) {
	entries := []Entry{{Date: time.Date(2026, 11, 8, 23, 30, 0, 0, time.UTC)}}
	location := time.FixedZone("WAT", 3600)
	groups := GroupByDate(entries, location)
	if len(groups["2026-11-09"]) != 1 { t.Fatalf("expected commit grouped under local date 2026-11-09") }
}
