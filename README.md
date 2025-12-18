# mcp-partner-desktop
the desktop application for [mcp-partner](https://github.com/mcp-partner/mcp-partner.github.io)

use wails to package


## Dev And Build

### Clone project
```shell
git clone https://github.com/mcp-partner/mcp-partner-desktop.git
```

### Clone mcp-partner project
```shell
cd mcp-partner-desktop
git clone https://github.com/mcp-partner/mcp-partner.github.io.git ./mcp-partner
rm -rf ./mcp-partner/node_modules ./mcp-partner/package-lock.json
```

### Build
- build with devtools
```shell
wails build -devtools
```

- build with devtools and run in linux
```shell
wails build -devtools && build/bin/mcp-partner
```

build result:
- windows: `mcp-partner.exe`
- mac: `mcp-partner.app`
- linux: `mcp-partner`


## Install
### Ubuntu
- install mcp-partner-desktop
    ```shell
    # get ~ path
    echo $HOME
    
    # after build mcp-partner-desktop
    # install mcp-partner-desktop
    cp build/bin/mcp-partner $HOME/.local/bin/mcp-partner
    
    # install mcp-partner.png
    cp mcp-partner/public/icon_256px.png $HOME/.local/share/icons/mcp-partner.png
    
    # install mcp-partner.desktop
    cp ./mcp-partner.desktop $HOME/.local/share/applications/
    ```
- remove mcp-partner-desktop
    ```shell
    # remove mcp-partner-desktop
    rm $HOME/.local/bin/mcp-partner
    
    # remove mcp-partner.png
    rm $HOME/.local/share/icons/mcp-partner.png
    
    # remove mcp-partner.desktop
    rm $HOME/.local/share/applications/mcp-partner.desktop
    ```
