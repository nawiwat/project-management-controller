package logformatter

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

// ACNFormatter is a ACN standard type formatter
type ACNFormatter struct {
}

// Format to ACN standard logger
// %d{yyyy-MM-dd'T'HH:mm:ss.SSSXXX}|2|%-5level|acw-crypto-go-demo|,%X{clientIP},%X{transaction_id},%X{uuid}|%thread|main|%logger{0}|%msg%n
func (f *ACNFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var uuid string
	var ip string
	var thread string
	var logName string
	var transactionID string

	if entry.Data != nil {
		if ctxUUID, ok := entry.Data["uuid"].(string); ok {
			uuid = ctxUUID
		}

		if ctxIP, ok := entry.Data["ip"].(string); ok {
			ip = ctxIP
		}

		if ctxThread, ok := entry.Data["thread"].(string); ok {
			thread = ctxThread
		}

		if ctxLogName, ok := entry.Data["logName"].(string); ok {
			logName = ctxLogName
		}

		if ctxTransactionID, ok := entry.Data["transactionID"].(string); ok {
			transactionID = ctxTransactionID
		}
	}

	fields := fmt.Sprintf("%s|2|%s|project-management|,%s,%s,%s|%s|main|%s|%s\n",
		entry.Time.Format("2006-01-02T15:04:05.000Z07:00"),
		strings.ToUpper(entry.Level.String()),
		ip,
		transactionID,
		uuid,
		thread,
		logName,
		entry.Message)

	return []byte(fields), nil
}