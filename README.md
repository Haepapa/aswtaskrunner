# aswtaskrunner

A solution that powers up a server, runs a specified process (e.g. data backup) and then shuts it down to save power.

# Usage

`docker build -t my-golang-app .`
`docker run --network host my-golang-app`

## development

`LOGIN_PASSWORD=<password> go run . storage02_plug`

## build

`go build -ldflags "-X main.buildPassword=<password>" -o checkplug`

## run the server from docker

`docker run -it -v ./devices.json:/app/devices.json \
 -p 8000:80 clementnerma/tapo-rest \
 --tapo-email '' \
 --tapo-password '' \
 --auth-password ''`
