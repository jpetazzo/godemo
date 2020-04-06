FROM golang
COPY godemo.go .
ENV CGO_ENABLED=0
RUN go build godemo.go

FROM scratch
COPY --from=0 /go/godemo .
ENV LISTEN=:88
EXPOSE 88
CMD ["/godemo"]
