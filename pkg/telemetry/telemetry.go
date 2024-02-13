package telemetry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aquasecurity/trivy/pkg/log"
)

type TrivyInfo struct {
	Version      string    `json:"version ,omitempty"`
	VulnDBTime   time.Time `json:"vulnDB ,omitempty"`
	PolicyDBTime time.Time `json:"policyDB ,omitempty"`
	JavaDBTime   time.Time `json:"javaDB ,omitempty"`
}

type UsageInfo struct {
	Target string                 `json:"target ,omitempty"`
	Flags  map[string]interface{} `json:"flags ,omitempty"`
}

type EnvironmentInfo struct {
	Os      string `json:"os ,omitempty"`
	CpuArch string `json:"cpuArch ,omitempty"`
}

type TelemetryPayload struct {
	Trivy       TrivyInfo       `json:"trivy,omitempty"`
	Usage       UsageInfo       `json:"usage,omitempty"`
	Environment EnvironmentInfo `json:"environment,omitempty"`
}

// SendTelemetry sends the telemetry payload to the server
func SendTelemetry(payload TelemetryPayload, url string) error {
	//TODO: url validation?
	log.Logger.Debugf("sending telemetry to %s", url)
	log.Logger.Debugf("telemetry payload %v", payload)
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send telemetry: status code %d", resp.StatusCode)
	}

	log.Logger.Debugf("telemetry sent successfully")
	return nil
}
