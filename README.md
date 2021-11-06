> This is a fork of [github.com/matryer/is](https://github.com/matryer/is) that
> prints failures using t.Log instead of printing directly to os.Stderr so that
> failures are grouped properly with sub-test headers in `go test` output.
>
> Happy to submit this as a PR, just maintaining a fork for now since I'm not
> sure it'd even be of interest, and the API is tiny enough that I'm not
> worried about maintaining it.
