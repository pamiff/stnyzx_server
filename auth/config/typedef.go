package config

type Config struct {
	ServiceName string
	MiniProgram MiniProgram
	Jwt Jwt
}

type MiniProgram struct {
	ApiId string
	Secret string
}

type Jwt struct {
	SignSecret string
}