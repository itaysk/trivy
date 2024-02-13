package flag

import (
	"runtime"
	"strings"

	"github.com/aquasecurity/trivy/pkg/telemetry"
	"github.com/aquasecurity/trivy/pkg/version"
)

func makeTelemetryPayload(flags Flags, options Options, target string) telemetry.TelemetryPayload {
	usageflags := make(map[string]interface{})
	for _, group := range flags.Groups() {
		for _, f := range group.Flags() {
			if !f.IsSet() {
				continue
			}
			key := f.GetConfigName()
			value := f.GetValue()

			//workaround for https://github.com/spf13/viper/discussions/1766
			if f.GetName() == "" && strings.HasPrefix(key, "license.") {
				continue
			}

			switch f.(type) {
			case *Flag[[]string]:
				if len(f.GetValues()) == 0 {
					value = "***"
				}
			case *Flag[string]:
				if len(f.GetValues()) == 0 {
					value = "***"
				}
			}

			usageflags[key] = value
		}
	}
	usagePayload := telemetry.UsageInfo{
		Target: target,
		Flags:  usageflags,
	}

	v := version.NewVersionInfo(options.CacheDir)
	trivyPayload := telemetry.TrivyInfo{
		Version: v.Version,
	}
	if v.VulnerabilityDB != nil {
		trivyPayload.VulnDBTime = v.VulnerabilityDB.DownloadedAt
	}
	if v.PolicyBundle != nil {
		trivyPayload.PolicyDBTime = v.PolicyBundle.DownloadedAt
	}
	if v.JavaDB != nil {
		trivyPayload.JavaDBTime = v.JavaDB.DownloadedAt
	}

	envPayload := telemetry.EnvironmentInfo{
		Os:      runtime.GOOS,
		CpuArch: runtime.GOARCH,
	}

	payload := telemetry.TelemetryPayload{
		Trivy:       trivyPayload,
		Usage:       usagePayload,
		Environment: envPayload,
	}

	return payload
}
