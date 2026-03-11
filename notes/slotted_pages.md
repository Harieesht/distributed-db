# Slotted Page

A page is 4096 bytes. You need to store variable-length rows inside it, delete them safely, and give each row a stable ID that never changes even if the row moves. A slotted page solves all three by splitting the page into two sides that grow toward each other — slots from the left, rows from the right.

---

## How It Works

```
|header|slot0|slot1|←——— free space ———→|row1|row0|
0      6     10    14                 4070  4085  4096

slots grow →                        ← rows grow
free space = upper - lower
```

**Header** (6 bytes at byte 0)
```
byte 0-1  →  lower     (where next slot will be written)
byte 2-3  →  upper     (where next row will be written)
byte 4-5  →  numSlots  (total slots including dead ones)
```

**Slot** (4 bytes each)
```
byte 0-1  →  offset  (where the row starts in the page)
byte 2-3  →  length  (how many bytes the row is)

offset == 0 AND length == 0  →  dead slot (deleted row)
```

**Insert** — write row at upper, write slot at lower, update header, return numSlots as slotID.

**Get** — find slot at `6 + slotID*4`, read offset and length, return `page[offset:offset+length]`.

**Delete** — find slot at `6 + slotID*4`, write `{0, 0}`. Nothing moves.

---

## Key Points

- Slots grow right, rows grow left — free space is everything between them
- Page is full when `upper - lower < len(row) + 4`
- Deleting a row just marks the slot dead — bytes stay on disk until compaction
- Slot numbers never change — only the offset inside the slot entry changes
- SlotID returned from insert is `numSlots` before incrementing

---

## Connects To

- **Row Serialization** — the `[]byte` you pass into `InsertRow` is a serialized row
- **Tuple ID** — a TID is `(pageID, slotID)` — the slotID is what this page hands back on insert
- **Buffer Pool** — the page `[]byte` lives in the buffer pool in RAM; only flushed to disk when dirty