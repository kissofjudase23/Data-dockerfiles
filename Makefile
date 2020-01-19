
cur_dir = $(shell pwd)
ubuntu_path ?= $(cur_dir)/Ubuntu
ubuntu1604_path ?= $(ubuntu_path)/1604
ubuntu1804_path ?= $(ubuntu_path)/1804

centos_path ?= $(cur_dir)/Centos
centos7_path ?= $(centos_path)/7
centos8_path ?= $(centos_path)/8

.PHONY: u18_up
u18_up:
	@$(MAKE) -C  $(ubuntu1804_path) up

.PHONY: u18_down
u18_down:
	@$(MAKE) -C  $(ubuntu1804_path) down

.PHONY: u18_tty
u18_tty:
	@$(MAKE) -C  $(ubuntu1804_path) tty

.PHONY: c8_up
c8_up:
	@$(MAKE) -C  $(centos8_path) up

.PHONY: c8_down
c8_down:
	@$(MAKE) -C  $(centos8_path) down

.PHONY: c8_tty
c8_tty:
	@$(MAKE) -C  $(centos8_path) tty

