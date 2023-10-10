# freemium-engine
Engine for creating Freemium applications with sales automation and a serverless model

## dependencies

1. Go 1.8+
2. Wails (GUI lib)
3. NPM 9.5.1+

## how to import private lib to another project

add this to `.gitconfig` file:

```bash
git config \
 --global \
 url."https://${user}:${personal_access_token}@github.com".insteadOf \
 "https://github.com"
```

and set `user` & `personal_access_token` in environment variables

## TODO

1. Come up with something with the private key so that it is not in the code. Maybe it can be generated and stored somehow based on the data of the current device.
2. add crypto payment lib
3. build app engine with logic module

## TODO ideas

1. Add collection of email addresses via a P2P network, to which the developer will have access through asymmetric encryption.

Install libwebkitgtk on Ubuntu

```bash
sudo apt-get update
sudo apt-get install -qq software-properties-common
sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 3B4FE6ACC0B21F32
sudo add-apt-repository 'deb [trusted=yes] http://cz.archive.ubuntu.com/ubuntu bionic main universe'
sudo apt-get update
sudo apt-get install -qq libwebkitgtk-1.0-0
```

install wails

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

to check dependencies:

```bash
wails doctor
```

```
# Wails Dependencies
┌──────────────────────────────────────────────────────────────────────────┐
| Dependency | Package Name          | Status    | Version                 |
| *docker    | docker.io             | Installed | 20.10.18                |
| gcc        | build-essential       | Installed | 12.8ubuntu1             |
| libgtk-3   | libgtk-3-dev          | Available | 3.24.18-1ubuntu1        |
| libwebkit  | libwebkit2gtk-4.0-dev | Available | 2.38.6-0ubuntu0.20.04.1 |
| npm        | npm                   | Installed | 9.5.1                   |
| *nsis      | nsis                  | Available | 3.05-2                  |
| pkg-config | pkg-config            | Installed | 0.29.1-0ubuntu4         |
└──────────────────────── * - Optional Dependency ─────────────────────────┘
```

install another libs:

```bash
sudo aptitude install libwebkitgtk-3.0-dev
sudo aptitude install libgtk-3-dev
```

If `libgtk-3-dev` will not be installed, then skip the first sentence of `aptitude` and accept the second with downgrade of the associated libraries.

## run dev version

```bash
wails dev
```

## build for linux

```bash
wails build -platform linux/amd64
```

## build for windows

```bash
wails build -platform windows/amd64
```
