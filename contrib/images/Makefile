

all: fiammad 

fiammad: fiammad-rmi
	docker build --tag fiammachain/fiammad -f fiammad/Dockerfile \
    		$(shell git rev-parse --show-toplevel)

fiammad-rmi:
	docker rmi fiammachain/fiammad 2>/dev/null; true


.PHONY: all fiammad fiammad-rmi 
