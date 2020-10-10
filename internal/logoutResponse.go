package internal

import (
	"encoding/xml"
	"github.com/LoginRadius/go-saml/util"
	"time"
)

type LogoutResponse struct {
	XMLName      xml.Name
	XMLNSP       string     `xml:"xmlns:samlp,attr"`
	XMLNSL       string     `xml:"xmlns:saml,attr"`
	ID           string     `xml:"ID,attr"`
	Version      string     `xml:"Version,attr"`
	IssueInstant string     `xml:"IssueInstant,attr"`
	Destination  string     `xml:"Destination,attr"`
	InResponseTo string     `xml:"InResponseTo,attr,omitempty"`
	Issuer       Issuer     `xml:"Issuer"`
	Signature    *Signature `xml:"Signature,omitempty"`
	Status       Status     `xml:"Status"`
}

func NewLogoutResponse() *LogoutResponse {
	responseID := util.ID()
	issueInstant := time.Now().UTC().Format(time.RFC3339)
	response := &LogoutResponse{
		XMLName: xml.Name{
			Local: "samlp:LogoutResponse",
		},
		XMLNSP:       "urn:oasis:names:tc:SAML:2.0:protocol",
		XMLNSL:       "urn:oasis:names:tc:SAML:2.0:assertion",
		ID:           responseID,
		Version:      "2.0",
		IssueInstant: issueInstant,
		Destination:  "",
		InResponseTo: "",
		Issuer: Issuer{
			XMLName: xml.Name{
				Local: "saml:Issuer",
			},
		},
		Signature: &Signature{
			XMLName: xml.Name{
				Local: "ds:Signature",
			},
			SAMLSIG: "http://www.w3.org/2000/09/xmldsig#",
			SignedInfo: SignedInfo{
				XMLName: xml.Name{
					Local: "ds:SignedInfo",
				},
				CanonicalizationMethod: CanonicalizationMethod{
					XMLName: xml.Name{
						Local: "ds:CanonicalizationMethod",
					},
					Algorithm: "http://www.w3.org/2001/10/xml-exc-c14n#",
				},
				SignatureMethod: SignatureMethod{
					XMLName: xml.Name{
						Local: "ds:SignatureMethod",
					},
					Algorithm: "", // populated by SetSignatureAlgorithm
				},
				SamlsigReference: SamlsigReference{
					XMLName: xml.Name{
						Local: "ds:Reference",
					},
					URI: "#" + responseID,
					Transforms: Transforms{
						XMLName: xml.Name{
							Local: "ds:Transforms",
						},
						Transform: []Transform{Transform{
							XMLName: xml.Name{
								Local: "ds:Transform",
							},
							Algorithm: "http://www.w3.org/2000/09/xmldsig#enveloped-signature",
						}, Transform{
							XMLName: xml.Name{
								Local: "ds:Transform",
							},
							Algorithm: "http://www.w3.org/2001/10/xml-exc-c14n#",
						}},
					},
					DigestMethod: DigestMethod{
						XMLName: xml.Name{
							Local: "ds:DigestMethod",
						},
						Algorithm: "", // populated by SetDigestAlgorithm
					},
					DigestValue: DigestValue{
						XMLName: xml.Name{
							Local: "ds:DigestValue",
						},
					},
				},
			},
			SignatureValue: SignatureValue{
				XMLName: xml.Name{
					Local: "ds:SignatureValue",
				},
			},
			KeyInfo: KeyInfo{
				XMLName: xml.Name{
					Local: "ds:KeyInfo",
				},
				X509Data: X509Data{
					XMLName: xml.Name{
						Local: "ds:X509Data",
					},
					X509Certificate: X509Certificate{
						XMLName: xml.Name{
							Local: "ds:X509Certificate",
						},
					},
				},
			},
		},
		Status: Status{
			XMLName: xml.Name{
				Local: "samlp:Status",
			},
			StatusCode: StatusCode{
				XMLName: xml.Name{
					Local: "samlp:StatusCode",
				},
				Value: "urn:oasis:names:tc:SAML:2.0:status:Success",
			},
		},
	}
	return response
}

func (r *LogoutResponse) String() (string, error) {
	x, err := xml.MarshalIndent(r, "", "    ")
	if err != nil {
		return "", err
	}
	return string(x), nil
}

func (r *LogoutResponse) SetInResponseTo(inResponseTo string) {
	r.InResponseTo = inResponseTo
}

func (r *LogoutResponse) SetSignatureAlgorithm(alg string) {
	r.Signature.SignedInfo.SignatureMethod.Algorithm = alg
}

func (r *LogoutResponse) SetDigestAlgorithm(alg string) {
	r.Signature.SignedInfo.SamlsigReference.DigestMethod.Algorithm = alg
}

func (r *LogoutResponse) SignedXML(idpPrivateKey interface{}) (string, error) {
	xmlStr, err := r.String()
	if err != nil {
		return "", err
	}
	signedXML, err := util.Sign(xmlStr, idpPrivateKey)
	if err != nil {
		return "", err
	}
	return signedXML, nil
}
