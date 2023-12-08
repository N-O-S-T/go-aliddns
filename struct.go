package main

type Account struct {
	AccessKeyId  string          `json:"AccessKeyId"`
	AccessSecret string          `json:"AccessSecret"`
	MainDomain   string          `json:"MainDomain"`
	SubDomains   []SubDomainConf `json:"SubDomains"`
}
type SubDomainConf struct {
	SubDomain string `json:"SubDomain"`
	Ipv4      bool   `json:"ipv4"`
	Ipv6      bool   `json:"ipv6"`
	Cert      bool   `json:"cert"`
}
