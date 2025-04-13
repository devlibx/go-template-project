# Go Template Project

[![Go Reference](https://pkg.go.dev/badge/github.com/harishbohara/go-template-project.svg)](https://pkg.go.dev/github.com/harishbohara/go-template-project)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
![Go Version](https://img.shields.io/github/go-mod/go-version/harishbohara/go-template-project?filename=go.mod)

A production-ready Go server template project with a clean architecture, configuration management, and end-to-end testing setup.

## 🚀 Quick Start

```bash
# Clone the repository
git clone https://github.com/devlibx/go-template-project.git

# Navigate to the project
cd go-template-project

# Install dependencies
go mod tidy

# Run the server in development mode
sh build/run-local-dev.sh
```

## ✨ Features

- 🏗️ Clean Architecture
- ⚙️ Environment-based Configuration
- 🔒 Secure Defaults
- 🧪 E2E Testing Setup
- 📝 Structured Logging
- 🛠️ Development Tools
- 🔄 CI/CD Ready
- 📊 Metrics Integration
- 🔍 Performance Profiling
- 📨 Message Queue Integration

## 📋 Prerequisites

- Go 1.19 or higher
- Git
- Kafka (for messaging features)

## 📦 Installation

```bash
# Clone the repository
git clone https://github.com/devlibx/go-template-project.git

# Navigate to the project directory
cd go-template-project

# Install dependencies
go mod tidy
```

## ⚙️ Configuration

The project uses YAML configuration files located in the `config/` directory:

- `app.yaml`: Application-specific configuration
- `http.yaml`: HTTP server settings
- `messaging.yaml`: Messaging service configuration

### Using HTTP APIs

Define your HTTP APIs in `http.yaml`:

```yaml
# http.yaml
server_config:
  servers:
    jsonplaceholder:
      host: jsonplaceholder.typicode.com
      port: -1
      https: true
  apis:
    getPosts:
      method: GET
      path: /todos/{postId}
      server: jsonplaceholder
      timeout: 100000
```

Implement the HTTP client:

```go
// Client interface
type Client interface {
    GetPosts(ctx context.Context, postId string) (*PostDto, error)
}

// DTO for response
type PostDto struct {
    Id     int    `json:"id"`
    UserId int    `json:"userId"`
    Title  string `json:"title"`
}

// Client implementation
type client struct {
    gox.CrossFunction
    goxHttpApi.GoxHttpContext
}

// Create new client instance
func NewClient(cf gox.CrossFunction, goxHttpCtx goxHttpApi.GoxHttpContext) Client {
    return &client{
        CrossFunction:  cf,
        GoxHttpContext: goxHttpCtx,
    }
}

// Implementation of GetPosts method
func (c *client) GetPosts(ctx context.Context, postId string) (*PostDto, error) {
    type responseObj struct {
        Id     int    `json:"id"`
        UserId int    `json:"userId"`
        Title  string `json:"title"`
    }

    httpRequest := command.NewGoxRequestBuilder("getPosts").
        WithContentTypeJson().
        WithPathParam("postId", postId).
        Build()
    
    httpResponse, err := goxHttpApi.ExecuteHttp[responseObj, any](ctx, c, httpRequest)
    if err != nil {
        return nil, err
    }

    return &PostDto{
        Id:     httpResponse.Response.Id,
        UserId: httpResponse.Response.UserId,
        Title:  httpResponse.Response.Title,
    }, nil
}
```

### Using Messaging

Configure Kafka producers in `messaging.yaml`:

```yaml
messaging_config:
  enabled: true
  producers:
    myProducer:
      enabled: true
      type: kafka
      async: true
      topic: ${KAFKA_TOPIC_NAME}
      endpoint: ${KAFKA_BROKER_ENDPOINT}
      session.timeout.ms: 1000
      message_timeout_ms: 1000
      concurrency: 5
      properties:
        acks: 0
        linger.ms: 1000
        batch.size: 65536
```

Use the messaging factory (provided via dependency injection):

```go
type Service struct {
    fx.In
    MessagingFactory messaging.MessagingFactory
}

func (s *Service) SendMessage() error {
    // Get producer from messaging factory
    producer, err := s.MessagingFactory.GetProducer("metrics")
    if err != nil {
        return errors.Wrap(err, "failed to get producer")
    }

    // Send message to Kafka
    ch := producer.Send(ctx, &messaging.Message{
        Key:     "message-key",
        Payload: map[string]interface{}{"key": "value"},
    })

    // Wait for result
    result := <-ch
    if result.Err != nil {
        return errors.Wrap(result.Err, "failed to publish to kafka")
    }
    return nil
}
```

Implement a consumer:

```go
// Get consumer from messaging factory
consumer, err := mf.GetConsumer("notifications")
if err != nil {
    return errors.Wrap(err, "failed to get consumer")
}

// Start processing messages
err = consumer.Process(
    context.Background(),
    messaging.NewSimpleConsumeFunction(
        cf,
        "message-processor",
        func(message *messaging.Message) error {
            // Process the message
            slog.Info("Received message", "key", message.Key, "payload", message.Payload)
            return nil
        },
        func(message *messaging.Message, err error) {
            slog.Error("Failed to process message", "error", err)
        },
    ),
)
if err != nil {
    return errors.Wrap(err, "failed to start consumer")
}
```

### Using Metrics

Configure metrics in `app.yaml`:

```yaml
metric:
  enabled: true
  prefix: "app"
  reporting_interval_ms: 1000
  enable_prometheus: true
```

Use metrics through CrossFunction:

```go
type Service struct {
    fx.In
    gox.CrossFunction
}

func (s *Service) DoSomething() error {
    // Record counter metrics
    s.Metric().Counter("api_calls").Inc(1)
    
    // Record timing metrics
    timer := s.Metric().Timer("api_latency").Start()
    defer timer.Stop()
    
    // Record gauge metrics
    s.Metric().Gauge("active_connections").Update(10)
    
    // Add tags to metrics
    s.Metric().Counter("requests").WithTags(map[string]string{
        "endpoint": "/users",
        "method": "GET",
        "status": "200",
    }).Inc(1)
    
    return nil
}
```

Available metric types:

1. Message Send Metrics:
```
<prefix>_message_send_{topic}
Labels:
- status: ok | error
- error: produce_failed | failed_after_produce | timeout | payload_error
- mode: sync | async
```

2. Message Consume Metrics:
```
<prefix>_message_consumed_{topic}
Labels:
- status: ok | error
- error: <error types>
- mode: sync | async
```

### Performance Profiling (pprof)

The project includes pprof integration for performance profiling. To enable pprof, set `enable_pprof: true` in your `app.yaml`:

```yaml
app:
  enable_pprof: true
```

Once enabled, you can access the following pprof endpoints:
- CPU Profile: `http://localhost:6060/debug/pprof/profile`
- Heap Profile: `http://localhost:6060/debug/pprof/heap`
- Goroutine Profile: `http://localhost:6060/debug/pprof/goroutine`
- Thread Create Profile: `http://localhost:6060/debug/pprof/threadcreate`
- Block Profile: `http://localhost:6060/debug/pprof/block`

3. Use pprof tool:
```bash
# CPU profile analysis
go tool pprof http://localhost:6060/debug/pprof/profile

# Memory profile analysis
go tool pprof http://localhost:6060/debug/pprof/heap
```

### Metrics Integration

The project supports metrics integration with Prometheus. Configure metrics in `app.yaml`:

```yaml
metric:
  enabled: true
  prefix: "app"
  reporting_interval_ms: 1000
  enable_prometheus: true
  
  # Tracing configuration (optional)
  tracing:
    enabled: false
```

Access Prometheus metrics at: `http://localhost:8080/metrics`

## 🛠️ Usage

### Development

To start the server in development mode:

```bash
sh build/run-local-dev.sh
```

### Staging

To run in staging environment:

```bash
sh build/run-local-stage.sh
```

## 📂 Project Structure

```
.
├── build/                  # Build and runtime scripts
├── cmd/                    # Application entry points
│   ├── server/            # Server application
│   └── tools/             # CLI tools
├── config/                # Configuration files
├── docs/                  # Documentation
│   └── img/               # Images and diagrams
├── internal/              # Private application code
│   └── handler/           # HTTP handlers
└── tests/                 # Test suites
    └── e2e/              # End-to-end tests
```

## 🧪 End-to-End Tests

The project includes comprehensive end-to-end tests in the `tests/e2e` directory. These tests verify the entire system's functionality by making actual HTTP requests and validating responses.

### Test Structure

```
tests/e2e/
├── e2e_test.go           # Test setup and utilities
└── e2e_post_test.go      # API endpoint tests
```

### Test Setup

The E2E tests use the testify suite for structured testing. Here's how the test environment is set up:

1. Environment Setup:
```go
type e2eTestSuite struct {
    suite.Suite
    restyClient *resty.Client
    ctx         context.Context
    done        context.CancelFunc
}
```

2. Suite Setup (runs once before all tests):
```go
func (s *e2eTestSuite) SetupSuite() {
    // Setup test environment
    env.SetupE2ETestEnv()

    // Allocate free ports for test services
    mapping, err := httpHelper.AllocateFreePortsAndAssignToEnvironmentVariables(
        "TEST_SERVICE", "")

    // Configure HTTP client
    s.restyClient = resty.New()
    s.restyClient.HostURL = fmt.Sprintf(
        "http://localhost:%s/%s/api/v1", 
        os.Getenv("HTTP_PORT"), 
        os.Getenv("APP_NAME"))
    s.restyClient.SetHeader("x-client-id", os.Getenv("CLIENT_ID"))
    s.restyClient.SetHeader("x-access-token", os.Getenv("CLIENT_TOKEN"))

    // Start the application
    s.ctx, s.done = context.WithTimeout(context.Background(), 30*time.Second)
    ch := make(chan bool, 1)
    go func() {
        command.FullMain(s.ctx, ch)
    }()
    <-ch
}
```

Key features of the test setup:
- Uses testify's suite for organized testing
- Automatically allocates free ports for test services
- Configures a Resty HTTP client with proper headers
- Starts the full application in test mode
- Supports debug logging for HTTP requests
- Timeout handling with context

Example E2E test implementation:

```go
func (s *e2eTestSuite) TestPostApi() {
    s.T().Run("Get Post - Success", func(t *testing.T) {
        // Make HTTP request
        resp, err := s.restyClient.R().
            SetHeader("Content-Type", "application/json").
            Get("/post/1")
        assert.NoError(t, err)
        assert.Equal(t, 200, resp.StatusCode())

        // Parse and validate response
        respMap := gox.StringObjectMap{}
        err = serialization.JsonBytesToObject(resp.Body(), &respMap)
        assert.NoError(t, err)

        // Verify response data
        assert.Equal(t, 1, respMap.IntOrZero("id"))
        
        // Optional: Print response for debugging
        fmt.Println("Get Post - Success Result\n", 
            respMap.JsonPrettyStringIgnoreError())
    })
}
```

The tests use:
- Test suite pattern for shared setup/teardown
- Resty client for making HTTP requests
- StringObjectMap for flexible JSON handling
- Built-in testing package assertions

### Running E2E Tests

1. Ensure the test environment is properly configured:
```bash
export TEST_ENV=local
```

2. Run the tests:
```bash
# Run all E2E tests
go test ./tests/e2e/... -v

# Run specific test
go test ./tests/e2e/... -v -run TestPostApi
```

3. Test with coverage:
```bash
go test ./tests/e2e/... -v -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## 👨‍💻 Development

### Pre-commit Hook

The project includes a pre-commit hook to ensure code quality:

```bash
sh build/pre-commit.sh
```

## 🤔 When to Use This Template

This template is particularly useful for:

- Microservice development where you need a solid foundation
- Projects that require extensive configuration management
- Applications with complex messaging requirements
- Systems that need robust monitoring and metrics
- Services that will scale and need performance profiling

## 📊 Comparison with Alternatives

| Feature | Go Template Project | Standard Go Project | Other Templates |
|---------|:-------------------:|:-------------------:|:---------------:|
| Clean Architecture | ✅ | ❌ | Varies |
| Environment Config | ✅ | ❌ | ⚠️ |
| E2E Testing | ✅ | ❌ | ⚠️ |
| Messaging Integration | ✅ | ❌ | ❌ |
| Metrics & Profiling | ✅ | ❌ | ⚠️ |
| CI/CD Ready | ✅ | ❌ | ✅ |

## 🗺️ Roadmap

- [ ] GraphQL API support
- [ ] Container orchestration examples
- [ ] OpenTelemetry integration
- [ ] Serverless deployment examples
- [ ] Database migration tools

## 👥 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

Please read [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) for details on our code of conduct.

## 🙏 Acknowledgements

- [Uber's fx](https://github.com/uber-go/fx) for dependency injection
- [Testify](https://github.com/stretchr/testify) for testing
- All [contributors](https://github.com/yourusername/go-template-project/graphs/contributors) who participated in this project

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🔒 Security

If you discover a security vulnerability, please send an e-mail to security@example.com instead of using the issue tracker. All security vulnerabilities will be promptly addressed.
