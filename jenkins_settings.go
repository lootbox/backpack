package main

type jenkinsCfg struct {
	JenkinsURL      string
	JenkinsUser     string
	JenkinsPassword string
	JenkinsKey      string
}

func newJenkinsCfg() *jenkinsCfg {
	return &jenkinsCfg{
		JenkinsURL:      getEnv("JenkinsURL", ""),
		JenkinsUser:     getEnv("JenkinsUser", ""),
		JenkinsPassword: getEnv("JenkinsPassword", "admin"),
		JenkinsKey:      getEnv("JenkinsKey", "admin"),
	}
}
