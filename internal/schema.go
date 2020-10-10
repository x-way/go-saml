package internal

import (
	"encoding/xml"
	"time"
)

type AuthnRequest struct {
	XMLName                        xml.Name
	SAMLP                          string                `xml:"xmlns:samlp,attr"`
	SAML                           string                `xml:"xmlns:saml,attr"`
	ID                             string                `xml:"ID,attr"`
	Version                        string                `xml:"Version,attr"`
	ProtocolBinding                string                `xml:"ProtocolBinding,attr"`
	AssertionConsumerServiceURL    string                `xml:"AssertionConsumerServiceURL,attr"`
	Destination                    string                `xml:"Destination,attr"`
	IssueInstant                   time.Time             `xml:"IssueInstant,attr"`
	AssertionConsumerServiceIndex  int                   `xml:"AssertionConsumerServiceIndex,attr"`
	AttributeConsumingServiceIndex int                   `xml:"AttributeConsumingServiceIndex,attr"`
	Issuer                         Issuer                `xml:"Issuer"`
	NameIDPolicy                   NameIDPolicy          `xml:"NameIDPolicy"`
	RequestedAuthnContext          RequestedAuthnContext `xml:"RequestedAuthnContext"`
	Signature                      *Signature            `xml:"Signature,omitempty"`
	ForceAuthn                     string                `xml:"ForceAuthn"`
	IsPassive                      string                `xml:"IsPassive"`
	ProviderName                   string                `xml:"ProviderName"`
}

type Issuer struct {
	XMLName xml.Name
	SAML    string `xml:"xmlns:saml,attr,omitempty"`
	URL     string `xml:",innerxml"`
}

type NameIDPolicy struct {
	XMLName     xml.Name
	AllowCreate bool   `xml:"AllowCreate,attr"`
	Format      string `xml:"Format,attr"`
}

type RequestedAuthnContext struct {
	XMLName              xml.Name
	SAMLP                string               `xml:"xmlns:samlp,attr"`
	Comparison           string               `xml:"Comparison,attr"`
	AuthnContextClassRef AuthnContextClassRef `xml:"AuthnContextClassRef"`
}

type AuthnContextClassRef struct {
	XMLName   xml.Name
	SAML      string `xml:"xmlns:saml,attr,omitempty"`
	Transport string `xml:",innerxml"`
}

type Signature struct {
	XMLName        xml.Name
	SAMLSIG        string `xml:"xmlns:ds,attr"`
	SignedInfo     SignedInfo
	SignatureValue SignatureValue
	KeyInfo        KeyInfo
}

type SignedInfo struct {
	XMLName                xml.Name
	CanonicalizationMethod CanonicalizationMethod
	SignatureMethod        SignatureMethod
	SamlsigReference       SamlsigReference
}

type SignatureValue struct {
	XMLName xml.Name
	Value   string `xml:",innerxml"`
}

type KeyInfo struct {
	XMLName  xml.Name
	X509Data X509Data `xml:",innerxml"`
}

type CanonicalizationMethod struct {
	XMLName   xml.Name
	Algorithm string `xml:"Algorithm,attr"`
}

type SignatureMethod struct {
	XMLName   xml.Name
	Algorithm string `xml:"Algorithm,attr"`
}

type SamlsigReference struct {
	XMLName      xml.Name
	URI          string       `xml:"URI,attr"`
	Transforms   Transforms   `xml:",innerxml"`
	DigestMethod DigestMethod `xml:",innerxml"`
	DigestValue  DigestValue  `xml:",innerxml"`
}

type X509Data struct {
	XMLName         xml.Name
	X509Certificate X509Certificate `xml:",innerxml"`
}

type Transforms struct {
	XMLName   xml.Name
	Transform []Transform
}

type DigestMethod struct {
	XMLName   xml.Name
	Algorithm string `xml:"Algorithm,attr"`
}

type DigestValue struct {
	XMLName xml.Name
}

type X509Certificate struct {
	XMLName xml.Name
	Cert    string `xml:",innerxml"`
}

type Transform struct {
	XMLName   xml.Name
	Algorithm string `xml:"Algorithm,attr"`
}

type EntityDescriptor struct {
	XMLName  xml.Name
	DS       string `xml:"xmlns:ds,attr"`
	XMLNS    string `xml:"xmlns,attr"`
	MD       string `xml:"xmlns:md,attr"`
	EntityID string `xml:"entityID,attr"`

	Extensions      Extensions      `xml:"Extensions"`
	SPSSODescriptor SPSSODescriptor `xml:"SPSSODescriptor"`
}

type IDPEntityDescriptor struct {
	XMLName          xml.Name
	DS               string           `xml:"xmlns:ds,attr"`
	XMLNS            string           `xml:"xmlns,attr"`
	ASSERTION        string           `xml:"xmlns:assertion,attr"`
	EntityID         string           `xml:"entityID,attr"`
	IDPSSODescriptor IDPSSODescriptor `xml:"IDPSSODescriptor"`
}

type Extensions struct {
	XMLName xml.Name
	Alg     string `xml:"xmlns:alg,attr"`
	MDAttr  string `xml:"xmlns:mdattr,attr"`
	MDRPI   string `xml:"xmlns:mdrpi,attr"`

	EntityAttributes string `xml:"EntityAttributes"`
}

type IDPSSODescriptor struct {
	XMLName                    xml.Name
	WantAuthnRequestsSigned    string `xml:"WantAuthnRequestsSigned,attr"`
	ProtocolSupportEnumeration string `xml:"protocolSupportEnumeration,attr"`
	SigningKeyDescriptor       KeyDescriptor
	SingleSignOnService        []SingleSignOnService `xml:"SingleSignOnService"`
	SingleLogoutService        []SingleLogoutService `xml:"SingleLogoutService"`
}

