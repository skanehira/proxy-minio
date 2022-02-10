# proxy-minio
プロキシ(nginx)を通してｍinioにアクセスするサンプル

```sh
$ docker compose up -d
$ docker run -it --net proxy-minio -v $PWD:/go/src/proxy-minio --name proxy-api golang
$ cd go/src/proxy-minio && go run main.go
```
