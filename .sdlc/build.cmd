
@echo OFF
SETLOCAL ENABLEEXTENSIONS
SET "script_name=%~n0"
SET "script_path=%~0"
SET "script_dir=%~dp0"
rem # to avoid invalid directory name message calling %script_dir%\config.bat
cd %script_dir%
call config.bat
cd ..
set project_dir=%cd%

set module_name=%REPO_HOST%/%REPO_OWNER%/%REPO_NAME%
set exe_path=bin\%REPO_NAME%.exe

echo script_name   %script_name%
echo script_path   %script_path%
echo script_dir    %script_dir%
echo project_dir   %project_dir%
echo module_name   %module_name%
echo exe_path      %exe_path%

cd %project_dir%

IF EXIST %exe_path% DEL /F %exe_path%

@echo ON
SETLOCAL ENABLEDELAYEDEXPANSION
for /f %%x in ('dir /AD /B /S cmd') do (
    echo --- go build cmd %%x
    cd %%x
    set bin_name=%%~nx
call go build -mod vendor -ldflags "-s -X %module_name%/lib/core.Version=%APP_VERSION% -X %module_name%/lib/core.BuildTime=%TIMESTAMP% -X %module_name%/lib/core.GitCommit=win-dev-commit" ^
    -o %bin_dir%\!bin_name!.exe "%module_name%/cmd/!bin_name!"
)
