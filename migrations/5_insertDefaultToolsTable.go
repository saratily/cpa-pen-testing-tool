package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table default_tools...")
		_, err := db.Exec(`INSERT INTO default_tools(type, category, options, format, active, can_change, selected)
				VALUES('reachable', 'reconnaissance', '{{.URL}}', 'GET {{.URL}}', 1, 0, 0),
				('whois', 'reconnaissance', '{{.domain}}', 'whois {{.domain}}', 1, 0, 0),
				('ping', 'reconnaissance', '{{.URL}}', 'ping {{.URL}}', 1, 0, 0),
				('digIPv4', 'dns_enumeration', '{{.domain}}', 'dig {{.domain}} +short A', 1, 0, 0),
				('digIPv6', 'dns_enumeration', '{{.domain}}', 'dig {{.domain}} +short AAAA', 1, 0, 0),
				('digCNAME', 'dns_enumeration', '{{.domain}}', 'dig {{.domain}} cname', 1, 0, 0),
				('digMX', 'dns_enumeration', '{{.domain}}', 'dig {{.domain}} MX', 1, 0, 0),
				('digNS', 'dns_enumeration', '{{.domain}}', 'dig {{.domain}} NS', 1, 0, 0),
				('digTXT', 'dns_enumeration', '{{.domain}}', 'dig {{.domain}} TXT', 1, 0, 0),
				('digANY', 'dns_enumeration', '{{.domain}}', 'dig {{.domain}} ANY', 1, 0, 0),
				('digSOA', 'dns_enumeration', '{{.domain}}', 'dig {{.domain}} SOA', 1, 0, 0),
				('LookupIP', 'dns_enumeration', '{{.domain}}', 'nslookup {{.URL}} | awk "/^Address: / { print $2 }"', 1, 0, 0),
				('LookupCNAME', 'dns_enumeration', '{{.domain}}', 'nslookup -type=cname {{.domain}}', 1, 0, 0),
				('LookupMX', 'dns_enumeration', '{{.domain}}', 'nslookup -type=MX {{.domain}}', 1, 0, 0),
				('LookupNS', 'dns_enumeration', '{{.domain}}', 'nslookup -type=NS {{.domain}}', 1, 0, 0),
				('LookupSRV', 'dns_enumeration', '{{.domain}}', 'nslookup -type=SRV {{.domain}}', 1, 0, 0),
				('LookupTXT', 'dns_enumeration', '{{.domain}}', 'nslookup -type=TXT {{.domain}}', 1, 0, 0),
				('shodan', 'reconnaissance', '{{.ip_address}}', 'shodan host {{.ip_address}}', 1, 0, 0),
				('ffuf', 'web_enumeration', '', 'ffuf -u {{.URL}} -w /usr/share/wordlists/dirb/common.txt -p 1 fc 301', 1, 0, 0),
				('dirb', 'web_enumeration', '-w /usr/share/wordlists/dirb/common.txt', 'dirb {{.URL}} -w /usr/share/wordlists/dirb/common.txt', 1, 0, 0),
				('wfuzz', 'web_enumeration', '', 'wfuzz -c -w /usr/share/wordlists/dirb/common.txt {{.URL}}/FUZZ', 1, 0, 0),
				('Wappalyzer', 'web_enumeration', '', 'python wappy.py -u {{.URL}}', 1, 0, 0),
				('nmap', 'network_scanning', '', 'nmap -sV {{.URL}}', 1, 1, 0)
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
