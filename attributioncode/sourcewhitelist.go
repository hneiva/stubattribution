package attributioncode

import "regexp"

// isWhitelisted returns true if source is in the whitelist
func isWhitelisted(source string) bool {
	if sourceWhitelist[source] {
		return true
	}
	for _, r := range sourceWhitelistRegexps {
		if r.MatchString(source) {
			return true
		}
	}
	return false
}

var sourceWhitelistRegexps = []*regexp.Regexp{
	regexp.MustCompile(`^[\w-\.]*\.moz\.works$`),
	regexp.MustCompile(`^[\w-]*\.allizom\.org$`),
	regexp.MustCompile(`^www\.google(\.com?)?\.\w+$`),
	regexp.MustCompile(`^\w+\.google\.com$`),
	regexp.MustCompile(`^\w+\.search\.yahoo\.com$`),
	regexp.MustCompile(`^\w+\.facebook\.com$`),
	regexp.MustCompile(`^[\w-]+\.wikipedia\.org$`),
	regexp.MustCompile(`^(\w+\.)*firefox\.com$`),
	regexp.MustCompile(`^([\w-]*\.)?mozilla\.org$`),
}

var sourceWhitelist = map[string]bool{
	"about-home":        true,
	"aol":               true,
	"ask":               true,
	"bing":              true,
	"desktop-snippet":   true,
	"directory-tiles":   true,
	"firefox-dev-tools": true,
	"google":            true,
	"seznam":            true,
	"snippet":           true,
	"yahoo":             true,
	"yandex":            true,

	"firefox-com": true,

	"getfirefox-com":     true,
	"getfirefox.com":     true,
	"www.getfirefox.com": true,

	"accounts.firefox.com":           true,
	"activations.cdn.mozilla.net":    true,
	"answers.yahoo.com":              true,
	"ar.search.yahoo.com":            true,
	"at.search.yahoo.com":            true,
	"au.search.yahoo.com":            true,
	"bienvenido.toshiba.com":         true,
	"bing.com":                       true,
	"br.answers.yahoo.com":           true,
	"br.search.yahoo.com":            true,
	"br.yhs4.search.yahoo.com":       true,
	"ca.search.yahoo.com":            true,
	"ch.search.yahoo.com":            true,
	"chip.de":                        true,
	"cl.search.yahoo.com":            true,
	"cn.bing.com":                    true,
	"co.search.yahoo.com":            true,
	"co.yhs4.search.yahoo.com":       true,
	"cse.google.com":                 true,
	"cto.mail.ru":                    true,
	"de.search.yahoo.com":            true,
	"dk.search.yahoo.com":            true,
	"duckduckgo.com":                 true,
	"e.mail.ru":                      true,
	"email.seznam.cz":                true,
	"en-maktoob.search.yahoo.com":    true,
	"encrypted.google.com":           true,
	"es-mg42.mail.yahoo.com":         true,
	"es.search.yahoo.com":            true,
	"espanol.search.yahoo.com":       true,
	"extensions.aol.com":             true,
	"facebook":                       true,
	"facebook.com":                   true,
	"fi.search.yahoo.com":            true,
	"firefox":                        true,
	"firefox-browser":                true,
	"firefox.cz":                     true,
	"firefox.de":                     true,
	"firefox.mozilla.cz":             true,
	"firefox.no":                     true,
	"firefox.org":                    true,
	"firefox.si":                     true,
	"fr-mg42.mail.yahoo.com":         true,
	"fr.search.yahoo.com":            true,
	"fr.yhs4.search.yahoo.com":       true,
	"fx36start":                      true,
	"global.bing.com":                true,
	"go.mail.ru":                     true,
	"gr.search.yahoo.com":            true,
	"hangouts.google.com":            true,
	"help.ea.com":                    true,
	"help.mail.ru":                   true,
	"hk.messenger.yahoo.com":         true,
	"hk.search.yahoo.com":            true,
	"id.messenger.yahoo.com":         true,
	"id.search.yahoo.com":            true,
	"id.yhs4.search.yahoo.com":       true,
	"images.tanks.mail.ru":           true,
	"in.search.yahoo.com":            true,
	"in.yhs4.search.yahoo.com":       true,
	"it.search.yahoo.com":            true,
	"kongregate.com":                 true,
	"lite.qwant.com":                 true,
	"love.mail.ru":                   true,
	"mail.aol.com":                   true,
	"mail.de":                        true,
	"mail.google.com":                true,
	"mail.ru":                        true,
	"maktoob.search.yahoo.com":       true,
	"malaysia.search.yahoo.com":      true,
	"malaysia.yhs4.search.yahoo.com": true,
	"messenger.yahoo.com":            true,
	"mg.mail.yahoo.com":              true,
	"mozilla.ch":                     true,
	"mozilla.com":                    true,
	"mozilla.cz":                     true,
	"mozilla.de":                     true,
	"mozilla.ee":                     true,
	"mozilla.fi":                     true,
	"mozilla.hu":                     true,
	"mozilla.jp":                     true,
	"mozilla.lt":                     true,
	"mozilla.pl":                     true,
	"mozilla.ro":                     true,
	"mozilla.rs":                     true,
	"mozilla.si":                     true,
	"mozilla.sk":                     true,
	"mx.search.yahoo.com":            true,
	"mx.yhs4.search.yahoo.com":       true,
	"my.mail.ru":                     true,
	"myaccount.google.com":           true,
	"navigator-bs.gmx.com":           true,
	"navigator-bs.gmx.es":            true,
	"navigator-bs.gmx.fr":            true,
	"nl.search.yahoo.com":            true,
	"no.search.yahoo.com":            true,
	"nz.search.yahoo.com":            true,
	"otvet.mail.ru":                  true,
	"partnerads.ysm.yahoo.com":       true,
	"pe.search.yahoo.com":            true,
	"ph.search.yahoo.com":            true,
	"photos.google.com":              true,
	"pl.search.yahoo.com":            true,
	"plus.google.com":                true,
	"plus.url.google.com":            true,
	"poseidon.navigator-bs.gmx.com":  true,
	"qc.search.yahoo.com":            true,
	"r.search.yahoo.com":             true,
	"ro.search.yahoo.com":            true,
	"ru.search.yahoo.com":            true,
	"scholar.google.com":             true,
	"se.search.yahoo.com":            true,
	"se.yhs4.search.yahoo.com":       true,
	"search.1and1.com":               true,
	"search.seznam.cz":               true,
	"search.yahoo.com":               true,
	"search.yahoo.co.jp":             true,
	"sg.search.yahoo.com":            true,
	"softonic.com":                   true,
	"start.new.toshiba.com":          true,
	"start.toshiba.com":              true,
	"suche.gmx.at":                   true,
	"suche.gmx.net":                  true,
	"support.google.com":             true,
	"takeout.google.com":             true,
	"talkgadget.google.com":          true,
	"tanks.mail.ru":                  true,
	"taobao.com":                     true,
	"th.search.yahoo.com":            true,
	"thunderbird.mozilla.cz":         true,
	"toshiba.com":                    true,
	"tr.search.yahoo.com":            true,
	"tw.search.yahoo.com":            true,
	"tweetdeck.twitter.com":          true,
	"uk.search.yahoo.com":            true,
	"uk.yhs4.search.yahoo.com":       true,
	"us-mg5.mail.yahoo.com":          true,
	"us-mg6.mail.yahoo.com":          true,
	"us.search.yahoo.com":            true,
	"us.yhs4.search.yahoo.com":       true,
	"ve.search.yahoo.com":            true,
	"vn.search.yahoo.com":            true,
	"www.aol.com":                    true,
	"www.bing.com":                   true,
	"www.chip.de":                    true,
	"www.google.be":                  true,
	"www.google.bg":                  true,
	"www.google.ca":                  true,
	"www.google.com":                 true,
	"www.google.de":                  true,
	"www.google.dz":                  true,
	"www.google.es":                  true,
	"www.google.fr":                  true,
	"www.google.it":                  true,
	"www.google.pl":                  true,
	"www.google.ro":                  true,
	"www.google.se":                  true,
	"www.google.sr":                  true,
	"www.kongregate.com":             true,
	"www.miniclip.com":               true,
	"www.qwant.com":                  true,
	"www.seznam.cz":                  true,
	"www.softonic.com":               true,
	"www.toshiba.com":                true,
	"www.yahoo.com":                  true,
	"www.youtube.com":                true,
	"yandex.ru":                      true,
	"youtube.com":                    true,
}
