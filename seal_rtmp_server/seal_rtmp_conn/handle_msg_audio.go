package seal_rtmp_conn

import (
	"UtilsTools/identify_panic"
	"log"
)

func (rtmp *RtmpConn) handleMsgAudio(msg *MessageStream) (err error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err, ",panic at,", identify_panic.IdentifyPanic())
		}
	}()

	//log.Println("recv audio data, timestamp=", msg.header.timestamp)

	return
}
