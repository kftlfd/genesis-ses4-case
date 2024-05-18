FROM golang:1.21-alpine as build
WORKDIR /src
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \ 
    CGO_ENABLED=0 GOOS=linux go build -o /app .

FROM scratch
COPY --from=build /app /app
EXPOSE 8000
CMD [ "/app" ]
