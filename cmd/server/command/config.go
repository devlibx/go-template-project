package command

import (
	goxBaseConfig "github.com/devlibx/gox-base/v2/config"
	goxBaseMetrics "github.com/devlibx/gox-base/v2/metrics"
	goxHttpApi "github.com/devlibx/gox-http/v4/api"
	goxHttp "github.com/devlibx/gox-http/v4/command"
	goxMessaging "github.com/devlibx/gox-messaging/v2"
	cadenceConfig "github.com/devlibx/gox-workfkow/workflow/framework/cadence"
)

type ApplicationConfig struct {
	App                           *goxBaseConfig.App                        `yaml:"app"`
	MetricConfig                  *goxBaseMetrics.Config                    `yaml:"metric"`
	HttpConfig                    *goxHttp.Config                           `yaml:"server_config"`
	MessagingConfig               *goxMessaging.Configuration               `yaml:"messaging_config"`
	RequestResponseSecurityConfig *goxHttpApi.RequestResponseSecurityConfig `yaml:"gox_http_request_response_security_config"`
	CadenceConfig                 *cadenceConfig.Config                     `yaml:"cadence_config"`
}

func (a *ApplicationConfig) SetDefaults() {
	if a.CadenceConfig == nil {
		a.CadenceConfig = &cadenceConfig.Config{Disabled: true}
	}
}
