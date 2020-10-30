# tidy-photo

Organize your photo library according its taken date.

## Building
```sh
go build
```

## Usage
```sh
./tidy-photo [-s] [-o] [-t] [-p]
```

### Parameters
- **s**: The source path. Its default value is `./`.
- **o**: The output path. Its default value also is `./`.
- **t**: Test mode. If true only show the expected output. Its default is `false`.
- **p**: The output directory structure pattern. It follows the
Golang's [time/Time.Format](https://golang.org/pkg/time/#Time.Format) guide.

## Output
The default `p` value is `YYYY/YYYY-MM-DD` so the generated output will be this following structure:
```
YYYY
└── YYYY-MM-DD
    └── Filename.ext
```
Example:
```
2020
├── 2020-10-28
│   ├── DSC00001.ARW
│   └── DSC00002.ARW
└── 2020-10-29
    └── IMG_0001.CR3
```
