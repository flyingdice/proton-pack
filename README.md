# proton-pack

Cross the streams.

![](assets/cross-the-streams.jpeg)

## Status

This is under active development. Don't use it.

## What is this?

A sandbox for playing around with the following ideas:

* Building re-usable streaming primitives in golang
  * Ideally implementing user extensibility in WASM
* Perhaps "Binary compatible" with Java Kafka Streams
  * Can run side-by-side with Java KStreams and distribute tasks between
* Playing with distributed checkpointing like Flink

## Usage

Run `make help` to see list of available commands.

```
❯ make help
modules                        Tidy up and vendor go modules.
help                           Print Makefile usage.
```

## License

[AGPL v3](LICENSE)
