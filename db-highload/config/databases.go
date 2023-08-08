package config

type Databases struct {
	Oracle DatabaseOracle
}

type (
	DatabaseOracle struct {
		Host     string
		Port     int
		Service  string
		User     string
		Password string
	}
)
