package utils

import (
	"fmt"
	"log"

	"gopkg.in/ldap.v3"
)

// ConnLdap conn ldap
func ConnLdap() (*ldap.Conn, error) {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", APIConfig.LDAPConfig.Addr, APIConfig.LDAPConfig.Port))

	if err != nil {
		return nil, err
	}

	return l, nil
}

// UserAuthentication ldap user auth
func UserAuthentication(username, password string) (bool, error) {
	l, err := ConnLdap()
	defer func() {
		l.Close()
	}()
	if err != nil {
		return false, err
	}

	// TODO: ldap conn support tls
	/*err = l.StartTLS(&tls.Config{InsecureSkipVerify: true})
	if err != nil {
		return false, err
	}*/

	err = l.Bind(APIConfig.LDAPConfig.BindDN, APIConfig.LDAPConfig.BindPass)
	if err != nil {
		return false, err
	}

	searchRequest := ldap.NewSearchRequest(
		APIConfig.LDAPConfig.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf("(&(objectClass=organizationalPerson)(uid=%s))", username),
		[]string{"dn"},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		return false, err
	}

	if len(sr.Entries) != 1 {
		return false, fmt.Errorf("username %s not exists", username)
	}

	userDN := sr.Entries[0].DN

	err = l.Bind(userDN, password)
	if err != nil {
		return false, fmt.Errorf("username %s password is worng", username)
	}
	return true, nil
}

// LdapSearch ldap search
func LdapSearch() {
	l, err := ConnLdap()
	if err != nil {
		log.Fatal("conn error", err)
	}

	defer l.Close()

	searchResult := ldap.NewSearchRequest(
		APIConfig.LDAPConfig.BindDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		"(&(objectClass=organizationalPerson))",
		[]string{"dn", "cn"}, nil)

	sr, err := l.Search(searchResult)
	if err != nil {
		log.Fatal("search error", err)
	}

	for _, entry := range sr.Entries {
		fmt.Printf("%s: %v\n", entry.DN, entry.GetAttributeValue("cn"))
	}
}
