# Git Type
Check Git URL is SSH, HTTPS or HTTP

```golang
type GitType struct {
	Url            string
	Type           string
	RepositoryName string
}
```

## Installation

```shell script
go get -u -v github.com/supanadit/git-type
```

## How To Use
```golang
import "github.com/supanadit/git-type"
```


## For Example :

### For SSH
```golang
gitType, err := gittype.NewGitType("git@github.com:supanadit/devops-factory.git")
if err != nil {
    panic(err)
}
fmt.Println(gitType.IsHTTPORS()) // false
fmt.Println(gitType.IsHTTP()) // false
fmt.Println(gitType.IsHTTPS()) // false
fmt.Println(gitType.IsSSH()) // true
```

### For HTTP / HTTPS
```golang
gitType, err := gittype.NewGitType("https://github.com/supanadit/devops-factory.git")
if err != nil {
    panic(err)
}
fmt.Println(gitType.IsHTTPORS()) // true
fmt.Println(gitType.IsHTTP()) // false
fmt.Println(gitType.IsHTTPS()) // true
fmt.Println(gitType.IsSSH()) // false
```
