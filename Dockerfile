#build stage
# build go lang image
FROM golang:1.18.2-alpine as build-env

#set environment variable
ENV APP_NAME array
ENV CMD_PATH array_reverse.go

#copy application data into image
WORKDIR $GOPATH/gosrc2/${APP_NAME}
COPY . .

#build application
RUN go build -v -o /${APP_NAME} ${CMD_PATH}

#run stage
FROM alpine:3.14

#set environment variable
ENV APP_NAME array

#copy only required data into this image
COPY --from=build-env /${APP_NAME} .

#expose application port
EXPOSE 8081

#start app
CMD ./${APP_NAME}


