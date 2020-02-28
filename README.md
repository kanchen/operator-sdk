# operator-sdk
This project contains the Dockerfiles used to create Docker images to use operator-sdk

## goland operator-sdk

Based on [Operator SDK User Guide][operator_sdk_user_guide]
the prerequisites to use the operator-sdk

- [git][git_tool]
- [go][go_tool] version v1.12+.
- [mercurial][mercurial_tool] version 3.9+
- [docker][docker_tool] version 17.03+.
- [kubectl][kubectl_tool] version v1.11.3+.

The Dockerfile.go can be used to create an docker image with the tools installed

- [git][git_tool]
- [go][go_tool] version v1.14
- [mercurial][mercurial_tool] version latest
- [docker][docker_tool] version latest
- [kubectl][kubectl_tool] version latest

Use the following to create a docker image:

```sh
$ docker build -t kchen/operator-sdk:0.0.1 -f ./Dockerfile.go
```

Use the following to use the image:

```sh
$ run -it -v /var/run/docker.sock:/var/run/docker.sock kchen/operator-sdk:0.0.1 /bin/sh
```

[operator_sdk_user_guide]: https://github.com/operator-framework/operator-sdk/blob/master/doc/user-guide.md
[git_tool]: https://git-scm.com/downloads
[mercurial_tool]: https://www.mercurial-scm.org/downloads
[go_tool]: https://golang.org/dl/
[docker_tool]: https://docs.docker.com/install/
