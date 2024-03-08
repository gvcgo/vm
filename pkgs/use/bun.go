package use

import (
	"runtime"

	"github.com/gvcgo/goutils/pkgs/gutils"
	"github.com/gvcgo/version-manager/pkgs/conf"
	"github.com/gvcgo/version-manager/pkgs/use/installer"
)

var BunInstaller = &installer.Installer{
	AppName:   "bun",
	Version:   "1.0.9",
	Fetcher:   conf.GetFetcher(),
	IsZipFile: true,
	FlagFileGetter: func() []string {
		r := []string{"bun"}
		if runtime.GOOS == gutils.Windows {
			r = []string{"bun.exe"}
		}
		return r
	},
	BinListGetter: func() []string {
		r := []string{"bun"}
		if runtime.GOOS == gutils.Windows {
			r = []string{"bun.exe"}
		}
		return r
	},
	DUrlDecorator:      installer.DefaultDecorator,
	StoreMultiVersions: true,
}

func TestBun() {
	zf := BunInstaller.Download()
	BunInstaller.Unzip(zf)
	BunInstaller.Copy()
	BunInstaller.CreateVersionSymbol()
	BunInstaller.CreateBinarySymbol()
	BunInstaller.SetEnv()
}
