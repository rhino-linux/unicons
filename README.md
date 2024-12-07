# Unicons, the Unicorn Icon Pack
**Unicons**, also known as **the Unicorn Icon Pack**, is the Icon Pack designed for the [**Unicorn Desktop Environment**](https://rhinolinux.org/unicorn.html).

As Elsie (Lead Developer of Pacstall) says
> This is an icon pack, it provides icon pack.

## Development
To optimize SVGs and generate PNGs from scalable icons, first install the requirements:
```bash
sudo apt install nodejs golang-go librsvg2-bin # Debian and friends
brew install node go librsvg # macOS
```
```bash
npm install -g svgo # both do this
```
Then, run:
```bash
go run generate-sizes.go
```

## Installation
To install:
```bash
git clone https://github.com/rhino-linux/unicons
sudo mkdir -p /usr/share/icons/Unicons
sudo cp -r unicons/* /usr/share/icons/Unicons
sudo gtk-update-icon-cache /usr/share/icons/Unicons
```
Then (for Unicorn):
```
xfconf-query -c xsettings -p /Net/IconThemeName -s Unicons
gsettings set org.gnome.desktop.interface icon-theme Unicons
```

## Maintainers:
- oklopfer 
- AJStrong

## Contribute to Unicons
To start contributing to Unicons, check out the [Official Figma Templates](https://www.figma.com/community/file/1320453161902790267/unicons-template-kit) or the [SVG Templates](https://github.com/rhino-linux/unicons/tree/main/Templates).
