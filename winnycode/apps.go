package main

var categories = map[string][]string{
	"Browsers": {
		"Arc Browser",
		"Chromium",
		"Brave",
		"Vivaldi",
		"Microsoft Edge",
		"Opera",
		"Opera GX",
		"Firefox",
		"Firefox ESR",
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
		"hexchat",
		"Slack",
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
		"Filmora",
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
		"Visual Studio Code Insiders",
		"Visual Studio",
		"Sublime Text",
		"Python",
		"GoLang",
		"FreePascal",
	},
	"Imaging": {
		"Krita",
		"Blender",
		"GIMP",
		"Figma",
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
		"PDFsam Basic",
	},
	"Utilities": {
		"Anyburn",
		"MiniTool Partition Wizard",
		"BalenaEtcher (Installer)",
		"BalenaEtcher (Portable)",
		"Rufus",
		"Microsoft Powertoys",
		"WinDirStat",
		"WinMerge",
	},
	"Gaming": {
		"Steam",
		"EpicGames",
		"Rockstar Games Launcher",
		"EA App",
		"Ubisoft",
		"GeForce Experience",
	},
	"Windows Tweaks": {},
}

var applications = map[string]string{
	// Browsers
	"Arc Browser":          "https://releases.arc.net/windows/ArcInstaller.exe",
	"Chromium":             "https://github.com/Hibbiki/chromium-win64/releases/download/v125.0.6422.142-r1287751/mini_installer.sync.exe",
	"Brave":                "https://laptop-updates.brave.com/latest/winx64",
	"Vivaldi":              "https://downloads.vivaldi.com/stable/Vivaldi.6.7.3329.39.x64.exe",
	"Microsoft Edge":       "https://go.microsoft.com/fwlink/?linkid=2108834&Channel=Stable&language=en&brand=M100",
	"Opera":                "https://net.geo.opera.com/opera/stable/windows?utm_tryagain=yes&utm_source=(direct)&utm_medium=doc&utm_campaign=(direct)&http_referrer=missing&utm_site=opera_com&&utm_lastpage=opera.com/",
	"Firefox":              "https://download.mozilla.org/?product=firefox-latest-ssl&os=win64&lang=en-US&_gl=1*12wy6k4*_ga*NzE1Njk2OTQ0LjE3MDk1NDU0NTQ.*_ga_MQ7767QQQW*MTcxMTM3MDk2MC4zLjEuMTcxMTM3MDk5Mi4wLjAuMA..",
	"Firefox ESR":          "https://download.mozilla.org/?product=firefox-esr-latest-ssl&os=win64&lang=en-US",
	"LibreWolf":            "https://gitlab.com/api/v4/projects/44042130/packages/generic/librewolf/126.0-1/librewolf-126.0-1-windows-x86_64-setup.exe",
	"Waterfox":             "https://cdn1.waterfox.net/waterfox/releases/latest/windows",
	"Floorp Browser":       "https://github.com/Floorp-Projects/Floorp/releases/download/v11.13.3/floorp-win64.installer.exe",
	"Tor Browser":          "https://www.torproject.org/dist/torbrowser/13.0.13/tor-browser-windows-x86_64-portable-13.0.13.exe",
	"Avast Secure Browser": "https://www.avast.com/download-thank-you.php?product=ASB&locale=en-ww&direct=1",
	"Thorium":              "https://github.com/Alex313031/Thorium-Win/releases/download/M124.0.6367.218/thorium_AVX2_mini_installer.exe",
	"Mercury":              "https://github.com/Alex313031/Mercury/releases/download/v.123.0.1/mercury_123.0.1_win64_AVX2_installer.exe",

	// Messaging
	"Discord":         "https://discord.com/api/download?platform=win",
	"Zoom":            "https://zoom.us/client/latest/ZoomInstaller.exe",
	"Skype":           "https://get.skype.com/go/getskype-full",
	"Telegram":        "https://telegram.org/dl/desktop/win64",
	"Viber":           "https://download.cdn.viber.com/desktop/windows/ViberSetup.exe",
	"Microsoft Teams": "https://go.microsoft.com/fwlink/?linkid=2187217&clcid=0x409&culture=en-us&country=us",
	"Thunderbird":     "https://download.mozilla.org/?product=thunderbird-115.11.1-SSL&os=win64&lang=en-US",
	"Pidgin":          "https://downloads.sourceforge.net/project/pidgin/Pidgin/2.14.13/pidgin-2.14.13.exe?ts=gAAAAABmKRei1SaO3BfhlDAO_yy48-zrBAiIp0ALd6reVsHV_eJr72XeaiLxaEHx85yvdgoMJnlzAcBUIhJ57sOJ-6oWbzbEKQ%3D%3D&r=",
	"hexchat":         "https://github.com/hexchat/hexchat/releases/download/v2.16.2/HexChat.2.16.2.x64.exe",
	// Media
	"VLC Media Player": "https://get.videolan.org/vlc/3.0.20/win64/vlc-3.0.20-win64.exe",
	"Spotify":          "https://download.scdn.co/SpotifyFullSetup.exe",
	"Audacity":         "https://github.com/audacity/audacity/releases/download/Audacity-3.5.1/audacity-win-3.5.1-64bit.exe",
	"iTunes":           "https://www.apple.com/itunes/download/win64",
	"OBS":              "https://cdn-fastly.obsproject.com/downloads/OBS-Studio-30.1.2-Full-Installer-x64.exe",
	"Streamlabs":       "https://streamlabs.com/streamlabs-desktop/download?sdb=0",
	"Kdenlive":         "https://download.kde.org/stable/kdenlive/24.02/windows/kdenlive-24.02.1.exe",
	"Winamp":           "https://download.winamp.com/winamp/winamp_latest_full.exe",
	"foobar2000":       "https://www.foobar2000.org/files/foobar2000-x64_v2.1.4.exe",
	"GOM Player":       "https://cdn.gomlab.com/gretech/player/GOMPLAYERGLOBALSETUP_CHROME.EXE",
	"Filmora":          "https://download.wondershare.com/filmora_full846.exe",

	// File Sharing
	"qBittorrent": "https://downloads.sourceforge.net/project/qbittorrent/qbittorrent-win32/qbittorrent-4.6.4/qbittorrent_4.6.4_x64_setup.exe?ts=gAAAAABmKRc7Y2knpGy-LjaP4-DSZ0q5I7HZ8Atwqe_XtQ4XAF7ALpbsZLabSQ7Nw2nEdp4SKg0zuBj32SqENsbzQ1D7pHQbpg%3D%3D&r=https%3A%2F%2Fsourceforge.net%2Fprojects%2Fqbittorrent%2Ffiles%2Flatest%2Fdownload",

	// Compression
	"Winrar x86": "https://www.rarlab.com/rar/wrar601.exe",
	"Winrar x64": "https://www.win-rar.com/fileadmin/winrar-versions/winrar/winrar-x64-700.exe",

	"7-Zip":  "https://www.7-zip.org/a/7z2106-x64.exe",
	"PeaZip": "https://osdn.net/frs/redir.php?m=netix&f=peazip%2F7.10.0%2Fpeazip_portable-7.10.0.WINDOWS.exe",

	// Dev  Tools
	"WinSCP":                      "https://winscp.net/download/files/202403251303916f4d129e8c7043e8537e47cb5d5f5b/WinSCP-6.3.2-Setup.exe",
	"Notepad++":                   "https://github.com/notepad-plus-plus/notepad-plus-plus/releases/download/v8.6.4/npp.8.6.4.Installer.x64.exe",
	"PuTTY":                       "https://the.earth.li/~sgtatham/putty/latest/w64/putty.exe",
	"FileZilla":                   "https://download.filezilla-project.org/client/FileZilla_3.66.5_win64_sponsored2-setup.exe",
	"Eclipse":                     "https://www.eclipse.org/downloads/download.php?file=/oomph/epp/2024-03/R/eclipse-inst-jre-win64.exe",
	"Visual Studio Code":          "https://code.visualstudio.com/sha/download?build=stable&os=win32-x64-user",
	"Visual Studio Code Insiders": "https://code.visualstudio.com/sha/download?build=insider&os=win32-x64-user",
	"Visual Studio":               "https://c2rsetup.officeapps.live.com/c2r/downloadVS.aspx?sku=community&channel=Release&version=VS2022&source=VSLandingPage&cid=2030:967bcd8b3df0dbedb8ebccd40b730d58",
	"Sublime Text":                "https://download.sublimetext.com/sublime_text_build_4169_x64_setup.exe",
	"Python":                      "https://www.python.org/ftp/python/3.12.4/python-3.12.4-amd64.exe",
	"GoLang":                      "https://go.dev/dl/go1.22.4.windows-amd64.msi",
	"FreePascal":                  "https://sourceforge.net/projects/freepascal/files/latest/download",

	// Imaging
	"Figma":    "https://www.figma.com/download/desktop/win",
	"Krita":    "https://download.kde.org/stable/krita/5.2.2/krita-x64-5.2.2-setup.exe",
	"Blender":  "https://ftp.nluug.nl/pub/graphics/blender/release/Blender4.1/blender-4.1.1-windows-x64.msi",
	"GIMP":     "https://download.gimp.org/gimp/v2.10/windows/gimp-2.10.38-setup.exe",
	"Inkscape": "https://inkscape.org/gallery/item/44617/inkscape-1.3.2_2023-11-25_091e20e-x64.exe",

	// Security
	"AVG AntiVirus Free":   "https://bits.avcdn.net/productfamily_ANTIVIRUS/insttype_FREE/platform_WIN_AVG/installertype_ONLINE/build_RELEASE/cookie_mmm_bav_tst_007_402_a?ref=clid_1296527067.1711656977--seid_1711656976--senu_1&alt=en-ww",
	"Avast Free Antivirus": "https://www.avast.com/download-thank-you.php?product=FAV-PPC&locale=en-ww&direct=1",

	// Utilities
	"Anyburn":                   "https://www.anyburn.com/anyburn_setup_x64.exe",
	"MiniTool Partition Wizard": "https://cdn2.minitool.com/?p=pw&e=pw-free",
	"BalenaEtcher (Installer)":  "https://github.com/balena-io/etcher/releases/download/v1.19.21/balenaEtcher-1.19.21.Setup.exe",
	"BalenaEtcher (Portable)":   "https://github.com/balena-io/etcher/releases/download/v1.18.11/balenaEtcher-Portable-1.18.11.exe",
	"Rufus":                     "https://github.com/pbatard/rufus/releases/download/v4.5/rufus-4.5.exe",
	"Microsoft Powertoys":       "https://github.com/microsoft/PowerToys/releases/download/v0.81.1/PowerToysUserSetup-0.81.1-x64.exe",
	"Lively Wallpaper":          "https://github.com/rocksdanister/lively/releases/download/v2.1.0.6/lively_setup_x86_full_v2106.exe",
	"WinDirStat":                "https://downloads.sourceforge.net/project/windirstat/windirstat/1.1.2%20installer%20re-release%20%28more%20languages%21%29/windirstat1_1_2_setup.exe?ts=gAAAAABmZChj0YIjHo4SH5PRhMPA7eQd95BLYMq0ywIOzH3X1YKnuNgheGRK8crVLn1ymmGDCZUv0Js7RMGh_q8uIy6m_XFBww%3D%3D&r=https%3A%2F%2Fsourceforge.net%2Fprojects%2Fwindirstat%2Ffiles%2Fwindirstat%2F1.1.2%2520installer%2520re-release%2520%2528more%2520languages%2521%2529%2Fwindirstat1_1_2_setup.exe%2Fdownload%3Fuse_mirror%3Daltushost-swe",
	"WinMerge":                  "https://downloads.sourceforge.net/winmerge/WinMerge-2.16.40-x64-Setup.exe",

	// Office
	"Foxit PDF Reader": "https://www.foxit.com/downloads/latest.html?product=Foxit-Reader&platform=Windows&version=&package_type=&language=English&distID=",
	"OpenOffice":       "https://sourceforge.net/projects/openofficeorg.mirror/files/4.1.15/binaries/en-US/Apache_OpenOffice_4.1.15_Win_x86_install_en-US.exe/download",
	"OnlyOffice":       "https://download.onlyoffice.com/install/desktop/editors/windows/distrib/onlyoffice/DesktopEditors_x64.exe",
	"LibreOffice":      "https://www.libreoffice.org/donate/dl/win-x86_64/24.2.3/en-US/LibreOffice_24.2.3_Win_x86-64.msi",
	"PDFsam Basic":     "https://download7.pdfsam.org/get-app.aspx?configld=DECC253B-6466-450B-B829-296A7FBAB00A&uid=1007261&cmp=PDFsam_Basic&ref=pdfsam.org/in-app&wid=6848",

	// Gaming
	"Ubisoft":                 "https://ubi.li/4vxt9",
	"Steam":                   "https://cdn.akamai.steamstatic.com/client/installer/SteamSetup.exe",
	"Rockstar Games Launcher": "https://gamedownloads.rockstargames.com/public/installer/Rockstar-Games-Launcher.exe",
	"EA App":                  "https://origin-a.akamaihd.net/EA-Desktop-Client-Download/installer-releases/EAappInstaller.exe",
	"EpicGames":               "https://launcher-public-service-prod06.ol.epicgames.com/launcher/api/installer/download/EpicGamesLauncherInstaller.msi",
	"GeForce Experience":      "https://uk.download.nvidia.com/GFE/GFEClient/3.28.0.412/GeForce_Experience_v3.28.0.412.exe",
	"Slack":                   "https://slack.com/downloads/instructions/windows?ddl=1&build=win64",
}
