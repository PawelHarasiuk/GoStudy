FROM golang:1.22 as builder_stage

WORKDIR app/

COPY . .

RUN go mod tidy

RUN GOOS=linux GOARCH=amd64 go build -o /main main.go

FROM public.ecr.aws/lambda/go:latest

COPY --from=builder_stage /main /var/task/main

CMD ["main"]