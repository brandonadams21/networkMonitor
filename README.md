# Network Monitoring Tool

This project is a network monitoring tool designed to capture and log network traffic data, providing insights into network activity for security and analysis purposes.

## Features

- Captures and logs network packets, including source and destination IP addresses, protocols, ports, and payload data.
- Supports human-readable payload data for easier analysis.
- JSON-encoded log entries for better readability and easier parsing.

## Installation

1. **Clone the Repository**:

    ```sh
    git clone https://github.com/yourusername/networkMonitor.git
    cd networkMonitor
    ```

2. **Install Dependencies**:

    Ensure you have Go installed. Then, install the required Go packages:

    ```sh
    go get github.com/google/gopacket
    go get github.com/google/gopacket/pcap
    ```

3. **Set Up Directory Structure**:

    Ensure your directory structure is set up as follows:

    ```
    networkMonitor/
    ├── cmd/
    │   └── monitor/
    │       └── main.go
    ├── pkg/
    │   ├── capture/
    │   │   └── capture.go
    │   ├── logger/
    │   │   └── logger.go
    │   ├── logparser/
    │       └── logparser.go
    ├── config/
    │   └── config.go
    ├── logs/
    │   └── network_traffic.log
    ├── scripts/
    │   └── run.sh
    ├── go.mod
    ├── go.sum
    ├── README.md
    ```

## Usage

### Running the Packet Capture

To start capturing network packets, ensure you are using the correct network interface name (e.g., `en0` for macOS). Then, run the script:

```sh
sudo ./scripts/run.sh
```

### Parsing the Log File

To parse and display the log entries, run the program with the `parse` argument:

```sh
./scripts/run.sh parse
```

### Viewing the Log File

The captured network traffic data will be logged to `logs/network_traffic.log`. You can view it using:

```sh
tail -f logs/network_traffic.log
```

## Resolving Git Issues

If you encounter the error "fatal: refusing to merge unrelated histories" when trying to pull from the remote repository, follow these steps:

### Allow Unrelated Histories

1. **Pull with the `--allow-unrelated-histories` Flag**:

    ```sh
    git pull origin main --allow-unrelated-histories
    ```

2. **Resolve Any Merge Conflicts**:

    After pulling the changes, you may encounter merge conflicts. Resolve these conflicts manually and then commit the changes.

3. **Commit the Merge**:

    ```sh
    git add .
    git commit -m "Merge unrelated histories"
    ```

### Alternative Approach

If you want to overwrite the local branch with the remote branch:

1. **Fetch the Remote Branch**:

    ```sh
    git fetch origin
    ```

2. **Reset the Local Branch**:

    ```sh
    git reset --hard origin/main
    ```

This command will discard all local changes and make the local branch identical to the remote branch.

