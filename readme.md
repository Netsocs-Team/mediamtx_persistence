## Mediamtx Persistence

This is a docker image that will persist the mediamtx config to a file.

It will also update the mediamtx config when the file changes.

## Docker Compose

```yaml
services:
    mediamtx:
        image: ghcr.io/netsocs-team/mediamtx_with_persistence/main:latest

        ports:
            - "9997:9997"
            - "8888:8888"
            - "8889:8889"
            - "8189:8189"
            - "9996:9996"
        volumes:
            - ./recordings:/recordings:rw
            - ./mediamtx:/etc/mediamtx
```

## Usage

1. Create a directory for the recordings

```bash
mkdir recordings
```

2. Create a mediamtx config file

```yaml
mkdir mediamtx
```

3. Run the docker compose

```bash
docker compose up
```

Done!
