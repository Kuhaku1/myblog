objects=myblog
version=1.1.0

build:
	# python

package:
	docker build $(objects}):$(version) .