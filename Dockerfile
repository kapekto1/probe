FROM golang:alpine as builder
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o probe .

FROM scratch
WORKDIR /
COPY --from=builder /app/probe .
EXPOSE 8080
CMD ["./probe"]
