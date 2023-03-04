PORT = 3000
NAME = python-api
FMT  = black --target-version=py311 --check --diff *.py || true
LINT = flake8 --max-line-length=88 --ignore=E203,W503
TEST = python -m pytest --verbose --junit-xml=junit.xml
RUN  = uvicorn main:app --host 0.0.0.0 --port 3000 --reload

setup:
	@pip install --upgrade pip > /dev/null 2>&1
	@pip install --requirement requirements.txt > /dev/null 2>&1
	@echo "$(NAME) Setup complete"

fmt:
	$(FMT)

lint:
	$(LINT)

test:
	$(TEST)

run:
	$(RUN)

all: setup fmt lint test run

.PHONY: setup fmt lint test run all
