package service

import (
	"inkafarma/devops_tools/jenkinscli"
	"inkafarma/webcindi/model"
)

func CrearConfig(env_id uint) *jenkinscli.JbaseJenkins {
	var Mjenkins = model.JenkinsModel{}
	Ijenkins, _ := Mjenkins.Get(env_id)

	return jenkinscli.New("/home/jaivic/go/src/inkafarma/devops_tools/jenkinscli/jenkins-cli.jar",
		Ijenkins.Ip,
		"8080",
		Ijenkins.Account,
		Ijenkins.Password, "")
}
func BastionUpdatePass(env_id uint, name string, pass string) {

	param := "USUARIO=" + name
	param1 := "PASSWORD=" + pass
	j := CrearConfig(env_id)
	j.BuildJob("devops-bastion-update-pass", param, "-p", param1)
}
