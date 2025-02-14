package app

import (
	"log"
	"time"

	"code.cloudfoundry.org/system-metrics-scraper/pkg/config"

	envstruct "code.cloudfoundry.org/go-envstruct"
)

type Config struct {
	// Loggregator Agent Certs
	ClientKeyPath  string `env:"CLIENT_KEY_PATH,  report, required"`
	ClientCertPath string `env:"CLIENT_CERT_PATH, report, required"`
	CACertPath     string `env:"CA_CERT_PATH,     report, required"`

	// System Metrics Agent Certs
	MetricsKeyPath    string `env:"SYSTEM_METRICS_KEY_PATH,     report, required"`
	MetricsCertPath   string `env:"SYSTEM_METRICS_CERT_PATH,    report, required"`
	MetricsCACertPath string `env:"SYSTEM_METRICS_CA_CERT_PATH, report, required"`
	MetricsCN         string `env:"SYSTEM_METRICS_CA_CN,        report, required"`

	// Leadership Election
	LeadershipServerAddr string `env:"LEADERSHIP_SERVER_ADDR,           report, required"`
	LeadershipKeyPath    string `env:"LEADERSHIP_ELECTION_KEY_PATH,     report, required"`
	LeadershipCertPath   string `env:"LEADERSHIP_ELECTION_CERT_PATH,    report, required"`
	LeadershipCACertPath string `env:"LEADERSHIP_ELECTION_CA_CERT_PATH, report, required"`

	LoggregatorIngressAddr string `env:"LOGGREGATOR_AGENT_ADDR, report, required"`
	DefaultSourceID        string `env:"DEFAULT_SOURCE_ID,      report, required"`

	ScrapeInterval time.Duration `env:"SCRAPE_INTERVAL, report"`
	ScrapeTimeout  time.Duration `env:"SCRAPE_TIMEOUT,  report"`
	ScrapePort     int           `env:"SCRAPE_PORT,     report, required"`

	DNSFile string `env:"DNS_FILE, report, required"`

	NatsHosts    []string `env:"NATS_HOSTS, required"`
	NatsCAPath   string   `env:"NATS_CA_PATH, required, report"`
	NatsCertPath string   `env:"NATS_CERT_PATH, required, report"`
	NatsKeyPath  string   `env:"NATS_KEY_PATH, required, report"`

	MetricsServer config.MetricsServer
}

func LoadConfig(log *log.Logger) Config {
	cfg := Config{
		ScrapeInterval: time.Minute,
		ScrapeTimeout:  time.Second,
	}

	if err := envstruct.Load(&cfg); err != nil {
		log.Fatal(err)
	}

	envstruct.WriteReport(&cfg) //nolint:errcheck

	return cfg
}
