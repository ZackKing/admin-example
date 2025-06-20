package conf

type AppConf struct {
	App     string `json:"app"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
	Debug   bool   `json:"debug"`
	Log     Log    `json:"log"`
	LogPath string `json:"logPath"`
}

type Log struct {
	Level string `json:"level"`
	Path  string `json:"path"`
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

type JwtConf struct {
	Exp        int    `json:"exp"`
	Iss        string `json:"iss"`
	SecretKey  string `json:"secretKey"`
	Algorithms string `json:"algorithms"`
}

type RedisConf struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}
