package avr

import (
	"fmt"
	"net"
	"strings"
	"time"
)

// AnthemClient is a client for Anthem AVRs
type AnthemClient struct {
	ServerURL string
	Port      string
}

var anthemAudioFormats = map[string]string{
	"0": "no audio",
	"1": "analog",
	"2": "pcm",
	"3": "dolby",
	"4": "dsd",
	"5": "dts",
	"6": "atmos",
}

// makeReq sends a command to the Anthem AVR via TCP and returns the response
func (c *AnthemClient) makeReq(command string) (string, error) {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", c.ServerURL, c.Port), 5*time.Second)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	cmd := fmt.Sprintf("%s;", command)
	log.Debugf("Sending Anthem command: %s", cmd)

	_, err = conn.Write([]byte(cmd))
	if err != nil {
		return "", err
	}

	conn.SetReadDeadline(time.Now().Add(5 * time.Second))

	var result []byte
	buf := make([]byte, 1)
	for {
		n, err := conn.Read(buf)
		if n > 0 {
			if buf[0] == ';' {
				break
			}
			result = append(result, buf[0])
		}
		if err != nil {
			break
		}
	}

	log.Debugf("Got Anthem result: %s", string(result))
	return string(result), nil
}

// GetCodec returns the current audio input format from the Anthem AVR
func (c *AnthemClient) GetCodec() (string, error) {
	resp, err := c.makeReq("Z1AIF?")
	if err != nil {
		return "", err
	}
	// response format: Z1AIF<n>
	if len(resp) < 6 {
		return "", fmt.Errorf("unexpected Anthem response: %s", resp)
	}
	code := strings.TrimPrefix(resp, "Z1AIF")
	if format, ok := anthemAudioFormats[code]; ok {
		return format, nil
	}
	return strings.ToLower(code), nil
}
