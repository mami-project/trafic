.DEFAULT_GOAL := help

NODES := $(shell docker-compose config --services | xargs)

up: ; docker-compose up -d

build-up: ; docker-compose up --build -d

down: ; docker-compose down

top: ; docker-compose top

#dash: ; bin/upload-dashboard.bash dashboards/overview.json

# Dynamically create each sh-<node> rule
define interactive_shell_template
sh-$(1): ; docker exec -ti $(1) bash
endef

$(foreach node,$(NODES),$(eval $(call interactive_shell_template,$(node))))

help:
	@echo
	@echo "Available targets:"
	@echo
	@echo " * up            start the testbed"
	@echo " * build-up      build images before starting the testbed"
	@echo " * down          stop the testbed"
	@echo " * top           display running processes"
	@echo " * sh-<node>     start interactive shell on <node>"
	@echo "                 where <node> is one of: $(NODES)"
	@echo

.PHONY: up build-up down top
