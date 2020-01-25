FROM golang:1.11.0
RUN mkdir /app
ADD . /app
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
EXPOSE 8081
RUN go build -o main .
CMD "/app/main"

# docker build -t expample .
# docker run -d -p 8088:8081 expample