<!-- PROJECT LOGO -->
<h1 align="center">Font Table</h1>
<p align="center">
  Package for transcoding font table files used in Diablo II.
  <br />
  <br />
  <a href="https://github.com/gravestench/font_table/issues">Report Bug</a>
  ·
  <a href="https://github.com/gravestench/font_table/issues">Request Feature</a>
</p>

<!-- ABOUT THE PROJECT -->
## About

The font table transcoder for Diablo II is a Go package that provides 
functionality to decode and encode font table files used in the popular 
game Diablo II.

This package also includes both command-line and graphical applications, 
extending its utility and ease of use.

## Getting Started

### Prerequisites
Before using the applications found in `cmd/`, ensure that you have installed 
[Go 1.16][golang] or a later version, and that your Go environment is correctly 
set up. To install the applications successfully, you will also need to define 
`$GOBIN` and add it to your `$PATH` environment variable. Here's how you can 
do it:
```shell
export GOBIN=$HOME/.gobin
mkdir -p $GOBIN
PATH=$PATH:$GOBIN
```

### Installation
Once `$GOBIN` is defined and added to your `$PATH`, you can easily build and 
install all apps located in `cmd/` with the following commands:

```shell
# Clone the repository and navigate to the directory
git clone http://github.com/gravestench/font
cd font

# Build and install inside $GOBIN
go build ./cmd/...
go install ./cmd/...
```

With the installation completed, you should now be able to run the applications 
inside `cmd/` directly from the command-line, such as `font-view`.

<!-- CONTRIBUTING -->
## Contributing

The Font repository follows a similar project structure to other transcoder 
libraries. The `~/pkg/` directory houses the core font transcoder library, 
while `~/cmd/` includes subdirectories for each CLI/GUI application that can be 
compiled.

Contributions to the project are **highly appreciated**. If you wish to 
contribute, follow these steps:

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<!-- MARKDOWN LINKS & IMAGES -->
[font]: https://github.com/gravestench/font
[golang]: https://golang.org/dl/