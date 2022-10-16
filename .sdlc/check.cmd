
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
set bin_dir=%project_dir%\bin

echo script_name   %script_name%
echo script_path   %script_path%
echo script_dir    %script_dir%
echo project_dir   %project_dir%
echo module_name   %module_name%
echo bin_dir       %bin_dir%

set buildmode=readonly
IF DEFINED SDLC_GO_VENDOR (
    echo Using Go vendor
    set GOPROXY=off
    set buildmode=vendor
)

cd %project_dir%

SETLOCAL ENABLEDELAYEDEXPANSION
for /f %%x in ('dir /AD /B /S lib') do (
    echo --- go test lib %%x
    cd %%x
    call go test -mod %buildmode% -race ./...
    call go test -mod %buildmode% -cover ./...
)

cd %project_dir%

SETLOCAL ENABLEDELAYEDEXPANSION
for /f %%x in ('dir /AD /B /S cmd') do (
    echo --- go test cmd %%x
    cd %%x
    set bin_name=%%~nx
    set exe_path=%bin_dir%\!bin_name!.exe
    echo Build %module_name% cmd !bin_name! into !exe_path!
    IF EXIST !exe_path! (
        echo Deleting old bin !exe_path!
        DEL /F !exe_path!
    )
    call go build -mod %buildmode%  ^
         -ldflags "-s -X %module_name%/lib/core.Version=%APP_VERSION% -X %module_name%/lib/core.BuildTime=%TIMESTAMP% -X %module_name%/lib/core.GitCommit=win-dev-commit" ^
         -o !exe_path! "%module_name%/cmd/!bin_name!"
    call go test -mod %buildmode% -cover ./...
)

cd %project_dir%
REM go test -mod %buildmode% -v %module_name%/e2e
