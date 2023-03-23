package conf

type AppConf struct {
	App  string `json:"app"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

type DbConf struct {
	Type         string `json:"type"`
	Host         string `json:"host"`
	Port         int    `json:"port"`
	Db           string `json:"db"`
	User         string `json:"user"`
	Password     string `json:"password"`
	Charset      string `json:"charset"`
	MaxConns     int    `json:"maxConns"`
	MaxIdleConns int    `json:"maxIdleConns"`
}

type SettingsConf struct {
	StopLogin bool `json:"stopLogin"`
}

type JwtConf struct {
	Exp        int    `json:"exp"`
	Iss        string `json:"iss"`
	SecretKey  string `json:"secretKey"`
	Algorithms string `json:"algorithms"`
}
