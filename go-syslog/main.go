package main

import (
	"flag"
	"fmt"
	"log"
	"log/syslog"
	"os"
	"strings"
	"time"
)

// 应用程序可以通过 UNIX domain sockets, UDP or TCP，向syslog守护进程发送日志。syslog守护进程可以在远端。 这样，就可以不用单独收集应用程序的日志了。
// golang提供了syslog 包，只需要调用Dial()，就可以连接syslog服务器，然后发送消息。
// 在写失败的情况下，syslog client会尝试重连syslog服务器，并重写。

// Dial定义如下

// func Dial(network, raddr string, priority Priority, tag string) (*Writer, error)

// Dial 建立一个到log守护进程的连接，这个守护进程的地址由参数raddr指定，连接方式由参数network定义(可以为tcp, udp，或空)，如果network是空，Dial会连接到本地的syslog服务器。

func main() {
	pNetwork := flag.String("network", "udp", "network, tcp/udp/(empty)")
	pAddr := flag.String("addr", "", "remote address, if network is tcp or upd, the empty address will default to :514")
	pPri := flag.String("pri", "local3.info", "syslog facility and severity (priority)")
	pTag := flag.String("tag", os.Args[0], "syslog tag, default os.Args[0]")
	flag.Parse()

	if *pAddr == "" {
		switch *pNetwork {
		case "udp", "tcp":
			*pAddr = ":514"
		}
	}

	sysLog, e := syslog.Dial(*pNetwork, *pAddr, parsePri(*pPri), *pTag)
	if e != nil {
		log.Fatal(e)
	}

	log.Printf("PID: %d", os.Getpid())

	for i := 0; ; i++ {
		// normal logs
		sysLog.Info(fmt.Sprintf("This is a test!, current iterater: %d", i))

		// you can even send JSON like messages like this:
		sysLog.Info(`@cee:{"key1":"value1", "key2":"value2"}`)

		// more details about JSON CEE format in rsyslog:
		// http://www.rsyslog.com/json-elasticsearch/

		time.Sleep(3 * time.Second)
	}
}

func parsePri(pri string) syslog.Priority {
	p := strings.IndexByte(pri, '.')
	facility, severity := pri[:p], pri[p+1:]

	s := strings.ToLower(severity)
	var facilityLevel, severityLevel uint8
	for k, v := range SeverityLevelsShort {
		if strings.HasPrefix(s, v) {
			severityLevel = k
			break
		}
	}

	s = strings.ToLower(facility)
	for k, v := range FacilityKeywords {
		if strings.HasPrefix(s, v) {
			facilityLevel = k
			break
		}
	}
	return syslog.Priority(facilityLevel*8 + severityLevel)

}

// https://github.com/influxdata/go-syslog/blob/develop/common/facility.go

// Facility maps facility numeric codes to facility string messages.
// As per RFC 5427.
var Facility = map[uint8]string{
	0:  "kernel messages",
	1:  "user-level messages",
	2:  "mail system messages",
	3:  "system daemons messages",
	4:  "authorization messages",
	5:  "messages generated internally by syslogd",
	6:  "line printer subsystem messages",
	7:  "network news subsystem messages",
	8:  "UUCP subsystem messages",
	9:  "clock daemon messages",
	10: "security/authorization messages",
	11: "ftp daemon messages",
	12: "NTP subsystem messages",
	13: "audit messages",
	14: "console messages",
	15: "clock daemon messages",
	16: "local use 0 (local0)",
	17: "local use 1 (local1)",
	18: "local use 2 (local2)",
	19: "local use 3 (local3)",
	20: "local use 4 (local4)",
	21: "local use 5 (local5)",
	22: "local use 6 (local6)",
	23: "local use 7 (local7)",
}

// FacilityKeywords maps facility numeric codes to facility keywords.
// As per RFC 5427.
var FacilityKeywords = map[uint8]string{
	0:  "kern",
	1:  "user",
	2:  "mail",
	3:  "daemon",
	4:  "auth",
	5:  "syslog",
	6:  "lpr",
	7:  "news",
	8:  "uucp",
	9:  "cron",
	10: "authpriv",
	11: "ftp",
	12: "ntp",
	13: "security",
	14: "console",
	15: "cron2",
	16: "local0",
	17: "local1",
	18: "local2",
	19: "local3",
	20: "local4",
	21: "local5",
	22: "local6",
	23: "local7",
}

// SeverityMessages maps severity levels to severity string messages.
var SeverityMessages = map[uint8]string{
	0: "system is unusable",
	1: "action must be taken immediately",
	2: "critical conditions",
	3: "error conditions",
	4: "warning conditions",
	5: "normal but significant condition",
	6: "informational messages",
	7: "debug-level messages",
}

// SeverityLevels maps seveirty levels to severity string levels.
var SeverityLevels = map[uint8]string{
	0: "emergency",
	1: "alert",
	2: "critical",
	3: "error",
	4: "warning",
	5: "notice",
	6: "informational",
	7: "debug",
}

// SeverityLevelsShort maps severity levels to severity short string levels
// as per https://github.com/torvalds/linux/blob/master/tools/include/linux/kern_levels.h and syslog(3).
var SeverityLevelsShort = map[uint8]string{
	0: "emerg",
	1: "alert",
	2: "crit",
	3: "err",
	4: "warning",
	5: "notice",
	6: "info",
	7: "debug",
}
