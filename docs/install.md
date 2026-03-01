# Installation

## Homebrew (macOS / Linux)

```shell
brew install OpenSyntaxHQ/tweak/tweak
```

## Go

```shell
go install github.com/OpenSyntaxHQ/tweak@latest
```

## Quick Install (curl)

```shell
curl -sfL https://raw.githubusercontent.com/OpenSyntaxHQ/tweak/main/install.sh | sh
```

## Binary Downloads

### macOS (Universal)

```shell
curl -L https://github.com/OpenSyntaxHQ/tweak/releases/latest/download/tweak_Darwin_all.tar.gz | tar xz
sudo mv tweak /usr/local/bin/
```

### Linux

| Architecture | Link |
|---|---|
| amd64 | [tweak_Linux_x86_64.tar.gz](https://github.com/OpenSyntaxHQ/tweak/releases/latest/download/tweak_Linux_x86_64.tar.gz) |
| arm64 | [tweak_Linux_arm64.tar.gz](https://github.com/OpenSyntaxHQ/tweak/releases/latest/download/tweak_Linux_arm64.tar.gz) |
| i386 | [tweak_Linux_i386.tar.gz](https://github.com/OpenSyntaxHQ/tweak/releases/latest/download/tweak_Linux_i386.tar.gz) |

### Windows

| Architecture | Link |
|---|---|
| amd64 | [tweak_Windows_x86_64.zip](https://github.com/OpenSyntaxHQ/tweak/releases/latest/download/tweak_Windows_x86_64.zip) |
| arm64 | [tweak_Windows_arm64.zip](https://github.com/OpenSyntaxHQ/tweak/releases/latest/download/tweak_Windows_arm64.zip) |
| i386 | [tweak_Windows_i386.zip](https://github.com/OpenSyntaxHQ/tweak/releases/latest/download/tweak_Windows_i386.zip) |

### Linux Packages

=== "Debian / Ubuntu"
    ```shell
    dpkg -i tweak_*.deb
    ```

=== "RHEL / Fedora"
    ```shell
    rpm -i tweak_*.rpm
    ```

=== "Arch Linux"
    ```shell
    pacman -U tweak_*.pkg.tar.zst
    ```