type SPSSODescriptor struct {
	XMLName                    xml.Name
	ProtocolSupportEnumeration string `xml:"protocolSupportEnumeration,attr"`
	SigningKeyDescriptor       KeyDescriptor
	EncryptionKeyDescriptor    KeyDescriptor
	// SingleLogoutService        SingleLogoutService `xml:"SingleLogoutService"`
	AssertionConsumerServices []AssertionConsumerService `xml:"AssertionConsumerServices"`
}

type EntityAttributes struct {
	XMLName xml.Name
	SAML    string `xml:"xmlns:saml,attr"`

	EntityAttributes []Attribute `xml:"Attribute"` // should be array??
}

type SPSSODescriptors struct {
}

type KeyDescriptor struct {
	XMLName xml.Name
	Use     string  `xml:"use,attr"`
	KeyInfo KeyInfo `xml:"KeyInfo"`
}

type SingleSignOnService struct {
	XMLName  xml.Name
	Index    string `xml:"index,attr"`
	Binding  string `xml:"Binding,attr"`
	Location string `xml:"Location,attr"`
}

type SingleLogoutService struct {
	XMLName  xml.Name
	Index    string `xml:"index,attr"`
	Binding  string `xml:"Binding,attr"`
	Location string `xml:"Location,attr"`
}

type AssertionConsumerService struct {
	XMLName  xml.Name
	Binding  string `xml:"Binding,attr"`
	Location string `xml:"Location,attr"`
	Index    string `xml:"index,attr"`
}

type Response struct {
	XMLName        xml.Name
	SAMLP          string    `xml:"xmlns:samlp,attr"`
	SAML           string    `xml:"xmlns:saml,attr"`
	Destination    string    `xml:"Destination,attr"`
	ID             string    `xml:"ID,attr"`
	Version        string    `xml:"Version,attr"`
	IssueInstant   string    `xml:"IssueInstant,attr"`
	InResponseTo   string    `xml:"InResponseTo,attr,omitempty"`
	Issuer         Issuer    `xml:"Issuer"`
	Status         Status    `xml:"Status"`
	Assertion      Assertion `xml:"Assertion"`
}

type Assertion struct {
	XMLName            xml.Name
	ID                 string    `xml:"ID,attr"`
	Version            string    `xml:"Version,attr"`
	XSI                string    `xml:"xmlns:xsi,attr"`
	XS                 string    `xml:"xmlns:xs,attr"`
	SAML               string    `xml:"xmlns:saml,attr"`
	IssueInstant       string    `xml:"IssueInstant,attr"`
	Issuer             Issuer    `xml:"Issuer"`
	Signature          Signature `xml:"Signature"`
	Subject            Subject
	Conditions         Conditions
	AuthnStatement     AuthnStatement `xml:"AuthnStatement,omitempty"`
	AttributeStatement AttributeStatement
}

type Conditions struct {
	XMLName             xml.Name
	NotBefore           string              `xml:",attr"`
	NotOnOrAfter        string              `xml:",attr"`
	AudienceRestriction AudienceRestriction `xml:"AudienceRestriction,omitempty"`
}

type AudienceRestriction struct {
	XMLName   xml.Name
	Audiences []Audience `xml:"Audience"`
}

type Audience struct {
	XMLName xml.Name
	Value   string `xml:",innerxml"`
}

type Subject struct {
	XMLName             xml.Name
	NameID              NameID
	SubjectConfirmation SubjectConfirmation
}

type SubjectConfirmation struct {
	XMLName                 xml.Name
	Method                  string `xml:",attr"`
	SubjectConfirmationData SubjectConfirmationData
}

type Status struct {
	XMLName    xml.Name
	StatusCode StatusCode `xml:"StatusCode"`
}

type SubjectConfirmationData struct {
	XMLName      xml.Name
	InResponseTo string `xml:",attr,omitempty"`
	NotOnOrAfter string `xml:",attr"`
	Recipient    string `xml:",attr"`
}

type NameID struct {
	XMLName         xml.Name
	XMLNS           string `xml:"xmlns:saml,attr,omitempty"`
	Format          string `xml:",attr"`
	SPNameQualifier string `xml:",attr,omitempty"`
	Value           string `xml:",innerxml"`
}

type StatusCode struct {
	XMLName xml.Name
	Value   string `xml:",attr"`
}

type AttributeValue struct {
	XMLName xml.Name
	XS      string `xml:"xmlns:xs,attr"`
	XSI     string `xml:"xmlns:xsi,attr"`
	Type    string `xml:"xsi:type,attr"`
	Value   string `xml:",innerxml"`
}

type Attribute struct {
	XMLName         xml.Name
	Name            string         `xml:",attr"`
	FriendlyName    string         `xml:",attr,omitempty"`
	NameFormat      string         `xml:",attr"`
	AttributeValues AttributeValue `xml:"AttributeValue"`
}

type AttributeStatement struct {
	XMLName    xml.Name
	Attributes []Attribute `xml:"Attribute"`
}

type AuthnStatement struct {
	XMLName             xml.Name
	AuthnInstant        string       `xml:",attr"`
	SessionNotOnOrAfter string       `xml:",attr,omitempty"`
	SessionIndex        string       `xml:",attr,omitempty"`
	AuthnContext        AuthnContext `xml:"AuthnContext"`
}

type AuthnContext struct {
	XMLName              xml.Name
	AuthnContextClassRef AuthnContextClassRef `xml:"AuthnContextClassRef"`
}

type SessionIndex struct {
	XMLName xml.Name
	Value   string `xml:",innerxml"`
}
