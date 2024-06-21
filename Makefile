# usage
# eg. make release VERSION=v0.0.1
# Binary name
BINARY=dmTool
# Builds the project
build:
		GO111MODULE=on go build -o ${BINARY} -ldflags "-X main.Version=${VERSION}"
		GO111MODULE=on go test -v
# Installs our project: copies binaries
install:
		GO111MODULE=on go install
release:
		# Clean
		go clean
		rm -rf dmTool-*.zip


		# Build for arm
#		go clean
#		CGO_ENABLED=0 GOOS=linux GOARCH=arm64 GO111MODULE=on go build -o export_dmTool -tags exp -ldflags "-s -w -X main.Version=${VERSION}"
#		go clean
#		CGO_ENABLED=0 GOOS=linux GOARCH=arm64 GO111MODULE=on go build -o import_dmTool -tags imp -ldflags "-s -w -X main.Version=${VERSION}"
#		rm -rf linux_arm/export-tool/*
#		rm -rf linux_arm/import-tool/*
#		mv export_dmTool linux_arm/export-tool
#		mv import_dmTool linux_arm/import-tool
#		cp settings.yaml start_export.sh linux_arm/export-tool
#		cp settings.yaml start_import.sh linux_arm/import-tool
#		rm -rf linux_arm/export-tool/.DS_Store
#		rm -rf linux_arm/import-tool/.DS_Store
#		zip -r ${BINARY}-linux-arm-x64-${VERSION}.zip linux_arm

		# Build for linux
#		go clean
#		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o export_dmTool -tags exp -ldflags "-s -w -X main.Version=${VERSION}"
#		go clean
#		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o import_dmTool -tags imp -ldflags "-s -w -X main.Version=${VERSION}"
#		rm -rf linux_x86/export-tool/*
#		rm -rf linux_x86/import-tool/*
#		mv export_dmTool linux_x86/export-tool
#		mv import_dmTool linux_x86/import-tool
#		cp settings.yaml start_export.sh linux_x86/export-tool
#		cp settings.yaml start_import.sh linux_x86/import-tool
#		rm -rf linux_x86/export-tool/.DS_Store
#		rm -rf linux_x86/import-tool/.DS_Store
#		zip -r ${BINARY}-linux-x64-${VERSION}.zip linux_x86

		# Build for win
		go clean
		CGO_ENABLED=0 GOOS=windows GOARCH=amd64 GO111MODULE=on go build -o exp_dmTool.exe -tags exp -ldflags "-s -w -X main.Version=%VERSION%"
		go clean
		CGO_ENABLED=0 GOOS=windows GOARCH=amd64 GO111MODULE=on go build -o imp_dmTool.exe -tags imp -ldflags "-s -w -X main.Version=%VERSION%"
        mkdir binary\win\dm-tool
        del /f /s /q binary\win\dm-tool\*
		move exp_dmTool.exe binary\win\dm-tool
		move imp_dmTool.exe binary\win\dm-tool
		copy example.yaml  binary\win\dm-tool
		copy start_export.bat binary\win\dm-tool
		copy start_import.bat binary\win\dm-tool
		"C:\Program Files\7-Zip\7z" a %BINARY%-win-x64-%VERSION%.zip binary\win\dm-tool
		go clean
# Cleans our projects: deletes binaries
clean:
		go clean
		del /f /s /q dmTool-win*.zip

.PHONY:  clean build