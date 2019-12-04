# temper

[![GoDoc][godoc image]][godoc link]
[![License Badge][license image]][LICENSE.txt]

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
  - `$ go get github.com/gotmc/libusb`

## Documentation

Documentation can be found at either:

- <https://godoc.org/github.com/gotmc/temper>
- <http://localhost:6060/pkg/github.com/gotmc/temper/> after running `$
  godoc -http=:6060`

## Contributing

[temper][] is developed using [Scott Chacon][]'s [GitHub Flow][]. To
contribute, fork [temper][], create a feature branch, and then
submit a [pull request][].  [GitHub Flow][] is summarized as:

- Anything in the `master` branch is deployable
- To work on something new, create a descriptively named branch off of
  `master` (e.g., `new-oauth2-scopes`)
- Commit to that branch locally and regularly push your work to the same
  named branch on the server
- When you need feedback or help, or you think the branch is ready for
  merging, open a [pull request][].
- After someone else has reviewed and signed off on the feature, you can
  merge it into master.
- Once it is merged and pushed to `master`, you can and *should* deploy
  immediately.

## Testing

Prior to submitting a [pull request][], please run:

```bash
$ gofmt
$ golint
$ go vet
$ go test
```

To update and view the test coverage report:

```bash
$ go test -coverprofile coverage.out
$ go tool cover -html coverage.out
```

## License

[temper][] is released under the MIT license. Please see the
[LICENSE.txt][] file for more information.


## Prior Art

Below are projects written in other languages:

- <https://github.com/elpeo/rbtemper> --- Ruby TEMPer library
- <https://github.com/bitplane/temper> --- Command line sensor logger
  for Temper1 devices
- <https://github.com/padelt/temper-python> --- libusb/PyUSB-based
  driver to read TEMPer USB HID devices (USB ID 0x0C45:0x7401) and serve
  as a NewSNMP passpersist module


[GitHub Flow]: http://scottchacon.com/2011/08/31/github-flow.html
[godoc image]: https://godoc.org/github.com/gotmc/temper?status.svg
[godoc link]: https://godoc.org/github.com/gotmc/temper
[libusb]: https://github.com/gotmc/libusb
[libusb-c]: http://libusb.info
[LICENSE.txt]: https://github.com/gotmc/temper/blob/master/LICENSE.txt
[license image]: https://img.shields.io/badge/license-MIT-blue.svg
[pull request]: https://help.github.com/articles/using-pull-requests
[Scott Chacon]: http://scottchacon.com/about.html
[temper]: https://github.com/gotmc/temper
