package flag

import (
	"testing"

	"github.com/aquasecurity/trivy/pkg/telemetry"
	"github.com/stretchr/testify/assert"
)

var minimalPayload = telemetry.TelemetryPayload{Trivy: telemetry.TrivyInfo{Version: "dev"}, Usage: telemetry.UsageInfo{Flags: map[string]interface{}{}}}

func TestMakeTelemetryPayload(t *testing.T) {
	// TODO: add more test cases, how to instantiate Flags?
	tests := []struct {
		name    string
		flags   Flags
		options Options
		args    []string
		target  string
		want    telemetry.TelemetryPayload
		wantErr bool
	}{
		{
			name:    "empty",
			flags:   Flags{},
			options: Options{},
			args:    []string{},
			target:  "",
			want:    minimalPayload,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makeTelemetryPayload(tt.flags, tt.options, tt.target)
			// EnvironmentInfo is runtime dependant and excluded from test
			got.Environment = telemetry.EnvironmentInfo{}
			assert.Equal(t, tt.want, got)
		})
	}
}
