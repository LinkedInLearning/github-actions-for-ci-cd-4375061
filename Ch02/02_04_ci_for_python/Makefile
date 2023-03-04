PORT ?= 3000
NAME ?= python-api
TEST ?= python -m pytest --verbose --junit-xml=junit.xml
RUN  ?= uvicorn main:app --host 0.0.0.0 --port 3000 --reload

.PHONY: setup
setup:
	@pip install --upgrade pip > /dev/null 2>&1
	@pip install --requirement requirements.txt > /dev/null 2>&1
	@echo "$(NAME) Setup complete"

.PHONY: test
test:
	$(TEST)

.PHONY: run
run:
	$(RUN)
