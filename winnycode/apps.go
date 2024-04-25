package main

var categories = map[string][]string{
	"Browsers": {
		"Chromium",
		"Brave",
		"Vivaldi",
		"Microsoft Edge",
		"Opera",
		"Opera GX",
		"Firefox",
		"LibreWolf",
		"Waterfox",
		"Floorp Browser",
		"Tor Browser",
		"Avast Secure Browser",
		"Thorium",
		"Mercury",
	},
	"Messaging": {
		"Discord",
		"Zoom",
		"Skype",
		"Telegram",
		"Viber",
		"Microsoft Teams",
		"Thunderbird",
		"Pidgin",
	},
	"Media": {
		"VLC Media Player",
		"Spotify",
		"Audacity",
		"iTunes",
		"OBS",
		"Streamlabs",
		"Kdenlive",
		"Winamp",
		"foobar2000",
		"GOM Player",
	},
	"File Sharing": {
		"qBittorrent",
	},
	"Compression": {
		"Winrar x64",
		"Winrar x86",
		"7-Zip",
		"PeaZip",
	},
	"Developer Tools": {
		"WinSCP",
		"Notepad++",
		"PuTTY",
		"Eclipse",
		"FileZilla",
		"Visual Studio Code",
		"Visual Studio",
		"Sublime Text",
		"Python",
		"GoLang",
	},
	"Imaging": {
		"Krita",
		"Blender",
		"GIMP",
		"Inkscape",
	},
	"Security": {
		"AVG AntiVirus Free",
		"Avast Free Antivirus",
	},
	"Office": {
		"OpenOffice",
		"OnlyOffice",
		"LibreOffice",
		"Foxit PDF Reader",
	},
	"Utilities": {
		"Anyburn",
		"MiniTool Partition Wizard",
		"BalenaEtcher (Installer)",
		"BalenaEtcher (Portable)",
		"Rufus",
	},
	"Gaming": {
		"Steam",
		"EpicGames",
		"Rockstar Games Launcher",
		"EA App",
		"Ubisoft",
	},
	"Windows Tweaks": {},
}

