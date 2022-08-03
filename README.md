# podterm: a CLI podcast client

`podterm` is a CLI podcast client written in Go.

## installation

Install through go:
```shell
go install github.com/PerpetualCreativity/podterm@latest
```

## usage

Where a channel name is required, only a short portion of the channel is generally required. `podterm` fuzzy-searches for the right channel and uses the best match. `clear` and `remove` are exceptions, requiring an unambiguously specified channel name as they are destructive.

| command                          | behavior                                                                                     |
|----------------------------------|----------------------------------------------------------------------------------------------|
| `add url`                        | adds a channel from the RSS url                                                              |
| `play channel [n]`               | plays the latest episode from the channel, or the Nth episode in reverse-chronological order |
| `list [channel]`                 | lists all channels, or if a channel is specified, episodes in a channel.                     |
| `info channel [I]`               | get detailed information about episode I                                                     |
| `download channel I [J]`         | download episode I. or all episodes from I to J                                              |
| `refresh [channel]`              | refresh all channel feeds, or only the specified channel feed                                |
| `clear channel` OR `clear --all` | delete all downloaded episodes from a channel, or from all channels                          |
| `remove channel`                 | removes channel and all episodes downloaded from the channel                                 |
