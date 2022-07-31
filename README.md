# podterm: a CLI podcast client

`podterm` is a CLI podcast client written in Go.

## installation

Install through go:
```shell
go install github.com/PerpetualCreativity/podterm@latest
```

## usage

| command                                            | supported? | behavior                                                                                                     |
|----------------------------------------------------|------------|--------------------------------------------------------------------------------------------------------------|
| `add url`                                          | yes        | adds a channel from the RSS feed                                                                             |
| `refresh title [-n, --number N] [-o, --overwrite]` | yes        | downloads the last N (default: 5) episodes from the channel, skipping already downloaded episodes by default |
| `refresh-all [-n, --number N]`                     | yes        | downloads the last N (default: 5) episodes from all channels                                                 |
| `list [channel]`                                   | yes        | lists all channels, or if channel is specified, all episodes in channel                                      |
| `play channel [n]`                                 | yes        | plays the latest episode from the channel, or the Nth episode in reverse-chronological order                 |
| `remove channel`                                   | yes        | removes channel and all episodes downloaded from the channel                                                 |
