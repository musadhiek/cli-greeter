package main

import (
	"flag"
	"os"
	"testing"
)

func TestGreet(t *testing.T) {
	tests := []struct {
		name      string
		flagValue string
		envValue  string
		want      string
		wantErr   bool
	}{
		{
			name:      "default value when no flag or env set",
			flagValue: "",
			envValue:  "",
			want:      "World",
			wantErr:   false,
		},
		{
			name:      "env variable overrides default",
			flagValue: "",
			envValue:  "Bob",
			want:      "Bob",
			wantErr:   false,
		},
		{
			name:      "flag overrides default",
			flagValue: "Alice",
			envValue:  "",
			want:      "Alice",
			wantErr:   false,
		},
		{
			name:      "flag takes precedence over env variable",
			flagValue: "Charlie",
			envValue:  "Bob",
			want:      "Charlie",
			wantErr:   false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup test environment
			args := []string{}
			if tc.flagValue != "" {
				args = []string{"-name", tc.flagValue}
			}

			if tc.envValue != "" {
				os.Setenv("GREET_NAME", tc.envValue)
			} else {
				os.Unsetenv("GREET_NAME")
			}

			// Clean up environment after test
			defer os.Unsetenv("GREET_NAME")

			// Execute test
			fs := flag.NewFlagSet("test", flag.ContinueOnError)
			got, err := greet(fs, args)

			// Verify results
			if (err != nil) != tc.wantErr {
				t.Errorf("greet() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if got != tc.want {
				t.Errorf("greet() = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestGreet_InvalidFlag(t *testing.T) {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	_, err := greet(fs, []string{"-invalid-flag"})
	if err == nil {
		t.Error("expected error for invalid flag, got nil")
	}
}
