package self

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gvcgo/goutils/pkgs/gutils"
	"github.com/gvcgo/version-manager/internal/cnf"
	"github.com/gvcgo/version-manager/internal/utils"
)

/*
set uninstall script.
*/
const (
	unInstallScriptName string = `vmr-uninstall`
)

var UnixRemoveScript string = `#!/bin/sh
cd ~
vmr Uins
rm -rf %s`

var WinRemoveScript string = `cd %HOMEPATH%
vmr Uins
rmdir /s /q `

var WinMingwRemoveScript string = `#!/bin/sh
powershell %s`

func SetUninstallScript() {
	script := fmt.Sprintf(UnixRemoveScript, cnf.GetVMRWorkDir())
	scriptName := unInstallScriptName
	if runtime.GOOS == gutils.Windows {
		script = WinRemoveScript + cnf.GetVMRWorkDir()
		scriptName = unInstallScriptName + ".bat"
	}

	scriptPath := filepath.Join(cnf.GetVMRWorkDir(), scriptName)
	os.WriteFile(scriptPath, []byte(script), os.ModePerm)

	if runtime.GOOS == gutils.Windows {
		mingwScriptPath := filepath.Join(cnf.GetVMRWorkDir(), unInstallScriptName+".sh")
		mingwScript := fmt.Sprintf(WinMingwRemoveScript, utils.ConvertWindowsPathToMingwPath(scriptPath))
		os.WriteFile(mingwScriptPath, []byte(mingwScript), os.ModePerm)
	}

	if runtime.GOOS != gutils.Windows {
		gutils.ExecuteSysCommand(true, "", "chmod", "+x", scriptPath)
	}
}
