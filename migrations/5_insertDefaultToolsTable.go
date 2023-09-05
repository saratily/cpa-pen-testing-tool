package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table default_tools...")
		_, err := db.Exec(`INSERT INTO default_tools(type, category, options, format, help, active, can_change, selected)
				VALUES('whois', 'reconnaissance', '{{.domain}}', 'whois {{.domain}}', 'whois is a query and response protocol that is used for querying databases that store an Internet resources registered users or assignees.', 1, 0, 0),
				('ping', 'reconnaissance', '{{.URL}}', 'ping {{.URL}}', 'ping is a computer network administration software utility used to test the reachability of a host on an Internet Protocol (IP) network', 1, 0, 0),
				('digIPv4', 'dns_enumeration', '{{.domain}}', 'dig {{.domain}} +short A', 'Find the IPv4 address (A) of a Hostname', 1, 0, 0),
				('digIPv6', 'dns_enumeration', '{{.domain}}', 'dig {{.domain}} +short AAAA', 'Find the IPv6 address (A) of a Hostname', 1, 0, 0),
				('digCNAME', 'dns_enumeration', '{{.domain}}', 'dig {{.domain}} cname', 'Find the CNAME record of a domain/host name', 1, 0, 0),
				('digMX', 'dns_enumeration', '{{.domain}}', 'dig {{.domain}} MX', 'Find the MX (mail exchanges) of a domain/host name', 1, 0, 0),
				('digNS', 'dns_enumeration', '{{.domain}}', 'dig {{.domain}} NS', 'Find the NS (name servers) of a domain name', 1, 0, 0),
				('digTXT', 'dns_enumeration', '{{.domain}}', 'dig {{.domain}} TXT', 'Find the TXT (text annotations) of a domain name', 1, 0, 0),
				('digANY', 'dns_enumeration', '{{.domain}}', 'dig {{.domain}} ANY', 'ANY query i.e. try all queries to get as much as information possible', 1, 0, 0),
				('digSOA', 'dns_enumeration', '{{.domain}}', 'dig {{.domain}} SOA', 'Show the SOA records for DNS zone name', 1, 0, 0),
				('LookupIP', 'dns_enumeration', '{{.domain}}', 'nslookup {{.URL}} | awk "/^Address: / { print $2 }"', 'Find the IP address (A) of a Hostname', 1, 0, 0),
				('LookupCNAME', 'dns_enumeration', '{{.domain}}', 'nslookup -type=cname {{.domain}}', 'Find the CNAME record of a domain/host name', 1, 0, 0),
				('LookupMX', 'dns_enumeration', '{{.domain}}', 'nslookup -type=MX {{.domain}}', 'Find the MX (mail exchanges) of a domain/host name', 1, 0, 0),
				('LookupNS', 'dns_enumeration', '{{.domain}}', 'nslookup -type=NS {{.domain}}', 'Find the NS (name servers) of a domain name', 1, 0, 0),
				('LookupSRV', 'dns_enumeration', '{{.domain}}', 'nslookup -type=SRV {{.domain}}', 'SRV attempts to resolve an SRV query of the service, protocol (tcp or udp), and domain name, sorted by priority and randomized by weight within a priority.', 1, 0, 0),
				('LookupTXT', 'dns_enumeration', '{{.domain}}', 'nslookup -type=TXT {{.domain}}', 'Find the TXT (text annotations) of a domain name', 1, 0, 0),
				('shodan', 'reconnaissance', '{{.domain}}', 'shodan stats {{.domain}}', 'Shodan is a search engine that lets users search for various types of servers (webcams, routers, servers, etc.) connected to the internet using a variety of filters', 1, 0, 0),
				('ffuf', 'web_enumeration', '{{.domain}}', 'ffuf -u http://{{.domain}}/FUZZ -w /usr/share/wordlists/dirb/common.txt -p 1', 'Ffuf is a lightning-fast subdomain fuzzer that’s highly customizable.', 1, 0, 0),
				('ffufNot301', 'web_enumeration', '{{.domain}}', 'ffuf -u http://FUFF.{{.domain}} -w /usr/share/wordlists/dirb/common.txt -p 1 fc 301', '', 1, 0, 0),
				('ffufSubdomain', 'web_enumeration', '{{.domain}}', 'ffuf -u http://api.{{.domain}}/FUZZ -w /usr/share/wordlists/dirb/common.txt -p 1', '', 1, 0, 0),
				('dirb', 'web_enumeration', '{{.domain}}', 'dirb http://{{.domain}} -w /usr/share/wordlists/dirb/common.txt', 'dirb is a Web Content Scanner. It looks for existing (and/or hidden) Web Objects. It basically works by launching a dictionary based attack against a web server and analyzing the responses.', 1, 0, 0),
				('wfuzz', 'web_enumeration', '{{.domain}}', 'wfuzz -c -w /usr/share/wordlists/dirb/common.txt http://{{.domain}}/FUZZ', 'wfuzz is one of the most powerful and versatile subdomain fuzzing tools out there. With Wfuzz, you can quickly and easily discover hidden subdomains on a target website, making it an essential tool for any penetration tester or ethical hacker.', 1, 0, 0),
				('wappalyzer', 'web_enumeration', '{{.domain}}', 'wappy -u {{.domain}}', 'wappalyzer is a browser extension used for identifying the technologies used by websites. It is an excellent tool for web reconnaissance and can be used by both security researchers and web developers.', 1, 0, 0),
				('nmap', 'network_scanning', '{{.domain}}', 'nmap -sT {{.domain}}', 'Nmap is used to discover hosts and services on a computer network by sending packets and analyzing the responses.', 1, 0, 0),
				('nikto', 'network_scanning', '{{.domain}}', 'nikto -h {{.domain}}', 'Nikto is a free software command-line vulnerability scanner that scans web servers for dangerous files/CGIs, outdated server software and other problems. It performs generic and server type specific checks. It also captures and prints any cookies received. ', 1, 0, 0)
		`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("delete from table default_tools...")
		_, err := db.Exec(`TRUNCATE TABLE default_tools`)
		return err
	})
}

// package main

// import (
// 	"fmt"

// 	"github.com/go-pg/migrations/v8"
// )

// func init() {
// 	migrations.MustRegisterTx(func(db migrations.DB) error {
// 		fmt.Println("creating table default_tools...")
// 		_, err := db.Exec(`INSERT INTO TABLE default_tools('type', 'category', 'options', 'format', 'selected')
// 		VALUES('dig', 'dns_enumeration', '+short', 'dig {{.URL}} +short', 1)
// 		)`)
// 		return err
// 	}, func(db migrations.DB) error {
// 		fmt.Println("delete from table default_tools...")
// 		_, err := db.Exec(`DELETE FROM TABLE default_tools`)
// 		return err
// 	})
// }
