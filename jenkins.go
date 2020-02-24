package main

import (
	"fmt"
	"github.com/bndr/gojenkins"
)

func server() *gojenkins.Jenkins {
	cfg := newJenkinsCfg()
	jenkins := gojenkins.CreateJenkins(
		nil,
		cfg.JenkinsURL,
		cfg.JenkinsUser,
		cfg.JenkinsPassword)

	jenkins, _ = jenkins.Init()
	return jenkins
}

func listNodes() {
	jenkins := server()
	nodes, _ := jenkins.GetAllNodes()

	for _, node := range nodes {
		node.Poll()
		nodeStatus, _ := node.IsOnline()
		if nodeStatus {
			fmt.Println("Node is Online: ", node.GetName())
		}
	}
}

func addSSHCreds() {
	cred := gojenkins.SSHCredentials{
		Scope:       "global",
		ID:          "sshID",
		Username:    "RANDONMANE",
		Passphrase:  "password",
		Description: "EXAMPLE",
		PrivateKeySource: &gojenkins.PrivateKeyFile{
			Value: "testValueofkey",
			Class: gojenkins.KeySourceOnMasterType,
		},
	}

	cm := gojenkins.CredentialsManager{J: server()}
	cm.Add("_", cred)
}
