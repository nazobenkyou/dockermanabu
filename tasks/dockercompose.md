# docker-composeのタスク

## 初めてから

今回docker-compose使いたいと思う、

シンプルで書いたので分からないことコメントしてください

さー！始まりましょ！

```bash
docker-compose up --build
```

これ二つのcontainer始まる:

- app
- mongodb

appやってみましょ

browserで`http://localhost:8080/users`入れよ

そしてconsoleでcurlしよう

```bash
curl -XPOST -d '{"name": "tero", "surname": "gohan"}' http://localhost:8080/users
```
