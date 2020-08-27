package bitconfig

// Config holds all application configuration values
type Config struct {
	Address      string
	App          string
	DBIP         string
	DBPort       int
	Env          string
	ReadTimeout  int
	WriteTimeout int
	Static       string
}
