package conf

type database struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     int
}

type cors struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	AllowCredentials bool
	ExposeHeaders    []string
}
