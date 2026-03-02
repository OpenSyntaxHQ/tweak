package tests

import (
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/OpenSyntaxHQ/tweak/processors"
)

func TestEpoch_UnixToDate(t *testing.T) {
	got, err := (processors.Epoch{}).Transform([]byte("0"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "1970") {
		t.Errorf("epoch 0 should contain 1970, got: %q", got)
	}
}

func TestEpoch_DateToUnix(t *testing.T) {
	got, err := (processors.Epoch{}).Transform([]byte("2023-01-01 00:00:00"))
	if err != nil {
		t.Fatal(err)
	}
	matched, _ := regexp.MatchString(`\d+`, got)
	if !matched {
		t.Errorf("expected unix timestamp in output, got: %q", got)
	}
}

func TestNow_Format(t *testing.T) {
	got, err := (processors.Now{}).Transform([]byte(""))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) == 0 {
		t.Error("now should produce non-empty output")
	}
}

func TestEpoch_MillisAndMicros(t *testing.T) {
	millis := "1704067200000"    // 2024-01-01 00:00:00 UTC
	micros := "1704067200000000" // 2024-01-01 00:00:00 UTC
	flags := []processors.Flag{
		{Name: "timezone", Short: "z", Value: "UTC", Type: processors.FlagString},
		{Name: "format", Short: "f", Value: "2006-01-02 15:04:05 MST", Type: processors.FlagString},
	}

	gotMillis, err := (processors.Epoch{}).Transform([]byte(millis), flags...)
	if err != nil {
		t.Fatalf("millis epoch conversion error: %v", err)
	}
	if !strings.Contains(gotMillis, "2024-01-01 00:00:00 UTC") {
		t.Fatalf("unexpected millis conversion output: %q", gotMillis)
	}

	gotMicros, err := (processors.Epoch{}).Transform([]byte(micros), flags...)
	if err != nil {
		t.Fatalf("micros epoch conversion error: %v", err)
	}
	if !strings.Contains(gotMicros, "2024-01-01 00:00:00 UTC") {
		t.Fatalf("unexpected micros conversion output: %q", gotMicros)
	}
}

func TestEpoch_InvalidInputAndTimezoneFallback(t *testing.T) {
	_, err := (processors.Epoch{}).Transform(
		[]byte("not-a-date"),
		processors.Flag{Name: "timezone", Short: "z", Value: "Invalid/Timezone", Type: processors.FlagString},
	)
	if err == nil {
		t.Fatal("expected parse error for invalid input")
	}

	got, err := (processors.Epoch{}).Transform(
		[]byte("0"),
		processors.Flag{Name: "timezone", Short: "z", Value: "Invalid/Timezone", Type: processors.FlagString},
		processors.Flag{Name: "format", Short: "f", Value: time.RFC3339, Type: processors.FlagString},
	)
	if err != nil {
		t.Fatalf("expected timezone fallback for invalid location, got error: %v", err)
	}
	if !strings.Contains(got, "1970-01-01T") {
		t.Fatalf("expected RFC3339 output for epoch 0, got %q", got)
	}
}

func TestNow_UTCFlag(t *testing.T) {
	got, err := (processors.Now{}).Transform(
		[]byte(""),
		processors.Flag{Name: "utc", Short: "u", Value: true, Type: processors.FlagBool},
		processors.Flag{Name: "format", Short: "f", Value: "2006-01-02 MST", Type: processors.FlagString},
	)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasSuffix(got, "UTC") {
		t.Fatalf("expected UTC suffix when utc flag set, got %q", got)
	}
}
