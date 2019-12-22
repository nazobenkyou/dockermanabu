# Docker 学ぶために

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