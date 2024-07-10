package command

import (
	goxBaseConfig "github.com/devlibx/gox-base/config"
	goxBaseMetrics "github.com/devlibx/gox-base/metrics"
	goxHttpApi "github.com/devlibx/gox-http/v2/api"
	goxHttp "github.com/devlibx/gox-http/v2/command"
	goxMessaging "github.com/devlibx/gox-messaging"
)

type ApplicationConfig struct {
	App                           *goxBaseConfig.App                        `yaml:"app"`
	MetricConfig                  *goxBaseMetrics.Config                    `yaml:"metric"`
	HttpConfig                    *goxHttp.Config                           `yaml:"server_config"`
	MessagingConfig               *goxMessaging.Configuration               `yaml:"messaging_config"`
	RequestResponseSecurityConfig *goxHttpApi.RequestResponseSecurityConfig `yaml:"gox_http_request_response_security_config"`
}
