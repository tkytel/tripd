# tripd

TRIP daemon - Telephony Routing Information Provider

## Requirements

* The platform should be linux/amd64, linux/arm/v7, linux/arm/v8, or linux/386.
* This software attempts to send an "unprivileged" ping via UDP.
    * On Linux, this must be enabled with the following `sysctl` command.
        ```
        sudo sysctl -w net.ipv4.ping_group_range="0 2147483647"
        ```

## License

MIT