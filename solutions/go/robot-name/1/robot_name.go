package robotname

import (
	cryptoRand "crypto/rand"
	"errors"
	"sync"
)

const (
	maxRobotNames = 26 * 26 * 1000 // 676,000
)

var (
	mu        sync.Mutex
	allocated = make(map[string]struct{})
)

// Robot represents a factory robot.
type Robot struct {
	name string
}

// Name returns the robot's name. It assigns a new unique random name the first time,
// and reuses the same name on subsequent calls until Reset is called.
func (r *Robot) Name() (string, error) {
	mu.Lock()
	if r.name != "" {
		name := r.name
		mu.Unlock()
		return name, nil
	}

	if len(allocated) >= maxRobotNames {
		mu.Unlock()
		return "", errors.New("no names available: exhausted all 676,000 names")
	}

	// Keep trying until we find a free name.
	for {
		mu.Unlock() // generate without holding the lock
		n, err := randomName()
		if err != nil {
			return "", err
		}

		mu.Lock()
		if _, taken := allocated[n]; !taken {
			allocated[n] = struct{}{}
			r.name = n
			mu.Unlock()
			return n, nil
		}

		// If taken, loop to try another name. Also check exhaustion quickly.
		if len(allocated) >= maxRobotNames {
			mu.Unlock()
			return "", errors.New("no names available: exhausted all 676,000 names")
		}
	}
}

// Reset wipes the robot's current name so it can be reassigned on the next Name() call.
func (r *Robot) Reset() {
	mu.Lock()
	if r.name != "" {
		delete(allocated, r.name)
		r.name = ""
	}
	mu.Unlock()
}

// randomName generates "LLDDD" (two uppercase letters + three digits) using CSPRNG.
func randomName() (string, error) {
	buf := make([]byte, 5)
	if _, err := cryptoRand.Read(buf); err != nil {
		return "", err
	}
	name := []byte{
		'A' + (buf[0] % 26),
		'A' + (buf[1] % 26),
		'0' + (buf[2] % 10),
		'0' + (buf[3] % 10),
		'0' + (buf[4] % 10),
	}
	return string(name), nil
}