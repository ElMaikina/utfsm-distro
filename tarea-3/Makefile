# List of directories and their corresponding Go files and counts
DIRS_AND_FILES := \
    data_node:data.go:1 \
    director:director.go:1 \
    dosh_bank:bank.go:1 \
    mercenarios:cpu.go:7 \
    mercenarios:player.go:1 \
    name_node:name.go:1

# Default target
all: run_all

# Rule to run the Go file in each directory the specified number of times
define run_go
    dir = $(word 1, $(subst :, ,$1))
    file = $(word 2, $(subst :, ,$1))
    count = $(word 3, $(subst :, ,$1))
    @echo "Running go run $(file) in $(dir) $(count) times"
    @for i in `seq 1 $(count)`; do \
        (cd $(dir) && go run $(file)); \
    done
endef

# Phony target to run all commands
run_all: $(DIRS_AND_FILES)
	@$(foreach df,$(DIRS_AND_FILES),$(call run_go,$(df)))

.PHONY: all run_all
