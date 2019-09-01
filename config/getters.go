package config

import (
	"fmt"
	"strings"
	"time"
	"github.com/wallix/awless/global"
)

func GetAWSRegion() string {
	if reg, ok := Config[RegionConfigKey]; ok && reg != "" {
		return fmt.Sprint(reg)
	}
	if reg, ok := Defaults["region"]; ok && reg != "" { // Compatibility with old key
		return fmt.Sprint(reg)
	}
	return ""
}

func GetUsername() string {
	if reg, ok := Config["user.username"]; ok && reg != "" {
		return fmt.Sprint(reg)
	}
	return ""
}

func GetUserId() string {
	if reg, ok := Config["user.id"]; ok && reg != "" {
		return fmt.Sprint(reg)
	}
	return ""
}

func GetToken() string {
	if reg, ok := Config["_token"]; ok && reg != "" {
		return fmt.Sprint(reg)
	}
	return ""
}

func GetPassword() string {
	if reg, ok := Config["_enc"]; ok && reg != "" {
		return fmt.Sprint(reg)
	}
	return ""
}

func GetPasswordPlainText() string {
	enc := []byte(GetPassword())
        pass := string(Decrypt(enc, global.ENC_PWD))
	return pass
}

const defaultAWSSessionProfile = "default"

func GetAWSProfile() string {
	if profile, ok := Config[ProfileConfigKey]; ok && profile != "" {
		return fmt.Sprint(profile)
	}
	if profile, ok := Defaults[ProfileConfigKey]; ok && profile != "" { // Compatibility with old key
		return fmt.Sprint(profile)
	}
	return defaultAWSSessionProfile
}

func GetAutosync() bool {
	return false
}

func GetSchedulerURL() string {
	/*
		if u, ok := Config[schedulerURL].(string); ok {
			return u
		}
	*/
	return ""
}

func GetConfigWithPrefix(prefix string) map[string]interface{} {
	conf := make(map[string]interface{})
	for k, v := range Config {
		if strings.HasPrefix(k, prefix) {
			conf[k] = v
		}
	}
	return conf
}

func getCheckUpgradeFrequency() time.Duration {
	if frequency, ok := Config[checkUpgradeFrequencyConfigKey].(int); ok {
		return time.Duration(frequency) * time.Hour
	}
	return 8 * time.Hour
}
