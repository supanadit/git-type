package gittype

import (
	"errors"
	"strings"
)

var sshType string = "ssh"
var httpType string = "http"
var httpsType string = "https"

type GitType struct {
	Url            string
	Type           string
	RepositoryName string
}

func NewGitType(url string) (GitType, error) {
	var gitType GitType
	var err error = nil
	knownType := false

	httpsProtocol := httpsType + "://"
	httpProtocol := httpType + "://"

	if url[0:len(httpsProtocol)] == httpsProtocol || url[0:len(httpProtocol)] == httpProtocol {
		splitPath := strings.Split(url, "/")
		repositoryNameSplit := strings.Split(splitPath[len(splitPath)-1], ".")
		if len(repositoryNameSplit) == 2 {
			if repositoryNameSplit[1] == "git" {
				gitType.Url = url
				gitType.RepositoryName = repositoryNameSplit[0]
				if url[0:len(httpsProtocol)] == httpsProtocol {
					gitType.Type = httpsType
				} else {
					gitType.Type = httpType
				}
				knownType = true
			}
		}
	} else {
		ssh := strings.Split(url, "@")
		if len(ssh) == 2 {
			domainAndPath := strings.Split(ssh[1], ":")
			if len(domainAndPath) == 2 {
				verifyDomain := strings.Split(domainAndPath[0], ".")
				if len(verifyDomain) == 2 {
					splitPath := strings.Split(domainAndPath[1], "/")
					repositoryNameSplit := strings.Split(splitPath[len(splitPath)-1], ".")
					if len(repositoryNameSplit) == 2 {
						if repositoryNameSplit[1] == "git" {
							gitType.Url = url
							gitType.RepositoryName = repositoryNameSplit[0]
							gitType.Type = sshType
							knownType = true
						}
					}
				}
			}
		}
	}

	if !knownType {
		err = errors.New("unknown URL type")
	}
	return gitType, err
}

func (gitType GitType) IsHTTPS() bool {
	valid := false
	if gitType.Type == httpsType {
		valid = true
	}
	return valid
}

func (gitType GitType) IsHTTP() bool {
	valid := false
	if gitType.Type == httpType {
		valid = true
	}
	return valid
}

func (gitType GitType) IsHTTPORS() bool {
	return gitType.IsHTTPS() || gitType.IsHTTP()
}

func (gitType GitType) IsSSH() bool {
	valid := false
	if gitType.Type == sshType {
		valid = true
	}
	return valid
}
