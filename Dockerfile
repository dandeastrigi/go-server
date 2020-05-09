FROM golang:1.14
ENV APP_NAME app
ENV PORT 8080

COPY /code/ /go/src/code
WORKDIR /go/src/code/${APP_NAME}

RUN go get ./
RUN go build -o ${APP_NAME}

#CMD ./${APP_NAME}

#EXPOSE ${PORT}