package config

type AppConfig struct {
	AppName            string
	JWTSecret          []byte
	JWTExpireInMinutes int64
}

var Config = AppConfig{
	AppName:            "test",
	JWTSecret:          []byte("very-secret"),
	JWTExpireInMinutes: 1500,
}
