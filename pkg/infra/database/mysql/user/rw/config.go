package ordersDataStore

type MySqlConfig struct {
	Database             string `yaml:"database"`
	Host                 string `yaml:"host"`
	Port                 int    `yaml:"port"`
	User                 string `yaml:"user"`
	Password             string `yaml:"password"`
	MaxIdleConnection    int    `yaml:"max_idle_connections"`
	MaxOpenConnection    int    `yaml:"max_open_connections"`
	ConnMaxLifetimeInSec int    `yaml:"connection_max_lifetime_sec"`
	ConnMaxIdleTimeInSec int    `yaml:"connection_max_idle_time_sec"`
}

func (m *MySqlConfig) SetupDefault() {
	if m.Host == "" {
		m.Host = "localhost"
	}
	if m.Port <= 0 {
		m.Port = 3306
	}
	if m.MaxIdleConnection <= 0 {
		m.MaxIdleConnection = 10
	}
	if m.MaxOpenConnection <= 0 {
		m.MaxOpenConnection = 10
	}
	if m.ConnMaxLifetimeInSec <= 0 {
		m.ConnMaxLifetimeInSec = 60
	}
	if m.ConnMaxIdleTimeInSec <= 0 {
		m.ConnMaxIdleTimeInSec = 60
	}
}

// GetDatabase returns the database name
func (m *MySqlConfig) GetDatabase() string {
	return m.Database
}

// GetHost returns the database host
func (m *MySqlConfig) GetHost() string {
	return m.Host
}

// GetPort returns the database port
func (m *MySqlConfig) GetPort() int {
	return m.Port
}

// GetUser returns the database user
func (m *MySqlConfig) GetUser() string {
	return m.User
}

// GetPassword returns the database password
func (m *MySqlConfig) GetPassword() string {
	return m.Password
}

// GetMaxIdleConnection returns the maximum idle connections
func (m *MySqlConfig) GetMaxIdleConnection() int {
	return m.MaxIdleConnection
}

// GetMaxOpenConnection returns the maximum open connections
func (m *MySqlConfig) GetMaxOpenConnection() int {
	return m.MaxOpenConnection
}

// GetConnMaxLifetimeInSec returns the connection max lifetime in seconds
func (m *MySqlConfig) GetConnMaxLifetimeInSec() int {
	return m.ConnMaxLifetimeInSec
}

// GetConnMaxIdleTimeInSec returns the connection max idle time in seconds
func (m *MySqlConfig) GetConnMaxIdleTimeInSec() int {
	return m.ConnMaxIdleTimeInSec
}
