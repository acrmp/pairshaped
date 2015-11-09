# pairshaped

A pre-commit hook that explodes unless there are two faces present.

## Installation

```
$ brew install homebrew/science/opencv
$ go get github.com/acrmp/pairshaped
```

Within the repo that you want to enforce pairing:

```
$ ln -s $(which pairshaped) .git/hooks/pre-commit
$ curl 'https://raw.githubusercontent.com/lazywei/go-opencv/master/samples/haarcascade_frontalface_alt.xml' > .git/hooks/haarcascade_frontalface_alt.xml
```

## Related

On a more serious note you may find git duet's `GIT_DUET_ROTATE_AUTHOR` useful if you
haven't come across it.
