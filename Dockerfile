FROM golang:1.19 as builder
WORKDIR /code
COPY ./ ./

RUN go mod download
RUN go mod verify

# `skaffold debug` sets SKAFFOLD_GO_GCFLAGS to disable compiler optimizations
ARG SKAFFOLD_GO_GCFLAGS
RUN go build -gcflags="${SKAFFOLD_GO_GCFLAGS}" -trimpath -o /app

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=builder /app /app

ENV GOTRACEBACK=single

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/app"]

#EXPOSE 8080
#CMD [ "/app" ]
