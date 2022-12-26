GO = go build 

TARGETS := all bench build check clean run 

SUBDIRS := $(wildcard */.)

$(TARGETS): $(SUBDIRS)
$(SUBDIRS):
	@$(MAKE) -C $@ $(MAKECMDGOALS)

.PHONY: $(TARGETS) $(SUBDIRS) cyclo runtime

