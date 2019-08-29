PROJECT	=	myblog
REGISTRY = registry.cn-qingdao.aliyuncs.com/eveisgd
VERSION	=	1.1.0

run:
	@echo ====================run====================
	go run .

build:
	@echo ====================build====================
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

package:build
	@echo ====================package====================
	docker build -t $(PROJECT):$(VERSION) .


publish: package
	@echo ====================publish====================
	docker tag $(PROJECT):$(VERSION) $(REGISTRY)/$(PROJECT):$(VERSION)
	docker push $(REGISTRY)/$(PROJECT):$(VERSION)

clean:
	rm -r ./myblog

package-clean: clean
	docker images | grep -E "($(PROJECT)\s+)" | awk '{print $$3}' | uniq | xargs -I {} docker rmi --force {}