# mcp-partner-desktop
the desktop application for [mcp-partner](https://github.com/Ericwyn/mcp-partner)

use wails to package


## Dev And Build

### Clone project
```shell
git clone https://github.com/Ericwyn/mcp-partner-desktop.git
```

### Clone mcp-partner project
```shell
cd mcp-partner-desktop
git clone https://github.com/Ericwyn/mcp-partner.git
```

### Build
- build
```shell
wails build
```

- build and run in linux
```shell
wails build && build/bin/mcp-partner
```

build result:
- windows: `mcp-partner.exe`
- mac: `mcp-partner.app`
- linux: `mcp-partner`

