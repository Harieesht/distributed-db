# Page I/O

The database is one file on disk divided into fixed 4096-byte chunks called pages.
Every read and write in the engine goes through a page.

Finding page N is just math — N × 4096. No scanning.

---

## How It Works

file.db

```

|<-- Page 0 -->|<-- Page 1 -->|<-- Page 2 -->|
0            4096           8192          12288

```

---

## Key Points

- os.Create truncates existing file — use intentionally

---

## Connects To

- **Slotted Page** — a page is just the []byte that slotted page reads and writes into
- **Buffer Pool** — caches pages in RAM so we don't hit disk on every call