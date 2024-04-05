# temper

[![GoDoc][godoc badge]][godoc link]
[![Go Report Card][report badge]][report card]
[![License Badge][license badge]][LICENSE.txt]

Go library to communicate with a TEMPer USB temperature monitors.

## Installation

```bash
$ go get github.com/gotmc/temper
```

## Dependencies

- [libusb C library][libusb-c] — Library for USB device access
  - OS X: `$ brew install libusb`
  - Debian/Ubuntu: `$ sudo apt-get install -y libusb-1.0-0 libusb-1.0-0-dev`
- [Go libusb][libusb] — Go bindings for the [libusb C library][libusb-c]
  - Add `require github.com/gotmc/libusb v1.0.22` to your `go.mod`

## Documentation

Documentation can be found at either:

- <https://godoc.org/github.com/gotmc/temper>
- <http://localhost:6060/pkg/github.com/gotmc/temper/> after running `$
godoc -http=:6060`

## Contributing

Contributions are welcome! To contribute please:

1. Fork the repository
2. Create a feature branch
3. Code
4. Submit a [pull request][]

### Testing

Prior to submitting a [pull request][], please run:

```bash
$ make check
$ make lint
```

To update and view the test coverage report:

```bash
$ make cover
```

## Prior Art

Below are projects written in other languages:

- <https://github.com/elpeo/rbtemper> — Ruby TEMPer library
- <https://github.com/bitplane/temper> — Command line sensor logger
  for Temper1 devices
- <https://github.com/padelt/temper-python> — libusb/PyUSB-based
  driver to read TEMPer USB HID devices (USB ID 0x0C45:0x7401) and serve
  as a NewSNMP passpersist module

## License

[temper][] is released under the MIT license. Please see the
[LICENSE.txt][] file for more information.

[godoc badge]: https://godoc.org/github.com/gotmc/temper?status.svg
[godoc link]: https://godoc.org/github.com/gotmc/temper
[libusb]: https://github.com/gotmc/libusb
[libusb-c]: http://libusb.info
[LICENSE.txt]: https://github.com/gotmc/temper/blob/master/LICENSE.txt
[license badge]: https://img.shields.io/badge/license-MIT-blue.svg
[pull request]: https://help.github.com/articles/using-pull-requests
[report badge]: https://goreportcard.com/badge/github.com/gotmc/temper
[report card]: https://goreportcard.com/report/github.com/gotmc/temper
[temper]: https://github.com/gotmc/temper
