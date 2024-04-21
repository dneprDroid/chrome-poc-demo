package injector 

import (
	"net/url"
	"crypto/x509"

	tls "boringssl.googlesource.com/boringssl/ssl/test/runner"
)

const (
	CERT_STATUS_OK = 0
)

const (
	kCertDisconnected = iota
	kCertConnecting
	kCertConnected 
)

func getCertInfoFromHost(host string) *SSLInfo {
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}
	conn, err := tls.Dial("tcp", host + ":443", conf)
	if err != nil {
		return nil
	}
	defer conn.Close()
	connState := conn.ConnectionState()
	certs := connState.PeerCertificates

	return &SSLInfo{
		keyExchangeGroup: int(connState.CurveID),
		peerSignatureAlgorithm: int(connState.PeerSignatureAlgorithm), 
		certStatus: CERT_STATUS_OK,
		connectionStatus: kCertConnected,
		certs: certs,
	}
}

func getCertInfoFromURL(pageUrl *url.URL) *SSLInfo {
	isHttps := pageUrl.Scheme == "https"
	if !isHttps {
		return nil
	}
	host := pageUrl.Hostname()
	return getCertInfoFromHost(host)
}

type SSLInfo struct {
	keyExchangeGroup int
	peerSignatureAlgorithm int 

	certStatus int
	connectionStatus int

	certs []*x509.Certificate
}

func persistCert(pickle *Pickle, sslInfo *SSLInfo) {
	certs := sslInfo.certs 
	if len(certs) == 0 {
		return
	}
	bigEndian := false

	pickle.WriteInt32(int32(len(certs)), bigEndian)
	pickle.WriteBytesString(certs[0].Raw)
		
	for i := 1; i < len(certs); i += 1 {
		pickle.WriteBytesString(certs[i].Raw)
	}
	pickle.WriteUint32(uint32(sslInfo.certStatus), bigEndian)

 	if sslInfo.connectionStatus != 0 {
		pickle.WriteInt32(int32(sslInfo.connectionStatus), bigEndian)
	}
}