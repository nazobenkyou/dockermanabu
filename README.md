# Docker 学ぶために

## Requirements

- Golang - https://golang.org/dl/
- Docker - https://docs.docker.com/install/
- Docker compose - https://github.com/docker/compose/releases
- Bash/zsh/... console/linux console on windows https://itsfoss.com/install-bash-on-windows/
- git - https://git-scm.com/downloads

## Get started

```bash
go mod download

CGO_ENABLED=0 go build
```

## Docker

### dockerignoreというえば

.dockerignoreファイルgitignoreと同じ目標だ、containerでコピーしないファイルということです

### dockefile

Dockerfileの中へ

### dockerコマンドについて

docker build -t \<ネーム\>:\<タッグ\> .

Dockerfileのところでビルドすることができます、ビルドときにARGというargumentsのことで書くことができます

docker run <ファイル> containerを動くすることです

### デバッグについて

containerの中にはエラー出てるときにこのコマンド使ってください

まずCMDのdockerfileコメントして
それでビルドと。。。

docker run -it \<イメージ\>:\<タッグ\> /bin/sh

itのparameterはiteractiveということで,そしてalpineはshスクリプトあるのでshのshellをしよう

そしてマイクロサービスのコマンドをして