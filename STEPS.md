Create 3 new service with command

```
kit n s users --module github.com/ibadi-id/packagemain/users
kit n s bugs --module github.com/ibadi-id/packagemain/bugs
kit n s notificator --module github.com/ibadi-id/packagemain/notificator
```

Set Function in each sevice

```
bugs:
Create(ctx context.Context, bug string) error

notificator
SendEmail(ctx context.Context, email string, content string) error

users:
Create(ctx context.Context, email string) error
```

Generate code for all service with middleware and gorilla mux

```
kit g s users --dmw --gorilla
kit g s bugs --dmw --gorilla
kit g s notificator -t grpc --dmw
```

Generate Docker Compose File

```
kit g d

add
WORKDIR /go/src/github.com/ibadi-id/projectmain/bugs

replace
RUN go get -t ./...
with
RUN go mod tidy

```

Change docker compose volume

```
- .:/go/src/github.com/ibadi-id/projectmain
```

Change notificator.proto

```
option go_package = "github.com/ibadi-id/packagemain/notificator/pkg/grpc/pb";

message SendEmailRequest {
string email =1;
string content =2;
}
message SendEmailReply {
string id =1;
}

cd ./projectmain/notificator/pkg/grpc/pb/
./compile.sh
```