var applications = map[string]string{
	// Browsers
	"Chromium":             "https://github.com/Hibbiki/chromium-win64/releases/download/v124.0.6367.79-r1274542/mini_installer.sync.exe",
	"Brave":                "https://laptop-updates.brave.com/latest/winx64",
	"Vivaldi":              "https://downloads.vivaldi.com/stable/Vivaldi.3.8.2259.37.exe",
	"Microsoft Edge":       "https://go.microsoft.com/fwlink/?linkid=2108834&Channel=Stable&language=en&brand=M100",
	"Opera":                "https://net.geo.opera.com/opera/stable/windows?utm_tryagain=yes&utm_source=(direct)&utm_medium=doc&utm_campaign=(direct)&http_referrer=missing&utm_site=opera_com&&utm_lastpage=opera.com/",
	"Firefox":              "https://download.mozilla.org/?product=firefox-latest-ssl&os=win64&lang=en-US&_gl=1*12wy6k4*_ga*NzE1Njk2OTQ0LjE3MDk1NDU0NTQ.*_ga_MQ7767QQQW*MTcxMTM3MDk2MC4zLjEuMTcxMTM3MDk5Mi4wLjAuMA..",
	"LibreWolf":            "https://gitlab.com/api/v4/projects/44042130/packages/generic/librewolf/125.0.2-1/librewolf-125.0.2-1-windows-x86_64-setup.exe",
	"Waterfox":             "https://cdn1.waterfox.net/waterfox/releases/G6.0.11/WINNT_x86_64/Waterfox%20Setup%20G6.0.11.exe",
	"Floorp Browser":       "https://github.com/Floorp-Projects/Floorp/releases/download/v11.12.1/floorp-stub.installer.exe",
	"Tor Browser":          "https://www.torproject.org/dist/torbrowser/13.0.13/tor-browser-windows-x86_64-portable-13.0.13.exe",
	"Avast Secure Browser": "https://www.avast.com/download-thank-you.php?product=ASB&locale=en-ww&direct=1",
	"Thorium":              "https://github.com/Alex313031/Thorium-Win/releases/download/M123.0.6312.133/thorium_AVX2_mini_installer.exe",
	"Mercury":              "https://github.com/Alex313031/Mercury/releases/download/v.123.0.1/mercury_123.0.1_win64_AVX2_installer.exe",

	// Messaging
	"Discord":         "https://discord.com/api/download?platform=win",
	"Zoom":            "https://zoom.us/client/latest/ZoomInstaller.exe",
	"Skype":           "https://get.skype.com/go/getskype-full",
	"Telegram":        "https://telegram.org/dl/desktop/win64",
	"Viber":           "https://download.cdn.viber.com/desktop/windows/ViberSetup.exe",
	"Microsoft Teams": "https://go.microsoft.com/fwlink/?linkid=2187217&clcid=0x409&culture=en-us&country=us",
	"Thunderbird":     "https://download.mozilla.org/?product=thunderbird-115.9.0-SSL&os=win64&lang=en-US",
	"Pidgin":          "https://downloads.sourceforge.net/project/pidgin/Pidgin/2.14.13/pidgin-2.14.13.exe?ts=gAAAAABmKRei1SaO3BfhlDAO_yy48-zrBAiIp0ALd6reVsHV_eJr72XeaiLxaEHx85yvdgoMJnlzAcBUIhJ57sOJ-6oWbzbEKQ%3D%3D&r=",

	// Media
	"VLC Media Player": "https://get.videolan.org/vlc/3.0.12/win64/vlc-3.0.12-win64.exe",
	"Spotify":          "https://download.scdn.co/SpotifyFullSetup.exe",
	"Audacity":         "https://github.com/audacity/audacity/releases/download/Audacity-3.4.2/audacity-win-3.4.2-64bit.exe",
	"iTunes":           "https://www.apple.com/itunes/download/win64",
	"OBS":              "https://cdn-fastly.obsproject.com/downloads/OBS-Studio-30.1.1-Full-Installer-x64.exe",
	"Streamlabs":       "https://streamlabs.com/streamlabs-desktop/download?sdb=0",
	"Kdenlive":         "https://download.kde.org/stable/kdenlive/24.02/windows/kdenlive-24.02.1.exe",
	"Winamp":           "https://download.winamp.com/winamp/winamp_latest_full.exe",
	"foobar2000":       "https://www.foobar2000.org/files/foobar2000-x64_v2.1.4.exe",
	"GOM Player":       "https://cdn.gomlab.com/gretech/player/GOMPLAYERGLOBALSETUP_CHROME.EXE",

	// File Sharing
	"qBittorrent": "https://downloads.sourceforge.net/project/qbittorrent/qbittorrent-win32/qbittorrent-4.6.4/qbittorrent_4.6.4_x64_setup.exe?ts=gAAAAABmKRc7Y2knpGy-LjaP4-DSZ0q5I7HZ8Atwqe_XtQ4XAF7ALpbsZLabSQ7Nw2nEdp4SKg0zuBj32SqENsbzQ1D7pHQbpg%3D%3D&r=https%3A%2F%2Fsourceforge.net%2Fprojects%2Fqbittorrent%2Ffiles%2Flatest%2Fdownload",

	// Compression
	"Winrar x86": "https://www.rarlab.com/rar/wrar601.exe",
	"Winrar x64": "https://www.win-rar.com/fileadmin/winrar-versions/winrar/winrar-x64-700.exe",

	"7-Zip":  "https://www.7-zip.org/a/7z2106-x64.exe",
	"PeaZip": "https://osdn.net/frs/redir.php?m=netix&f=peazip%2F7.10.0%2Fpeazip_portable-7.10.0.WINDOWS.exe",

	// Dev  Tools
	"WinSCP":             "https://winscp.net/download/files/202403251303916f4d129e8c7043e8537e47cb5d5f5b/WinSCP-6.3.2-Setup.exe",
	"Notepad++":          "https://github.com/notepad-plus-plus/notepad-plus-plus/releases/download/v8.6.4/npp.8.6.4.Installer.x64.exe",
	"PuTTY":              "https://the.earth.li/~sgtatham/putty/latest/w64/putty.exe",
	"FileZilla":          "https://download.filezilla-project.org/client/FileZilla_3.66.5_win64_sponsored2-setup.exe",
	"Eclipse":            "https://www.eclipse.org/downloads/download.php?file=/oomph/epp/2024-03/R/eclipse-inst-jre-win64.exe",
	"Visual Studio Code": "https://code.visualstudio.com/sha/download?build=stable&os=win32-x64-user",
	"Visual Studio":      "https://c2rsetup.officeapps.live.com/c2r/downloadVS.aspx?sku=community&channel=Release&version=VS2022&source=VSLandingPage&cid=2030:967bcd8b3df0dbedb8ebccd40b730d58",
	"Sublime Text":       "https://download.sublimetext.com/sublime_text_build_4169_x64_setup.exe",
	"Python":             "https://www.python.org/ftp/python/3.12.3/python-3.12.3-amd64.exe",
	"GoLang":             "https://go.dev/dl/go1.22.2.windows-amd64.msi",

	// Imaging
	"Krita":    "https://download.kde.org/stable/krita/5.2.2/krita-x64-5.2.2-setup.exe",
	"Blender":  "https://ftp.nluug.nl/pub/graphics/blender/release/Blender4.1/blender-4.1.1-windows-x64.msi",
	"GIMP":     "https://download.gimp.org/gimp/v2.10/windows/gimp-2.10.36-setup-1.exe",
	"Inkscape": "https://inkscape.org/gallery/item/44617/inkscape-1.3.2_2023-11-25_091e20e-x64.exe",

	// Security
	"AVG AntiVirus Free":   "https://bits.avcdn.net/productfamily_ANTIVIRUS/insttype_FREE/platform_WIN_AVG/installertype_ONLINE/build_RELEASE/cookie_mmm_bav_tst_007_402_a?ref=clid_1296527067.1711656977--seid_1711656976--senu_1&alt=en-ww",
	"Avast Free Antivirus": "https://www.avast.com/download-thank-you.php?product=FAV-PPC&locale=en-ww&direct=1",

	// Utilities
	"Anyburn":                   "https://www.anyburn.com/anyburn_setup_x64.exe",
	"MiniTool Partition Wizard": "https://cdn2.minitool.com/?p=pw&e=pw-free",
	"BalenaEtcher (Installer)":  "https://github.com/balena-io/etcher/releases/download/v1.18.11/balenaEtcher-Setup-1.18.11.exe",
	"BalenaEtcher (Portable)":   "https://github.com/balena-io/etcher/releases/download/v1.18.11/balenaEtcher-Portable-1.18.11.exe",
	"Rufus":                     "https://github.com/pbatard/rufus/releases/download/v4.4/rufus-4.4.exe",
	// Office
	"Foxit PDF Reader": "https://www.foxit.com/downloads/latest.html?product=Foxit-Reader&platform=Windows&version=&package_type=&language=English&distID=",
	"OpenOffice":       "https://sourceforge.net/projects/openofficeorg.mirror/files/4.1.15/binaries/en-US/Apache_OpenOffice_4.1.15_Win_x86_install_en-US.exe/download",
	"OnlyOffice":       "https://download.onlyoffice.com/install/desktop/editors/windows/distrib/onlyoffice/DesktopEditors_x64.exe",
	"LibreOffice":      "https://www.libreoffice.org/donate/dl/win-x86_64/24.2.2/ro/LibreOffice_24.2.2_Win_x86-64.msi",

	// Gaming
	"Ubisoft":                 "https://ubi.li/4vxt9",
	"Steam":                   "https://cdn.akamai.steamstatic.com/client/installer/SteamSetup.exe",
	"Rockstar Games Launcher": "https://gamedownloads.rockstargames.com/public/installer/Rockstar-Games-Launcher.exe",
	"EA App":                  "https://origin-a.akamaihd.net/EA-Desktop-Client-Download/installer-releases/EAappInstaller.exe",
	"EpicGames":               "https://launcher-public-service-prod06.ol.epicgames.com/launcher/api/installer/download/EpicGamesLauncherInstaller.msi",
}
