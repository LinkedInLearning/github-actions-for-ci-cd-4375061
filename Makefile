README_FILES = $(shell find . -name 'README.md' -not -path '*/node_modules/*' -not -path './.git/*')

spellcheck: $(README_FILES)
	# Foreach README.md file, run aspell
	@for file in $(README_FILES); do \
		aspell check --mode=markdown --lang=en "$${file}"; \
	done

clean:
	# Remove all backup files created by aspell
	@find . -name '*.md.bak' -exec rm -vf {} \;

all: spellcheck clean

.PHONY: all spellcheck clean
