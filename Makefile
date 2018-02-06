
help:
	@echo "Available commands:"
	@echo "make install			Install dependencies."
	@echo "make run			Run inside container."
	@echo "make build			Build binary."
	@echo "make test			Run tests."
	@echo "make coverage			Show coverage in html."
	@echo "make clean			Clean build files."

install:
	@echo "Make: Install"
	./scripts/container.sh scripts/install.sh

.PHONY: test
test:
	@echo "Make: Test"
	./scripts/test.sh

coverage:
	@echo "Make: Coverage"
	./scripts/coverage.sh

run:
	@echo "Make: Run"
	./scripts/container.sh scripts/run.sh

build:
	@echo "Make: Build"
	./scripts/container.sh scripts/build_amd64.sh

clean:
	@echo "Make: Clean"
	rm -rf vendor
	rm -rf temp
	rm -rf _tmp_views
	touch nobreak_yt.db && rm nobreak_yt.db
	touch coverage.out && rm coverage.out
	touch coverage.html && rm coverage.html
