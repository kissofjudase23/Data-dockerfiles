
cur_dir = $(shell pwd)
ubuntu_path ?= $(cur_dir)/Ubuntu
ubuntu1804_path ?= $(ubuntu_path)/1804

.PHONY: u18_up
u18_up:
	@$(MAKE) -C  $(ubuntu1804_path) up

.PHONY: u18_down
u18_down:
	@$(MAKE) -C  $(ubuntu1804_path) down

.PHONY: u18_tty
u18_tty:
	@$(MAKE) -C  $(ubuntu1804_path) tty



