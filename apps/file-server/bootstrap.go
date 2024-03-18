package main

import "github.com/cza14h/nino-work/pkg/bootstrap"

func main() {

	conf, etcdRegistry := bootstrap.CommonBootstrap("fileService.client")

	fileService, ok := conf.Service["fileService"]
	if !ok {
		panic("File Service is not configured")
	}

}
