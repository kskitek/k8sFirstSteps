
all: build-red build-blue

build-%:
	GOOS=linux go build -o $* ./$*
	upx ./$*/$*
	docker build -t kskitek/k8sfirststeps:$*-v1 -f Dockerfile-$* .
	docker push kskitek/k8sfirststeps:$*-v1
