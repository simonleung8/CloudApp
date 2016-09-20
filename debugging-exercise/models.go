package main

type apps_json struct {
	Resources []resources `json:"resources"`
}
type resources struct {
	Metadata metadata `json:"metadata"`
	Entity   entity   `json:"entity"`
}

type metadata struct {
	Guid string `json:"guid"`
}
type entity struct {
	Env                env    `json:"environment_json"`
	Name               string `json:"name"`
	HealthCheckTimeout int    `json:"health_check_timeout",null`
}

type env struct {
	AppID string `json:"app_id"`
}
