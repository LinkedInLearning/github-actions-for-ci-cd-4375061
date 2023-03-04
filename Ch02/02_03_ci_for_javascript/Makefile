PORT = 3000
NAME = javascript-api
LINT = npm run lint
TEST = npm run test
RUN  = npm start

setup:
	@npm install > /dev/null 2>&1
	@echo "$(NAME) Setup complete"

lint:
	$(LINT)

test:
	$(TEST)

run:
	$(RUN)

all: setup lint test run

.PHONY: setup lint test run all
