package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table default_tools...")
		_, err := db.Exec(`INSERT INTO default_tools(type, category, options, format, can_change, selected)
				VALUES('ping', 'reconnaissance', '', 'ping {url}', 'f', 't'),
				('whois', 'reconnaissance', '', 'whois {url}', 'f', 't'),		
				('dig', 'dns_enumeration', '+short', 'dig {url} +short', 'f', 't'),
				('dig', 'dns_enumeration', '+short A', 'dig {url} +short A', 't', 't'),
				('dig', 'dns_enumeration', '+short AAAA', 'dig {url} +short AAAA', 't', 't'),
				('dig', 'dns_enumeration', 'MX', 'dig {url} MX', 't', 't'),
				('dig', 'dns_enumeration', 'TXT', 'dig {url} TXT', 't', 't'),
				('dig', 'dns_enumeration', 'ANY', 'dig {url} ANY', 't', 't'),
				('dig', 'dns_enumeration', 'SOA', 'dig {url} SOA', 't', 't'),
				('nslookup', 'dns_enumeration', ' | awk "/^Address: / { print $2 }"', 'nslookup {url} | awk "/^Address: / { print $2 }"', 'f', 't'),
				('nslookup', 'dns_enumeration', '-type=A', 'nslookup -type=A {url}', 't', 't'),
				('nslookup', 'dns_enumeration', '-type=cname', 'nslookup -type=cname {url}', 't', 't'),
				('nslookup', 'dns_enumeration', '-type=SOA', 'nslookup -type=SOA {url}', 't', 't'),
				('nslookup', 'dns_enumeration', '-type=NX', 'nslookup -type=NX {url}', 't', 't'),
				('nslookup', 'dns_enumeration', '-type=MS', 'nslookup -type=MS {url}', 't', 't'),
				('nslookup', 'dns_enumeration', '-type=TXT', 'nslookup -type=TXT {url}', 't', 't'),
				('shodan', 'reconnaissance', '+short', 'shodan host {ip_address}', 'f', 't'),
				('ffuf', 'web_enumeration', '', 'ffuf -u {url} -w /usr/share/wordlists/dirb/common.txt -p 1 fc 301', 'f', 't'),
				('dirb', 'web_enumeration', '-w /usr/share/wordlists/dirb/common.txt', 'dirb {url} -w /usr/share/wordlists/dirb/common.txt', 'f', 't'),
				('wfuzz', 'web_enumeration', '', 'wfuzz -c -w /usr/share/wordlists/dirb/common.txt {url}/FUZZ', 'f', 't'),
				('Wappalyzer', 'web_enumeration', '', 'python wappy.py -u {url}', 'f', 't'),
				('nmap', 'network_scanning', '', 'nmap -sV {ip_address}', 't', 't')
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
// 		VALUES('dig', 'dns_enumeration', '+short', 'dig {url} +short', 't')
// 		)`)
// 		return err
// 	}, func(db migrations.DB) error {
// 		fmt.Println("delete from table default_tools...")
// 		_, err := db.Exec(`DELETE FROM TABLE default_tools`)
// 		return err
// 	})
// }
