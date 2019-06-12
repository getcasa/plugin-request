# Request
This plugin is a part of [Casa](https://github.com/getcasa), it's used to send http requests.

## Downloads
Use the integrated store in casa or [github releases](https://github.com/getcasa/plugin-request/releases).

## Build
```
sudo env CGO_ENABLED=1 GOOS=linux GOARCH=arm64 CC=aarch64-linux-gnu-gcc go build -buildmode=plugin -o request.so *.go
```

## Install
1. Extract `request.zip`
2. Move `request` folder to casa `plugins` folder
3. Restart casa
