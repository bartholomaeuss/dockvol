## dockvol

`dockvol` is a tiny Go-powered CLI meant to help with Docker volume hygiene before you kick off a backup. The current `backup` command interrogates Docker and makes sure a named volume is only attached to a single container, which prevents you from snapshotting a volume that is concurrently mounted by multiple services.

This repository is intentionally minimal so you can expand the command with your own backup logic (tarballing, remote sync, etc.) once volume safety has been verified.

---

### Features

- Lightweight Cobra-based CLI (`dockvol backup`) that shells out to Docker
- Verifies a volume is not attached to more than one container before backup
- Easy to extend with whatever archiving or copy commands match your stack

---

### Prerequisites

- Docker CLI available on your PATH
- Go 1.22+ if you want to build/run from source
- Access to the Docker daemon (local socket or remote context)

---

### Building / Running

```bash
git clone https://github.com/your-org/dockvol.git
cd dockvol
go run ./main.go backup --volume my-volume
```

To install as a binary:

```bash
go build -o dockvol .
./dockvol backup --volume my-volume
```

---

### Command Reference

| Command | Description | Flags |
| ------- | ----------- | ----- |
| `dockvol backup` | Checks if the provided Docker volume is attached to more than one container. Extend this command with your backup logic after the safety check passes. | `--volume, -v` (required) Name of the Docker volume to verify |

The command exits with a non-zero status if `docker ps -aq --filter volume=<name>` returns more than one container ID, which protects you from archiving a live, shared volume.

---

### Example: Alpine Container + Named Volume

The snippet below shows a realistic local workflow that demonstrates Docker volume handling and how `dockvol` fits in right before a backup step.

```bash
# 1. Create a demo volume
docker volume create demo-data

# 2. Populate it using a one-off Alpine container
docker run --rm \
  -v demo-data:/data \
  alpine:3.19 \
  sh -c 'echo "Hello from Alpine" > /data/message.txt'

# 3. Confirm file exists (optional)
docker run --rm \
  -v demo-data:/data \
  alpine:3.19 \
  cat /data/message.txt

# 4. Run dockvol to ensure only one container is attached
go run ./main.go backup --volume demo-data

# If you built a binary, run:
# ./dockvol backup --volume demo-data
```

At step 4, `dockvol` prints the container IDs that reference the volume and fails fast if multiple IDs are detected. Hook your preferred backup workflow (tar + gzip, `docker run --rm --volumes-from`, `restic`, etc.) after this safety gate.

---

### Extending

- Replace the placeholder backup logic in `cmd/backup.go` with commands that create a compressed archive of the volume contents.
- Add additional flags (e.g., `--destination`, `--compressor`, `--dry-run`) via Cobra to make the CLI production-ready.
- Wire it into CI/CD or scheduled jobs so stale volumes are automatically validated before snapshots are taken.

---

### License

MIT — see `LICENSE` for details.
