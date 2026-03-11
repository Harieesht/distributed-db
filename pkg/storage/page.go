package storage

import (
	"encoding/binary"
	"errors"
)

type page struct {
	bytes []byte
}

//initializing a slotted page with a header
func InitPage() *page {

	bytes := make([]byte,PageSize)
	binary.LittleEndian.PutUint16(bytes[0:2],6) // first two bytes stores the lower end
	binary.LittleEndian.PutUint16(bytes[2:4],4096) //next to bytes stores the upper end
	binary.LittleEndian.PutUint16(bytes[4:6],0) // next two records stores the nuber of sets

	return &page{bytes:bytes}
}

func (p *page) Insert(row []byte) (int, error) {
    lowerBound := binary.LittleEndian.Uint16(p.bytes[0:2])
    upperBound := binary.LittleEndian.Uint16(p.bytes[2:4])
    numSlots   := binary.LittleEndian.Uint16(p.bytes[4:6])

    if upperBound-lowerBound < uint16(len(row))+4 {
        return 0, errors.New("not enough space in the page")
    }

    newUpperBound := upperBound - uint16(len(row))

    // write row
    copy(p.bytes[newUpperBound:], row)

    // write slot entry
    binary.LittleEndian.PutUint16(p.bytes[lowerBound:lowerBound+2], newUpperBound)
    binary.LittleEndian.PutUint16(p.bytes[lowerBound+2:lowerBound+4], uint16(len(row)))

    // update header
    binary.LittleEndian.PutUint16(p.bytes[0:2], lowerBound+4)
    binary.LittleEndian.PutUint16(p.bytes[2:4], newUpperBound)
    binary.LittleEndian.PutUint16(p.bytes[4:6], numSlots+1)

    return int(numSlots), nil
}

func (p *page) Get(slotID int) ([]byte, error) {
    numSlots := int(binary.LittleEndian.Uint16(p.bytes[4:6]))

    if slotID >= numSlots {
        return nil, errors.New("slot id out of range")
    }

    slot := 6 + slotID*4
    rowOffset    := binary.LittleEndian.Uint16(p.bytes[slot : slot+2])
    rowLength    := binary.LittleEndian.Uint16(p.bytes[slot+2 : slot+4])

    if rowOffset == 0 && rowLength == 0 {
        return nil, errors.New("slot is deleted")
    }

    return p.bytes[rowOffset : rowOffset+rowLength], nil
}

func(p *page) Delete(slotID int) error {

	numSlots := binary.LittleEndian.Uint16(p.bytes[4:6])

	if slotID >= int(numSlots) {
		return errors.New("slot id out of range")
	}

	slot := 6+slotID*4

	// setting rowoffset and row length as 0 for future page compaction
	binary.LittleEndian.PutUint16(p.bytes[slot:slot+2],0)
	binary.LittleEndian.PutUint16(p.bytes[slot+2:slot+4],0)

	return nil
}
