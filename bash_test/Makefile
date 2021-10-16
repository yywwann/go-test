# GO related variables
GOBASE		:= ${shell pwd}
GOBIN		:= ${GOBASE}/bin
MAIN_PATH   := ${GOBASE}/${MAIN_DIR}
TESTOUT		:= ${GOBASE}/coverage.out
TIMESTAMP	:= ${shell date '+%Y/%m/%d %H:%M'}
PKGS		:= ${shell go list ./... | grep -v /vendor | grep -v /venv}

cat:
	@echo ┌ runing cat
	@echo ├ GOBASE = ${GOBASE}
	@echo ├ GOBIN = ${GOBIN}
	@echo ├ MAIN_PATH = ${MAIN_PATH}
	@echo ├ TESTOUT = ${TESTOUT}
	@echo ├ TIMESTAMP = ${TIMESTAMP}
	@#echo ├ PKGS = ${PKGS}
	@echo └ done

test:
	@echo ┌ runing test
	@bash test.sh
	@bash build/test.sh
	@echo └ done




