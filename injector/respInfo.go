package injector 

import (
	"time"
	"net/url"
	"strconv"
)

const (
	 RESPONSE_INFO_VERSION = 3

	 RESPONSE_INFO_MINIMUM_VERSION = 3
   
	 RESPONSE_INFO_VERSION_MASK = 0xFF
   
	 RESPONSE_INFO_HAS_CERT = 1 << 8
   
	 RESPONSE_INFO_HAS_SECURITY_BITS = 1 << 9
   
	 RESPONSE_INFO_HAS_CERT_STATUS = 1 << 10
   
	 RESPONSE_INFO_HAS_VARY_DATA = 1 << 11
   
	 RESPONSE_INFO_TRUNCATED = 1 << 12
   
	 RESPONSE_INFO_WAS_SPDY = 1 << 13
   
	 RESPONSE_INFO_WAS_ALPN = 1 << 14
   
	 RESPONSE_INFO_WAS_PROXY = 1 << 15
   
	 RESPONSE_INFO_HAS_SSL_CONNECTION_STATUS = 1 << 16
   
	 RESPONSE_INFO_HAS_ALPN_NEGOTIATED_PROTOCOL = 1 << 17
   
	 RESPONSE_INFO_HAS_CONNECTION_INFO = 1 << 18

	 RESPONSE_INFO_USE_HTTP_AUTHENTICATION = 1 << 19
   
	 RESPONSE_INFO_HAS_SIGNED_CERTIFICATE_TIMESTAMPS = 1 << 20
   
	 RESPONSE_INFO_UNUSED_SINCE_PREFETCH = 1 << 21

	 RESPONSE_INFO_HAS_KEY_EXCHANGE_GROUP = 1 << 22
   
	 RESPONSE_INFO_PKP_BYPASSED = 1 << 23
   
	 RESPONSE_INFO_HAS_STALENESS = 1 << 24
   
	 RESPONSE_INFO_HAS_PEER_SIGNATURE_ALGORITHM = 1 << 25
   
	 RESPONSE_INFO_RESTRICTED_PREFETCH = 1 << 26
   
	 RESPONSE_INFO_HAS_DNS_ALIASES = 1 << 27
)

const (
	kTimeTToMicrosecondsOffset = int64(11644473600000000)
)

func timeToInt64(t time.Time) int64 {
	micr := t.UnixNano() / int64(time.Microsecond)
	return micr + kTimeTToMicrosecondsOffset
}

func generateHeaderData(
	pickle *Pickle,
	contentSize int, 
	contentType string,
	currDate time.Time, 
	expiresDate time.Time,
) {
	headers := generateHttpHeaders(
		contentSize, 
		contentType,
		currDate, 
		expiresDate,
	)
	data := make([]byte, 0)
	data = append(data, []byte("HTTP/1.1 200 OK")...)
	data = append(data, 0x0)
	for _, header := range headers {
		data = append(data, []byte(header)...)
		data = append(data, 0x0)
	}
	data = append(data, 0x0)
	pickle.WriteInt32(int32(len(data)), false)
	pickle.WriteBytes(data)
}

func (self *ChromeInjector) persistRespInfo(
	pageUrl *url.URL,
	cacheKey string,
) []byte {
	currDate := time.Now()

	expireDate := self.ExpireDate
	if expireDate.IsZero() {
		expireDate = currDate.Add(time.Hour * 24 * 365)
	}

	requestTime := currDate
	responseTime := currDate.Add(time.Second)

	contentData := []byte(self.Content)

	sslInfo := getCertInfoFromURL(pageUrl)
	flags := getRespFlags(pageUrl, sslInfo)
	pickle := NewPickle()

	bigEndian := false 
	pickle.WriteInt32(int32(flags), bigEndian)
	pickle.WriteInt64(timeToInt64(requestTime), bigEndian)
	pickle.WriteInt64(timeToInt64(responseTime), bigEndian)

	generateHeaderData(
		pickle,
		len(contentData), 
		self.ContentType,
		currDate, 
		expireDate,
	)
	
	if sslInfo != nil {
		persistCert(pickle, sslInfo)
	}

	host := pageUrl.Hostname()
	pickle.WriteString(host)

	port := pageUrl.Port()
	portInt, _ := strconv.Atoi(port)

	pickle.WriteUInt16(uint16(portInt), bigEndian)
	pickle.WriteUInt16(0x1, bigEndian)

	if sslInfo != nil && sslInfo.keyExchangeGroup != 0 {
		pickle.WriteInt32(int32(sslInfo.keyExchangeGroup), bigEndian)
	}
	if sslInfo != nil && sslInfo.peerSignatureAlgorithm != 0 {
		pickle.WriteInt32(int32(sslInfo.peerSignatureAlgorithm), bigEndian)
	}
	return pickle.Bytes()
}

func getRespFlags(pageUrl *url.URL, sslInfo *SSlInfo) int {
	flags := RESPONSE_INFO_VERSION

	isSslInfoValid := sslInfo != nil
	isSslInfoPkpBypassed := isSslInfoValid

	hasConnectionInfo := true 

	if isSslInfoValid {
		flags |= RESPONSE_INFO_HAS_CERT
		flags |= RESPONSE_INFO_HAS_CERT_STATUS

		if sslInfo.keyExchangeGroup != 0 {
			flags |= RESPONSE_INFO_HAS_KEY_EXCHANGE_GROUP
		}
		if sslInfo.connectionStatus != 0 {
			flags |= RESPONSE_INFO_HAS_SSL_CONNECTION_STATUS
		}
		if sslInfo.peerSignatureAlgorithm != 0 {
			flags |= RESPONSE_INFO_HAS_PEER_SIGNATURE_ALGORITHM
		}
	}
	if hasConnectionInfo {
		flags |= RESPONSE_INFO_HAS_CONNECTION_INFO
	}
	if isSslInfoPkpBypassed {
		flags |= RESPONSE_INFO_PKP_BYPASSED
	}
	return flags
}