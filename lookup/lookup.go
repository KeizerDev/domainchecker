package lookup

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/domainr/whois"
	"github.com/gosuri/uilive"
	"github.com/gosuri/uitable"
	"golang.org/x/text/language"
)

const (
	defaultLanguage = "en"
	defaultRegion   = "US"
)

type domaincheck struct {
	Domain, Available string
	Id                int
}

var domainExtensions = []string{"com", "net", "org", "nl", "de", "io"}

var domainscheck = []domaincheck{}

// Provider interface provides a way to build the URI
// for each provider.

// Providers tracks loaded providers.
// var Providers map[string]Provider

func init() {
	// 	Providers = make(map[string]Provider)
}

// AddProvider should be called within your provider's init() func.
// This will register the provider so it can be used.
// func AddProvider(name string, provider Provider) {
// 	Providers[name] = provider
// }

// Search builds a search URL and opens it in your browser.
func DoLookUp(p string, qwhois string, verbose bool) error {
	writer := uilive.New()
	writer.Start()

	table := uitable.New()
	table.MaxColWidth = 50
	if strings.HasSuffix(qwhois, ".*") {
		qwhois = strings.TrimSuffix(qwhois, ".*")

		for i, domain := range domainExtensions {
			domainscheck = append(domainscheck, domaincheck{fmt.Sprintf("%s.%s", qwhois, domain), "progress", i})
			table.AddRow("[☓]", fmt.Sprintf("%s.%s", qwhois, domain))
		}

		// lookup.doWhois(verbose, qwhois) // Add array support
	}

	fmt.Fprintln(writer, table)
	writer.Stop() // flush and stop rendering

	fmt.Printf(doWhois(qwhois, verbose))
	if verbose {
		fmt.Printf("%s\n", qwhois)
	}

	return nil
}

func doWhois(qwhois string, verbose bool) string {
	request, err := whois.NewRequest(qwhois)
	FatalIf(err)

	response, err := whois.DefaultClient.Fetch(request)
	FatalIf(err)

	if verbose {
		fmt.Printf("%s\n", response)
	}

	r := regexp.MustCompile(`(No match)|(^NOT FOUND)|(^Not fo|AVAILABLE)|(^No Data Fou|has not been regi|No entri)|(Status: free)|(.nl is free)`)
	if v := response.String(); r.MatchString(v) {
		return "beschikbaar\n"
	} else {
		return "NIET beschikbaar\n"
	}
}

// Region returns the users region code.
// Eg. "US", "GB", etc
func Region() string {
	l := locale()

	tag, err := language.Parse(l)
	if err != nil {
		return defaultRegion
	}

	region, _ := tag.Region()

	return region.String()
}

// Language returns the users language code.
// Eg. "en", "es", etc
func Language() string {
	l := locale()

	tag, err := language.Parse(l)
	if err != nil {
		return defaultLanguage
	}

	base, _ := tag.Base()

	return base.String()
}

func locale() string {
	lang := os.Getenv("LANG")
	if lang == "" {
		return ""
	}

	locale := strings.Split(lang, ".")[0]

	return locale
}

func FatalIf(err error) {
	if err == nil {
		return
	}
	fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	os.Exit(-1)
}
