
cur_dir = $(shell pwd)
ubuntu_path ?= $(cur_dir)/Ubuntu
ubuntu1604_path ?= $(ubuntu_path)/1604
ubuntu1804_path ?= $(ubuntu_path)/1804

centos_path ?= $(cur_dir)/Centos
centos7_path ?= $(centos_path)/7
centos8_path ?= $(centos_path)/8

.PHONY: u18_up u18_down u18_tty c8_up c8_down c8_tty
u18_up:
	@$(MAKE) -C  $(ubuntu1804_path) up

u18_down:
	@$(MAKE) -C  $(ubuntu1804_path) down

u18_tty:
	@$(MAKE) -C  $(ubuntu1804_path) tty

c8_up:
	@$(MAKE) -C  $(centos8_path) up

c8_down:
	@$(MAKE) -C  $(centos8_path) down

c8_tty:
	@$(MAKE) -C  $(centos8_path) tty
