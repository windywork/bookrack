package tools

import (
	"fmt"
	"github.com/fanux/fist/tools"
	"github.com/wonderivan/logger"
	"gopkg.in/ldap.v3"
	"strings"
)

func getLdapSearchResult(user string) (*ldap.SearchResult, error) {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", "fist.lameleg.com", 31389))
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer l.Close()

	err = l.Bind("cn=admin,dc=sealyun,dc=com", "admin")
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	ldapdn := strings.Split("cn=admin,dc=sealyun,dc=com", ",")
	dc := ldapdn[1] + string(',') + ldapdn[2] // dc=sealyun,dc=com
	searchRequest := ldap.NewSearchRequest(
		dc, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=*)(uid=%s))", user),
		[]string{"dn", "cn"},
		nil,
	)
	sr, err := l.Search(searchRequest)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	if len(sr.Entries) == 0 {
		logger.Error(tools.ErrLdapUserNotExists)
		return nil, tools.ErrLdapUserNotExists
	}
	return sr, nil
}

func authenticationLdap(user, password string) error {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", "fist.lameleg.com", 31389))
	if err != nil {
		logger.Error(err)
		return err
	}
	defer l.Close()

	sr, err := getLdapSearchResult(user)
	if err != nil {
		logger.Error(err)
		return err
	}
	userdn := sr.Entries[0].DN
	// Bind as the user to verify their password
	err = l.Bind(userdn, password)
	if err != nil {
		logger.Error(err)
		return err
	}
	logger.Info("user authenticated")
	return nil
}

func getLdapUserCn(user string) string {
	sr, err := getLdapSearchResult(user)
	if err != nil {
		logger.Error(err)
		return "ldap-" + user
	}
	return sr.Entries[0].GetAttributeValue("cn") //get nickname by ldap cn
}
