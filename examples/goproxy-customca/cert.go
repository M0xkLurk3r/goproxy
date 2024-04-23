package main

import (
	"crypto/tls"
	"crypto/x509"

	"github.com/elazarl/goproxy"
)

var caCert = []byte(`-----BEGIN CERTIFICATE-----
MIIDkzCCAnugAwIBAgIJAKe/ZGdfcHdPMA0GCSqGSIb3DQEBCwUAMGAxCzAJBgNV
BAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBX
aWRnaXRzIFB0eSBMdGQxGTAXBgNVBAMMEGRlbW8gZm9yIGdvcHJveHkwHhcNMTYw
OTI3MTQzNzQ3WhcNMTkwOTI3MTQzNzQ3WjBgMQswCQYDVQQGEwJBVTETMBEGA1UE
CAwKU29tZS1TdGF0ZTEhMB8GA1UECgwYSW50ZXJuZXQgV2lkZ2l0cyBQdHkgTHRk
MRkwFwYDVQQDDBBkZW1vIGZvciBnb3Byb3h5MIIBIjANBgkqhkiG9w0BAQEFAAOC
AQ8AMIIBCgKCAQEA2+W48YZoch72zj0a+ZlyFVY2q2MWmqsEY9f/u53fAeTxvPE6
1/DnqsydnA3FnGvxw9Dz0oZO6xG+PZvp+lhN07NZbuXK1nie8IpxCa342axpu4C0
69lZwxikpGyJO4IL5ywp/qfb5a2DxPTAyQOQ8ROAaydoEmktRp25yicnQ2yeZW//
1SIQxt7gRxQIGmuOQ/Gqr/XN/z2cZdbGJVRUvQXk7N6NhQiCX1zlmp1hzUW9jwC+
JEKKF1XVpQbc94Bo5supxhkKJ70CREPy8TH9mAUcQUZQRohnPvvt/lKneYAGhjHK
vhpajwlbMMSocVXFvY7o/IqIE/+ZUeQTs1SUwQIDAQABo1AwTjAdBgNVHQ4EFgQU
GnlWcIbfsWJW7GId+6xZIK8YlFEwHwYDVR0jBBgwFoAUGnlWcIbfsWJW7GId+6xZ
IK8YlFEwDAYDVR0TBAUwAwEB/zANBgkqhkiG9w0BAQsFAAOCAQEAoFUjSD15rKlY
xudzyVlr6n0fRNhITkiZMX3JlFOvtHNYif8RfK4TH/oHNBTmle69AgixjMgy8GGd
H90prytGQ5zCs1tKcCFsN5gRSgdAkc2PpRFOK6u8HwOITV5lV7sjucsddXJcOJbQ
4fyVe47V9TTxI+A7lRnUP2HYTR1Bd0R/IgRAH57d1ZHs7omHIuQ+Ea8ph2ppXMnP
DXVOlZ9zfczSnPnQoomqULOU9Fq2ycyi8Y/ROtAHP6O7wCFbYHXhxojdaHSdhkcd
troTflFMD2/4O6MtBKbHxSmEG6H0FBYz5xUZhZq7WUH24V3xYsfge29/lOCd5/Xf
A+j0RJc/lQ==
-----END CERTIFICATE-----`)

var caKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA2+W48YZoch72zj0a+ZlyFVY2q2MWmqsEY9f/u53fAeTxvPE6
1/DnqsydnA3FnGvxw9Dz0oZO6xG+PZvp+lhN07NZbuXK1nie8IpxCa342axpu4C0
69lZwxikpGyJO4IL5ywp/qfb5a2DxPTAyQOQ8ROAaydoEmktRp25yicnQ2yeZW//
1SIQxt7gRxQIGmuOQ/Gqr/XN/z2cZdbGJVRUvQXk7N6NhQiCX1zlmp1hzUW9jwC+
JEKKF1XVpQbc94Bo5supxhkKJ70CREPy8TH9mAUcQUZQRohnPvvt/lKneYAGhjHK
vhpajwlbMMSocVXFvY7o/IqIE/+ZUeQTs1SUwQIDAQABAoIBAHK94ww8W0G5QIWL
Qwkc9XeGvg4eLUxVknva2Ll4fkZJxY4WveKx9OCd1lv4n7WoacYIwUGIDaQBZShW
s/eKnkmqGy+PvpC87gqL4sHvQpuqqJ1LYpxylLEFqduWOuGPUVC2Lc+QnWCycsCS
CgqZzsbMq0S+kkKRGSvw32JJneZCzqLgLNssQNVk+Gm6SI3s4jJsGPesjhnvoPaa
xZK14uFpltaA05GSTDaQeZJFEdnnb3f/eNPc2xMEfi0S2ZlJ6Q92WJEOepAetDlR
cRFi004bNyTb4Bphg8s4+9Cti5is199aFkGCRDWxeqEnc6aMY3Ezu9Qg3uttLVUd
uy830GUCgYEA7qS0X+9UH1R02L3aoANyADVbFt2ZpUwQGauw9WM92pH52xeHAw1S
ohus6FI3OC8xQq2CN525tGLUbFDZnNZ3YQHqFsfgevfnTs1//gbKXomitev0oFKh
VT+WYS4lkgYtPlXzhdGuk32q99T/wIocAguvCUY3PiA7yBz93ReyausCgYEA6+P8
bugMqT8qjoiz1q/YCfxsw9bAGWjlVqme2xmp256AKtxvCf1BPsToAaJU3nFi3vkw
ICLxUWAYoMBODJ3YnbOsIZOavdXZwYHv54JqwqFealC3DG0Du6fZYZdiY8pK+E6m
3fiYzP1WoVK5tU4bH8ibuIQvpcI8j7Gy0cV6/AMCgYBHl7fZNAZro72uLD7DVGVF
9LvP/0kR0uDdoqli5JPw12w6szM40i1hHqZfyBJy042WsFDpeHL2z9Nkb1jpeVm1
C4r7rJkGqwqElJf6UHUzqVzb8N6hnkhyN7JYkyyIQzwdgFGfaslRzBiXYxoa3BQM
9Q5c3OjDxY3JuhDa3DoVYwKBgDNqrWJLSD832oHZAEIicBe1IswJKjQfriWWsV6W
mHSbdtpg0/88aZVR/DQm+xLFakSp0jifBTS0momngRu06Dtvp2xmLQuF6oIIXY97
2ON1owvPbibSOEcWDgb8pWCU/oRjOHIXts6vxctCKeKAFN93raGphm0+Ck9T72NU
BTubAoGBAMEhI/Wy9wAETuXwN84AhmPdQsyCyp37YKt2ZKaqu37x9v2iL8JTbPEz
pdBzkA2Gc0Wdb6ekIzRrTsJQl+c/0m9byFHsRsxXW2HnezfOFX1H4qAmF6KWP0ub
M8aIn6Rab4sNPSrvKGrU6rFpv/6M33eegzldVnV9ku6uPJI1fFTC
-----END RSA PRIVATE KEY-----`)

func setCA(caCert, caKey []byte) error {
	goproxyCa, err := tls.X509KeyPair(caCert, caKey)
	if err != nil {
		return err
	}
	if goproxyCa.Leaf, err = x509.ParseCertificate(goproxyCa.Certificate[0]); err != nil {
		return err
	}
	goproxy.GoproxyCa = goproxyCa
	goproxy.OkConnect = &goproxy.ConnectAction{Action: goproxy.ConnectAccept, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	goproxy.MitmConnect = &goproxy.ConnectAction{Action: goproxy.ConnectMitm, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	goproxy.HTTPMitmConnect = &goproxy.ConnectAction{Action: goproxy.ConnectHTTPMitm, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	goproxy.RejectConnect = &goproxy.ConnectAction{Action: goproxy.ConnectReject, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	goproxy.DefaultHijackTLSConfig = goproxy.TLSConfigFromCA(&goproxyCa)
	return nil
}
