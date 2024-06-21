mode con: cols=150 lines=60
@echo off
set /p input=please input your database backup absolute path like 'd:\app\oa.sql':^


import_vastTool.exe %input%
pause