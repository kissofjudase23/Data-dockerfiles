
cur_dir = $(shell pwd)
ubuntu_path ?= $(cur_dir)/Ubuntu
ubuntu1804_path ?= $(ubuntu_path)/1804


u18_up:
	@$(MAKE) -C  $(ubuntu1804_path) up

u18_down:
	@$(MAKE) -C  $(ubuntu1804_path) down

u18_tty:
	@$(MAKE) -C  $(ubuntu1804_path) tty



