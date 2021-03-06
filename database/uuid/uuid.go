// implementation from 	"github.com/google/uuid"
package uuid

import (
	"crypto/rand"
	"io"
)

//const (
//	lillian    = 2299160          // Julian day of 15 Oct 1582
//	unix       = 2440587          // Julian day of 1 Jan 1970
//	epoch      = unix - lillian   // Days between epochs
//	g1582      = epoch * 86400    // seconds between epochs
//	g1582ns100 = g1582 * 10000000 // 100s of a nanoseconds between epochs
//)
//
//var (
//	nodeMu     sync.Mutex
//	timeMu     sync.Mutex
//	nodeID     [6]byte // hardware for version 1 UUIDs
//	zeroID     [6]byte // nodeID with only 0's
//	clockSeq   uint16
//	lasttime   uint64
//	timeNow    = time.Now      // for testing
//	interfaces []net.Interface // cached list of interfaces
//	rander     = rand.Reader   // random function
//	ifname     string
//)
//
//// A UUID is a 128 bit (16 byte) Universal Unique Identifier as defined in RFC 4122.
type UUID []byte

//type UUID string

//
//// A Time represents a time as the number of 100's of nanoseconds since 15 Oct 1582.
//type Time int64
//
//// New returns a Version 1 UUID based on the current NodeID and clock
//// sequence, and the current time.  If the NodeID has not been set by SetNodeID
//// or SetNodeInterface then it will be set automatically.  If the NodeID cannot
//// be set NewUUID returns nil.  If clock sequence has not been set by
//// SetClockSequence then it will be set automatically.  If getTime fails to
//// return the current NewUUID returns nil and an error.
////
//// In most cases, New should be used.
//func New() (UUID, error) {
//	nodeMu.Lock()
//	if nodeID == zeroID {
//		setNodeInterface("")
//	}
//	nodeMu.Unlock()
//
//	var uuid UUID
//	now, seq, err := getTime()
//	if err != nil {
//		return uuid, err
//	}
//
//	timeLow := uint32(now & 0xffffffff)
//	timeMid := uint16((now >> 32) & 0xffff)
//	timeHi := uint16((now >> 48) & 0x0fff)
//	timeHi |= 0x1000 // Version 1
//
//	binary.BigEndian.PutUint32(uuid[0:], timeLow)
//	binary.BigEndian.PutUint16(uuid[4:], timeMid)
//	binary.BigEndian.PutUint16(uuid[6:], timeHi)
//	binary.BigEndian.PutUint16(uuid[8:], seq)
//	copy(uuid[10:], nodeID[:])
//
//	return uuid, nil
//}
//
//func setNodeInterface(name string) bool {
//	iname, addr := getHardwareInterface(name) // null implementation for js
//	if iname != "" && addr != nil {
//		ifname = iname
//		copy(nodeID[:], addr)
//		return true
//	}
//
//	// We found no interfaces with a valid hardware address.  If name
//	// does not specify a specific interface generate a random Node ID
//	// (section 4.1.6)
//	if name == "" {
//		ifname = "random"
//		randomBits(nodeID[:])
//		return true
//	}
//	return false
//}
//
//// getHardwareInterface returns the name and hardware address of interface name.
//// If name is "" then the name and hardware address of one of the system's
//// interfaces is returned.  If no interfaces are found (name does not exist or
//// there are no interfaces) then "", nil is returned.
////
//// Only addresses of at least 6 bytes are returned.
//func getHardwareInterface(name string) (string, []byte) {
//	if interfaces == nil {
//		var err error
//		interfaces, err = net.Interfaces()
//		if err != nil {
//			return "", nil
//		}
//	}
//	for _, ifs := range interfaces {
//		if len(ifs.HardwareAddr) >= 6 && (name == "" || name == ifs.Name) {
//			return ifs.Name, ifs.HardwareAddr
//		}
//	}
//	return "", nil
//}
//
//// getTime returns the current Time (100s of nanoseconds since 15 Oct 1582) and
//// clock sequence as well as adjusting the clock sequence as needed.  An error
//// is returned if the current time cannot be determined.
//func getTime() (Time, uint16, error) {
//	defer timeMu.Unlock()
//	timeMu.Lock()
//	t := timeNow()
//	// If we don't have a clock sequence already, set one.
//	if clockSeq == 0 {
//		setClockSequence(-1)
//	}
//	now := uint64(t.UnixNano()/100) + g1582ns100
//
//	// If time has gone backwards with this clock sequence then we
//	// increment the clock sequence
//	if now <= lasttime {
//		clockSeq = ((clockSeq + 1) & 0x3fff) | 0x8000
//	}
//	lasttime = now
//	return Time(now), clockSeq, nil
//}
//
//// randomBits completely fills slice b with random data.
//func randomBits(b []byte) {
//	if _, err := io.ReadFull(rander, b); err != nil {
//		panic(err.Error()) // rand should never fail
//	}
//}
//
//func setClockSequence(seq int) {
//	if seq == -1 {
//		var b [2]byte
//		randomBits(b[:]) // clock sequence
//		seq = int(b[0])<<8 | int(b[1])
//	}
//	oldSeq := clockSeq
//	clockSeq = uint16(seq&0x3fff) | 0x8000 // Set our variant
//	if oldSeq != clockSeq {
//		lasttime = 0
//	}
//}

// newUUID generates a random UUID according to RFC 4122
func New() (UUID, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return uuid, err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return uuid, nil
	//return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
