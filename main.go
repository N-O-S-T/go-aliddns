//go:build darwin || linux || windows

package main

import (
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"go-aliddns/aliddns"
	"os"
)

var isTest = flag.Bool("t", false, "测试用例，开启测试")
var certIs = flag.Bool("ct", false, "启用certbot dns更新")
var Key = flag.String("k", os.Getenv("CERTBOT_VALIDATION"), "自动获取系统certbot参数，也可以进行传参")

func main() {
	flag.Parse()
	if flag.Parsed() {
		var qs = new(aliddns.QueryStruct)
		var account Account

		conf.MustLoad(*configFile, &account)
		// 传参
		qs.IsTest = *isTest
		qs.AccessKeyId = account.AccessKeyId
		qs.AccessSecret = account.AccessSecret
		qs.MainDomain = account.MainDomain
		for _, subDomainConf := range account.SubDomains {
			qs.SubDomain = subDomainConf.SubDomain
			// 此处暂未完成
			if *certIs && subDomainConf.Cert {
				dnsHeader := "_acme-challenge."
				qs.ValueType = "TXT"
				qs.SubDomain = dnsHeader + qs.SubDomain
				qs.Value = *Key
				qs.DnsCheck()
				continue
			}
			if subDomainConf.Ipv4 {
				qs.ValueType = "A"
				qs.GetOutBoundIP()
				qs.DnsCheck()
			}
			if subDomainConf.Ipv6 {
				qs.ValueType = "AAAA"
				qs.GetOutBoundIP()
				qs.DnsCheck()
			}
		}
	}
}
