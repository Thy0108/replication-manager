// replication-manager - Replication Manager Monitoring and CLI for MariaDB and MySQL
// Authors: Guillaume Lefranc <guillaume@signal18.io>
//          Stephane Varoqui  <stephane@mariadb.com>
// This source code is licensed under the GNU General Public License, version 3.

package receiver

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tanji/replication-manager/graphite/logging"
	"github.com/tanji/replication-manager/graphite/points"
)

func TestPickle(t *testing.T) {
	// > python
	// >>> import pickle, struct
	// >>> listOfMetricTuples = [("hello.world", (1452200952, 42))]
	// >>> payload = pickle.dumps(listOfMetricTuples, protocol=2)
	// >>> header = struct.pack("!L", len(payload))
	// >>> message = header + payload
	// >>> print repr(message)
	// '\x00\x00\x00#\x80\x02]q\x00U\x0bhello.worldq\x01J\xf8\xd3\x8eVK*\x86q\x02\x86q\x03a.'

	test := newTCPTestCase(t, true)
	defer test.Finish()

	test.Send("\x00\x00\x00#\x80\x02]q\x00U\x0bhello.worldq\x01J\xf8\xd3\x8eVK*\x86q\x02\x86q\x03a.")

	time.Sleep(10 * time.Millisecond)

	select {
	case msg := <-test.rcvChan:
		test.Eq(msg, points.OnePoint("hello.world", 42, 1452200952))
	default:
		t.Fatalf("Message #0 not received")
	}
}

func TestBadPickle(t *testing.T) {
	assert := assert.New(t)
	test := newTCPTestCase(t, true)
	defer test.Finish()

	logging.Test(func(log logging.TestOut) {
		test.Send("\x00\x00\x00#\x80\x02]q\x00q\x0bhello.worldq\x01Rixf8\xd3\x8eVK*\x86q\x02\x86q\x03a.")
		time.Sleep(10 * time.Millisecond)
		assert.Contains(log.String(), "I [pickle] Can't unpickle message")
	})
}

// https://github.com/tanji/replication-manager/graphite/issues/30
func TestPickleMemoryError(t *testing.T) {
	assert := assert.New(t)
	test := newTCPTestCase(t, true)
	defer test.Finish()

	logging.Test(func(log logging.TestOut) {
		test.Send("\x80\x00\x00\x01") // 2Gb message length
		time.Sleep(10 * time.Millisecond)

		assert.Contains(log.String(), "W [pickle] Bad message")
	})
}
