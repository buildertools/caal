# RULE 1: Any containers created by this Makefile should be automatically cleaned up

PWD := $(shell pwd)

clean:
	rm bin/caal-exported
	docker rmi buildertools/caal:dev

# Using a Node ecosystem tool for service functional tests
prepare:
	docker run -it --rm -v "$(shell pwd)":/work -w /work node:4.1.2 npm install

reiterate:
	docker-compose build caal
	docker-compose up -d --no-deps --force-recreate caal

iterate:
	docker-compose kill
	docker-compose rm
	docker-compose up -d

client:
	docker-compose up -d --no-deps --force-recreate client

stop:
	docker-compose stop

# Using `docker cp` to copy a file out of an image requires three steps:
#  1. Create a container from the target image
#  2. Run `docker cp` to copy the file 
#  3. Remove the temporary container
# The biggest problem with this handshake is the need to maintain references to the
# target container. This is compounded in Makefiles. So, forget about that nonsense.
#   Instead I'm using a volume and a copy command from a self-destructing container.
# This has the nice property of being able to run in a single step and potentially 
# performing more complex copy operations.
build:
	docker build -t buildertools/caal:dev -f build.df .
	docker run --rm -v $(PWD)/bin:/xfer buildertools/caal:dev /bin/sh -c 'cp /go/bin/caal* /xfer/'

release: build
	docker build -t buildertools/caal:latest -f release.df .
	docker tag buildertools/caal:latest buildertools/caal:poc

