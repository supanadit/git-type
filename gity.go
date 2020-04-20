package gity

import (
	"errors"
	"strings"
)

// sshType is constant variable
const sshType string = "ssh"

// httpType is constant variable
const httpType string = "http"

// httpsType is constant variable
const httpsType string = "https"

// Type is the default model provided by the library
type Type struct {
	Url            string
	Type           string
	RepositoryName string
}

// Check is the function to check type of git
func Check(url string) (t Type, err error) {
	knownType := false

	httpsProtocol := httpsType + "://"
	httpProtocol := httpType + "://"

	if url[0:len(httpsProtocol)] == httpsProtocol || url[0:len(httpProtocol)] == httpProtocol {
		splitPath := strings.Split(url, "/")
		repositoryNameSplit := strings.Split(splitPath[len(splitPath)-1], ".")
		if len(repositoryNameSplit) == 2 {
			if repositoryNameSplit[1] == "git" {
				t.Url = url
				t.RepositoryName = repositoryNameSplit[0]
				if url[0:len(httpsProtocol)] == httpsProtocol {
					t.Type = httpsType
				} else {
					t.Type = httpType
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
							t.Url = url
							t.RepositoryName = repositoryNameSplit[0]
							t.Type = sshType
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
	return t, err
}

// IsHTTPS will return true if this url is type of HTTPS
func (gitType Type) IsHTTPS() bool {
	valid := false
	if gitType.Type == httpsType {
		valid = true
	}
	return valid
}

// IsHTTP will return true if this url is type of HTTP
func (gitType Type) IsHTTP() bool {
	valid := false
	if gitType.Type == httpType {
		valid = true
	}
	return valid
}

// IsHTTPORS will return true if this url is type of HTTP or HTTPS
func (gitType Type) IsHTTPORS() bool {
	return gitType.IsHTTPS() || gitType.IsHTTP()
}

// IsSSH will return true if this url is type of SSH
func (gitType Type) IsSSH() bool {
	valid := false
	if gitType.Type == sshType {
		valid = true
	}
	return valid
}
