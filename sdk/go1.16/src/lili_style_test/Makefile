SSH_HOST ?= ssh-1.mc.lolipop.jp
SSH_PORT ?= 34203
SSH_USER ?= morning-yakushima-5627
ROTATE   ?= 5

GO ?= GO111MODULE=on go

RELEASE_FILES = server src sheets

build:
	$(GO) build -o server

server: build
	./server

deploy:
	$(eval SUFFIX = $(shell date '+%Y%m%d-%H%M%S'))
	$(eval KEEP = $(shell expr ${ROTATE} + 1))
	rm -rf ./tmp.dist
	mkdir tmp.dist/

	cp -r ${RELEASE_FILES} tmp.dist/
	ssh -p ${SSH_PORT} ${SSH_USER}@${SSH_HOST} 'mkdir -p /var/app/releases'
	scp -r -P ${SSH_PORT} ./tmp.dist ${SSH_USER}@${SSH_HOST}:/var/app/releases/${SUFFIX}
	ssh -p ${SSH_PORT} ${SSH_USER}@${SSH_HOST} 'rm -f /var/app/current'
	ssh -p ${SSH_PORT} ${SSH_USER}@${SSH_HOST} 'ln -sf /var/app/releases/${SUFFIX} /var/app/current'
	ssh -p ${SSH_PORT} ${SSH_USER}@${SSH_HOST} 'chmod +x /var/app/current/server'
	ssh -p ${SSH_PORT} ${SSH_USER}@${SSH_HOST} 'ls -t /var/app/releases/* | tail -n+${KEEP} | xargs --no-run-if-empty rm -rf'
	rm server

logs-out:
	ssh -p ${SSH_PORT} ${SSH_USER}@${SSH_HOST} "find /var/log/container/ -type f -name haconiwa.out | xargs ls --full-time | sort -k6,7 | awk '{print \$$9}' | xargs -i cat {}"

logs-err:
	ssh -p ${SSH_PORT} ${SSH_USER}@${SSH_HOST} "find /var/log/container/ -type f -name haconiwa.err | xargs ls --full-time | sort -k6,7 | awk '{print \$$9}' | xargs -i cat {}"