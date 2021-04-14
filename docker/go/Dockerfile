################################
# STEP 1 build executable binary
################################
FROM golang:alpine AS stage1
WORKDIR /build
COPY /go .
# smaller binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s -extldflags "-static"' main.go

################################
# STEP 2 build a small image
################################
FROM scratch
COPY --from=stage1 /build/main /main
ENTRYPOINT [ "/main" ]
CMD [ "add-row neu-student, My Name, ID0000001" ]