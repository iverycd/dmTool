set BINARY=dmTool
set VERSION=v%1


del /f /s /q dmTool-win*.zip
rem Build for win
C:\Users\Administrator\sdk\go1.22.4\bin\go.exe clean
C:\Users\Administrator\sdk\go1.22.4\bin\go.exe build -o exp_dmTool.exe -tags exp -ldflags "-s -w -X main.Version=%VERSION%"
C:\Users\Administrator\sdk\go1.22.4\bin\go.exe clean
C:\Users\Administrator\sdk\go1.22.4\bin\go.exe build -o imp_dmTool.exe -tags imp -ldflags "-s -w -X main.Version=%VERSION%"
mkdir binary\win\dm-tool
del /f /s /q binary\win\dm-tool\*
move exp_dmTool.exe binary\win\dm-tool
move imp_dmTool.exe binary\win\dm-tool
copy settings.yaml  binary\win\dm-tool
copy start_export.bat binary\win\dm-tool
copy start_import.bat binary\win\dm-tool
ROBOCOPY dm_client binary\win\dm-tool\dm_client /E

"C:\Program Files\7-Zip\7z" a %BINARY%-win-x64-%VERSION%.zip binary\win\dm-tool
pause