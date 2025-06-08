# tripd

TRIP daemon - Telephony Routing Information Provider

## Requirements

* The platform should be linux/amd64, linux/arm/v7, linux/arm/v8, or linux/386.

## Setup

* This software attempts to send an "unprivileged" ping via UDP
    * On Linux, this must be enabled with the following `sysctl` command.
        ```sh
        $ sudo sysctl -w net.ipv4.ping_group_range="0 2147483647"
        ```
* Create `config.toml` by referring to `config.example.toml`
    * Place it in the same directory as `compose.yaml`.
        ```sh
        $ tree
        .
        ├── config.toml
        └── compose.yaml
        ```
* Fire it up
    ```
    $ docker compose up -d
    ```
* You can also expose tripd service to the Internet via tunneling service (e.g. Cloudflare Tunnel)

## Update

```sh
$ docker compose pull
$ docker compose up -d
```

## License

MIT