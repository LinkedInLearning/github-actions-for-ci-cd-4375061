PORT ?= 3000
NAME ?= javascript-api
TEST ?= npm run test
RUN  ?= npm start

setup:
	@npm install > /dev/null 2>&1
	@echo "$(NAME) Setup complete"

test:
	$(TEST)

run:
	$(RUN)
