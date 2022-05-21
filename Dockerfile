FROM golang:latest



WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o go-dock .

EXPOSE 8300

CMD [ "./go-dock" ]