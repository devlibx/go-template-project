package command

import (
	goxBaseConfig "github.com/devlibx/gox-base/v2/config"
	goxBaseMetrics "github.com/devlibx/gox-base/v2/metrics"
	goxHttpApi "github.com/devlibx/gox-http/v3/api"
	goxHttp "github.com/devlibx/gox-http/v3/command"
	goxMessaging "github.com/devlibx/gox-messaging/v2"
)

type ApplicationConfig struct {
	App                           *goxBaseConfig.App                        `yaml:"app"`
	MetricConfig                  *goxBaseMetrics.Config                    `yaml:"metric"`
	HttpConfig                    *goxHttp.Config                           `yaml:"server_config"`
	MessagingConfig               *goxMessaging.Configuration               `yaml:"messaging_config"`
	RequestResponseSecurityConfig *goxHttpApi.RequestResponseSecurityConfig `yaml:"gox_http_request_response_security_config"`
}
