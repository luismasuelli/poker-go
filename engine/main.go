package engine

// All the games have at least one table. A typical game will have
// few tables, at least 1. Multi-table tournaments can have thousands
// of tables instead. TableID 0xffffffff is reserved for lobby in MTT.
// TableID 0 means "no table in particular".
type TableID uint32

// Tables will have a small number of seats. SeatID 0 means "no seat
// in particular".
type SeatID uint8

// Player requests will be big numbers to enumerate transactions that
// can be tracked in future messages. RequestID 0 means "no request
// in particular" and should never be used, to avoid confusion.
type RequestID uint64
