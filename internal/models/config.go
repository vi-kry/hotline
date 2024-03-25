package models

type Config struct {
	Env        string `yaml:"env" env-default:"local"`
	HTTPServer `yaml:"http_server"`
	ISSO       `yaml:"isso"`
	Token      `yaml:"token"`
}

type HTTPServer struct {
	Port string `yaml:"port" env-default:"8082"`
	Path string `yaml:"path" env-default:"hotline/v1"`
}

type ISSO struct {
	URL       string `yaml:"url" env-required:"true"`
	ClientID  string `yaml:"client_id" env-required:"true"`
	GrantType string `yaml:"grant_type" env-required:"true"`
}

type Token struct {
	Secret    string `yaml:"secret"`
	ExpiresAt int    `yaml:"expires_at"`
}
