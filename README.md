# go-study-swagger
project to study golang with swagger

## config
- add alias on .bashrc
`alias swagger="docker run --rm -t --user $(id -u):$(id -g) -e GOPATH=$HOME/go:/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger"`
