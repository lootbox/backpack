package main

import "fmt"

func main() {
	// functions checking
	listNodes()
	addSSHCreds()
	fmt.Println(getIP(6), getIP(4))
	fmt.Println(filterAMI("LatestRedHatAMIFilter"))
	fmt.Println(getLatestAMI("LatestRedHatAMIFilter"))
	fmt.Println(listCNAMES())
	fmt.Println(getDomainNameByIP("10.80.100.80"))
	fmt.Println(getIPByDomainName("domain.company.tech."))
	fmt.Println(getEC2InstanceByName("domain.company.tech."))
}
