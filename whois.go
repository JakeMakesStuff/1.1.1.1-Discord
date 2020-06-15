package main

import (
	"context"
	"github.com/andersfylling/disgord"
	"net"
)

// WHOISIP runs the WHOIS request on a IP.
func WHOISIP(args []string, s disgord.Session, m *disgord.Message) {
	ip := args[1]

	// Get the response
	WHOISData, err := FetchWHOISIPJSON(ip)
	if err != nil {
		_, _ = s.SendMsg(context.TODO(), m.ChannelID, "Could not fetch whois information for `"+ip+"` due to an error")
		return
	}
	if WHOISData.ResultsInfo.Count < 1 {
		_, _ = s.SendMsg(context.TODO(), m.ChannelID, "No whois results found for `"+ip+"`")
		return
	}

	// Paginate data and send to channel
	paginator.Paginate(s, m, []string{WrapDataTitle(ip, FormatWHOISIPData(*WHOISData))}, "```\n", "\n```")
}

// WHOISDomain runs the WHOIS request on a domain.
func WHOISDomain(args []string, s disgord.Session, m *disgord.Message) {
	domain := args[1]

	// Get the response
	WHOISData, err := FetchWHOISDomainJSON(domain)
	if err != nil {
		_, _ = s.SendMsg(context.TODO(), m.ChannelID, "Could not fetch whois information for `"+domain+"` due to an error")
		return
	}
	if !WHOISData.Success {
		_, _ = s.SendMsg(context.TODO(), m.ChannelID, "Could not fetch whois information for `"+domain+"` due to an error")
		return
	}

	// Paginate data and send to channel
	paginator.Paginate(s, m, []string{WrapDataTitle(domain, FormatWHOISDomainData(*WHOISData))}, "```\n", "\n```")
}

// WHOIS runs the WHOIS request on either a domain or IP.
func WHOIS(args []string, s disgord.Session, m *disgord.Message) {
	// Validate ip
	ip := args[1]
	if net.ParseIP(ip) == nil {
		WHOISDomain(args, s, m)
		return
	}

	WHOISIP(args, s, m)
}
