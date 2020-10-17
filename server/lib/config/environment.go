package config

func EnvironmentName(osEnv string) string {
	if osEnv == "" {
		osEnv = "local"
	}

	return osEnv
}
