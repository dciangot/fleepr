FROM golang:alpine as BUILD

RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo  -o fleepr .

FROM alpine as APP

COPY --from=0 /app/fleepr /usr/bin/fleepr

ENTRYPOINT ["fleepr"]