default: build

build: 
	@WITH_BUILDKIT=1 docker build -t spam:latest -f ./Dockerfile . 
