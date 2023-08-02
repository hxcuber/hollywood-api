FROM golang:latest
LABEL authors="hxcuber"

WORKDIR /Users/hxcuber/Desktop/goo/hollywood-api/api/

COPY api/go.mod api/go.sum  ./
RUN go mod download && go mod verify

COPY . .
RUN go build -C api/cmd/hollywood-api/

ENTRYPOINT ["api/cmd/hollywood-api/hollywood-api"]
EXPOSE 3000