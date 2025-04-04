package main

import "fmt"

func main() {
	websites := map[string]string{
		"SitePoint":  "https://www.sitepoint.com",
		"GitHub":     "https://github.com",
		"GoLang":     "https://golang.org",
		"Google":     "https://www.google.com",
		"Docker":     "https://www.docker.com",
		"Kubernetes": "https://kubernetes.io",
	}

	//fmt.Println(websites)
	fmt.Println(websites["GitHub"])

	delete(websites, "GitHub")
	//fmt.Println(websites)

	userNames := make([]string, 2)
	userNames[0] = "John Doe"
	userNames = append(userNames, "Jane Doe")
	fmt.Println(userNames)

	courses := make(intMap, 2)
	courses["Math"] = 90
	courses["Science"] = 85
	fmt.Println(courses)

	for i, val := range websites {
		fmt.Printf("%s: %s\n", i, val)
	}
}

type intMap map[string]int
