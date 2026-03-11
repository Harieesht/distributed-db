# DistributedDB

A distributed relational database engine built from scratch in Go.

> Work in progress. Currently in Phase 1 — Storage Foundation.

---

## Phase 1 — Storage Foundation

Building the core storage engine. No libraries. Just bytes on disk.

- [x] Page read / write (DiskManager)
- [x] Slotted page layout
- [ ] Row serialization
- [ ] B-tree insert and search
- [ ] Buffer pool manager
- [ ] Write-Ahead Log (WAL)

---

## Run

```bash
go run cmd/minidb/main.go
```

## Test

```bash
go test ./...
```