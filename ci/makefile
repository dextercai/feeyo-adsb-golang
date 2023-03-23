NAME=feeyo-adsb-golang
BINDIR=/build
VERSION=$(shell git describe --tags || echo "unknown")
BUILDTIME=$(shell date "+%F %T %Z" -u)
GOBUILD=CGO_ENABLED=0 go build -trimpath -ldflags '-X "dextercai.com/feeyo-adsb-golang/constant.Version=$(VERSION)" \
		-X "dextercai.com/feeyo-adsb-golang/constant.BuildTime=$(BUILDTIME)" \
		-w -s -buildid='

all: docker # Most used

docker:
	$(GOBUILD) -o $(BINDIR)/$(NAME)

