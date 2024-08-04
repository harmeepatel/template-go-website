FROM golang:1.22

WORKDIR /src

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

Run go build -o bin/ii

EXPOSE 80

ENV DEPLOY_PLATFORM="prod"
CMD ["bin/ii"]
