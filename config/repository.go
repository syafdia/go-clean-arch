package config

type TokenRepositoryConfig struct {
	TokenExpiration int // Hour.
}

type RepositoryConfig struct {
	TokenRepository TokenRepositoryConfig
}
