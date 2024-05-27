# gombobulator

Cloud-native enterprise-grade extremely parallel processing unit, also called "the calculator".

Toy project to learn more go, kubernetes and AWS. Makes very little sense ;)

## Running

```sh
go build
go run . &
curl 127.0.0.1:7890/add --data "[1,2,3]"
```