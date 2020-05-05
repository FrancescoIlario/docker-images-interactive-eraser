############################
# STEP 1 build executable binary
############################
FROM golang:1.14.1-buster as builder

# Ensure ca-certficates are up to date
RUN update-ca-certificates

WORKDIR $GOPATH/src/github.com/FrancescoIlario/go-std-tmpl

# use modules
COPY go.mod .

ENV GO111MODULE=on
RUN go mod download \
    && go mod verify

COPY . .

# Build the binary
RUN mkdir -p /app/build \
    && go get -u ./... \
    && CGO_ENABLED=0 go build \
		-ldflags="-w -s" \
		-a \
		-installsuffix cgo \
		-o /app/build/go_std_app \
		cmd/go-std-tmpl/main.go

############################
# STEP 2 build a small image
############################
# using base nonroot image
# user:group is nobody:nobody, uid:gid = 65534:65534
FROM gcr.io/distroless/static

# Copy our static executable
COPY --from=builder /app/build/go_std_app /usr/bin/go_std_app

# Run the produced binary.
ENTRYPOINT ["/usr/bin/go_std_app"]
