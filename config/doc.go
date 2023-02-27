/*
Package config holds the configuration file definitions for mox.conf (Static)
and domains.conf (Dynamic).

Annotated empty/default configuration files you could use as a starting point
for your mox.conf and domains.conf, as generated by "mox config
describe-static" and "mox config describe-domains":

# mox.conf

	# Directory where all data is stored, e.g. queue, accounts and messages, ACME TLS
	# certs/keys. If this is a relative path, it is relative to the directory of
	# mox.conf.
	DataDir:

	# Default log level, one of: error, info, debug, trace, traceauth, tracedata.
	# Trace logs SMTP and IMAP protocol transcripts, with traceauth also messages with
	# passwords, and tracedata on top of that also the full data exchanges (full
	# messages), which can be a large amount of data.
	LogLevel:

	# Overrides of log level per package (e.g. queue, smtpclient, smtpserver,
	# imapserver, spf, dkim, dmarc, dmarcdb, autotls, junk, mtasts, tlsrpt).
	# (optional)
	PackageLogLevels:
		x:

	# User to switch to after binding to all sockets as root. Default: mox. If the
	# value is not a known user, it is parsed as integer and used as uid and gid.
	# (optional)
	User:

	# If true, do not automatically fix file permissions when starting up. By default,
	# mox will ensure reasonable owner/permissions on the working, data and config
	# directories (and files), and mox binary (if present). (optional)
	NoFixPermissions: false

	# Full hostname of system, e.g. mail.<domain>
	Hostname:

	# If enabled, a single DNS TXT lookup of _updates.xmox.nl is done every 24h to
	# check for a new release. Each time a new release is found, a changelog is
	# fetched from https://updates.xmox.nl and delivered to the postmaster mailbox.
	# (optional)
	CheckUpdates: false

	# Global TLS configuration, e.g. for additional Certificate Authorities.
	# (optional)
	TLS:

		# (optional)
		CA:

			# (optional)
			AdditionalToSystem: false

			# (optional)
			CertFiles:
				-

	# Automatic TLS configuration with ACME, e.g. through Let's Encrypt. The key is a
	# name referenced in TLS configs, e.g. letsencrypt. (optional)
	ACME:
		x:

			# For letsencrypt, use https://acme-v02.api.letsencrypt.org/directory.
			DirectoryURL:

			# How long before expiration to renew the certificate. Default is 30 days.
			# (optional)
			RenewBefore: 0s

			# Email address to register at ACME provider. The provider can email you when
			# certificates are about to expire. If you configure an address for which email is
			# delivered by this server, keep in mind that TLS misconfigurations could result
			# in such notification emails not arriving.
			ContactEmail:

			# TLS port for ACME validation, 443 by default. You should only override this if
			# you cannot listen on port 443 directly. ACME will make requests to port 443, so
			# you'll have to add an external mechanism to get the connection here, e.g. by
			# configuring port forwarding. (optional)
			Port: 0

	# File containing hash of admin password, for authentication in the web admin
	# pages (if enabled). (optional)
	AdminPasswordFile:

	# Listeners are groups of IP addresses and services enabled on those IP addresses,
	# such as SMTP/IMAP or internal endpoints for administration or Prometheus
	# metrics. All listeners with SMTP/IMAP services enabled will serve all configured
	# domains.
	Listeners:
		x:

			# Use 0.0.0.0 to listen on all IPv4 and/or :: to listen on all IPv6 addresses.
			IPs:
				-

			# If empty, the config global Hostname is used. (optional)
			Hostname:

			# For SMTP/IMAP STARTTLS, direct TLS and HTTPS connections. (optional)
			TLS:

				# Name of provider from top-level configuration to use for ACME, e.g. letsencrypt.
				# (optional)
				ACME:

				# (optional)
				KeyCerts:
					-

						# Certificate including intermediate CA certificates, in PEM format.
						CertFile:

						# Private key for certificate, in PEM format. PKCS8 is recommended, but PKCS1 and
						# EC private keys are recognized as well.
						KeyFile:

				# Minimum TLS version. Default: TLSv1.2. (optional)
				MinVersion:

			# Maximum size in bytes accepted incoming and outgoing messages. Default is 100MB.
			# (optional)
			SMTPMaxMessageSize: 0

			# (optional)
			SMTP:
				Enabled: false

				# Default 25. (optional)
				Port: 0

				# Do not offer STARTTLS to secure the connection. Not recommended. (optional)
				NoSTARTTLS: false

				# Do not accept incoming messages if STARTTLS is not active. Can be used in
				# combination with a strict MTA-STS policy. A remote SMTP server may not support
				# TLS and may not be able to deliver messages. (optional)
				RequireSTARTTLS: false

				# Addresses of DNS block lists for incoming messages. Block lists are only
				# consulted for connections/messages without enough reputation to make an
				# accept/reject decision. This prevents sending IPs of all communications to the
				# block list provider. If any of the listed DNSBLs contains a requested IP
				# address, the message is rejected as spam. The DNSBLs are checked for healthiness
				# before use, at most once per 4 hours. Example DNSBLs: sbl.spamhaus.org,
				# bl.spamcop.net (optional)
				DNSBLs:
					-

			# SMTP for submitting email, e.g. by email applications. Starts out in plain text,
			# can be upgraded to TLS with the STARTTLS command. Prefer using Submissions which
			# is always a TLS connection. (optional)
			Submission:
				Enabled: false

				# Default 587. (optional)
				Port: 0

				# Do not require STARTTLS. Since users must login, this means password may be sent
				# without encryption. Not recommended. (optional)
				NoRequireSTARTTLS: false

			# SMTP over TLS for submitting email, by email applications. Requires a TLS
			# config. (optional)
			Submissions:
				Enabled: false

				# Default 465. (optional)
				Port: 0

			# IMAP for reading email, by email applications. Starts out in plain text, can be
			# upgraded to TLS with the STARTTLS command. Prefer using IMAPS instead which is
			# always a TLS connection. (optional)
			IMAP:
				Enabled: false

				# Default 143. (optional)
				Port: 0

				# Enable this only when the connection is otherwise encrypted (e.g. through a
				# VPN). (optional)
				NoRequireSTARTTLS: false

			# IMAP over TLS for reading email, by email applications. Requires a TLS config.
			# (optional)
			IMAPS:
				Enabled: false

				# Default 993. (optional)
				Port: 0

			# Account web interface, for email users wanting to change their accounts, e.g.
			# set new password, set new delivery rulesets. (optional)
			AccountHTTP:
				Enabled: false

				# Default 80. (optional)
				Port: 0

			# Account web interface listener for HTTPS. Requires a TLS config. (optional)
			AccountHTTPS:
				Enabled: false

				# Default 80. (optional)
				Port: 0

			# Admin web interface, for managing domains, accounts, etc. Served at /admin/.
			# Preferrably only enable on non-public IPs. (optional)
			AdminHTTP:
				Enabled: false

				# Default 80. (optional)
				Port: 0

			# Admin web interface listener for HTTPS. Requires a TLS config. Preferrably only
			# enable on non-public IPs. (optional)
			AdminHTTPS:
				Enabled: false

				# Default 443. (optional)
				Port: 0

			# Serve prometheus metrics, for monitoring. You should not enable this on a public
			# IP. (optional)
			MetricsHTTP:
				Enabled: false

				# Default 8010. (optional)
				Port: 0

			# Serve /debug/pprof/ for profiling a running mox instance. Do not enable this on
			# a public IP! (optional)
			PprofHTTP:
				Enabled: false

				# Default 8011. (optional)
				Port: 0

			# Serve autoconfiguration/autodiscovery to simplify configuring email
			# applications, will use port 443. Requires a TLS config. (optional)
			AutoconfigHTTPS:
				Enabled: false

				# TLS port, 443 by default. You should only override this if you cannot listen on
				# port 443 directly. Autoconfig requests will be made to port 443, so you'll have
				# to add an external mechanism to get the connection here, e.g. by configuring
				# port forwarding. (optional)
				Port: 0

				# If set, plain HTTP instead of HTTPS is spoken on the configured port. Can be
				# useful when the autoconfig domain is reverse proxied. (optional)
				NonTLS: false

			# Serve MTA-STS policies describing SMTP TLS requirements. Requires a TLS config.
			# (optional)
			MTASTSHTTPS:
				Enabled: false

				# TLS port, 443 by default. You should only override this if you cannot listen on
				# port 443 directly. MTA-STS requests will be made to port 443, so you'll have to
				# add an external mechanism to get the connection here, e.g. by configuring port
				# forwarding. (optional)
				Port: 0

				# If set, plain HTTP instead of HTTPS is spoken on the configured port. Can be
				# useful when the mta-sts domain is reverse proxied. (optional)
				NonTLS: false

	# Destination for emails delivered to postmaster address.
	Postmaster:
		Account:

		# E.g. Postmaster or Inbox.
		Mailbox:

	# Mailboxes to create when adding an account. Inbox is always created. If no
	# mailboxes are specified, the following are automatically created: Sent, Archive,
	# Trash, Drafts and Junk. (optional)
	DefaultMailboxes:
		-

# domains.conf

	# Domains for which email is accepted. For internationalized domains, use their
	# IDNA names in UTF-8.
	Domains:
		x:

			# Free-form description of domain. (optional)
			Description:

			# If not empty, only the string before the separator is used to for email delivery
			# decisions. For example, if set to "+", you+anything@example.com will be
			# delivered to you@example.com. (optional)
			LocalpartCatchallSeparator:

			# If set, upper/lower case is relevant for email delivery. (optional)
			LocalpartCaseSensitive: false

			# With DKIM signing, a domain is taking responsibility for (content of) emails it
			# sends, letting receiving mail servers build up a (hopefully positive) reputation
			# of the domain, which can help with mail delivery. (optional)
			DKIM:

				# Emails can be DKIM signed. Config parameters are per selector. A DNS record must
				# be created for each selector. Add the name to Sign to use the selector for
				# signing messages.
				Selectors:
					x:

						# sha256 (default) or (older, not recommended) sha1 (optional)
						Hash:

						# (optional)
						Canonicalization:

							# If set, some modifications to the headers (mostly whitespace) are allowed.
							HeaderRelaxed: false

							# If set, some whitespace modifications to the message body are allowed.
							BodyRelaxed: false

						# Headers to sign with DKIM. If empty, a reasonable default set of headers is
						# selected. (optional)
						Headers:
							-

						# If set, don't prevent duplicate headers from being added. Not recommended.
						# (optional)
						DontSealHeaders: false

						# Period a signature is valid after signing, as duration, e.g. 72h. The period
						# should be enough for delivery at the final destination, potentially with several
						# hops/relays. In the order of days at least. (optional)
						Expiration:

						# Either an RSA or ed25519 private key file in PKCS8 PEM form.
						PrivateKeyFile:

				# List of selectors that emails will be signed with. (optional)
				Sign:
					-

			# With DMARC, a domain publishes, in DNS, a policy on how other mail servers
			# should handle incoming messages with the From-header matching this domain and/or
			# subdomain (depending on the configured alignment). Receiving mail servers use
			# this to build up a reputation of this domain, which can help with mail delivery.
			# A domain can also publish an email address to which reports about DMARC
			# verification results can be sent by verifying mail servers, useful for
			# monitoring. Incoming DMARC reports are automatically parsed, validated, added to
			# metrics and stored in the reporting database for later display in the admin web
			# pages. (optional)
			DMARC:

				# Address-part before the @ that accepts DMARC reports. Must be
				# non-internationalized. Recommended value: dmarc-reports.
				Localpart:

				# Account to deliver to.
				Account:

				# Mailbox to deliver to, e.g. DMARC.
				Mailbox:

			# With MTA-STS a domain publishes, in DNS, presence of a policy for
			# using/requiring TLS for SMTP connections. The policy is served over HTTPS.
			# (optional)
			MTASTS:

				# Policies are versioned. The version must be specified in the DNS record. If you
				# change a policy, first change it in mox, then update the DNS record.
				PolicyID:

				# testing, enforce or none. If set to enforce, a remote SMTP server will not
				# deliver email to us if it cannot make a TLS connection.
				Mode:

				# How long a remote mail server is allowed to cache a policy. Typically 1 or
				# several weeks.
				MaxAge: 0s

				# List of server names allowed for SMTP. If empty, the configured hostname is set.
				# Host names can contain a wildcard (*) as a leading label (matching a single
				# label, e.g. *.example matches host.example, not sub.host.example). (optional)
				MX:
					-

			# With TLSRPT a domain specifies in DNS where reports about encountered SMTP TLS
			# behaviour should be sent. Useful for monitoring. Incoming TLS reports are
			# automatically parsed, validated, added to metrics and stored in the reporting
			# database for later display in the admin web pages. (optional)
			TLSRPT:

				# Address-part before the @ that accepts TLSRPT reports. Recommended value:
				# tls-reports.
				Localpart:

				# Account to deliver to.
				Account:

				# Mailbox to deliver to, e.g. TLSRPT.
				Mailbox:

	# Accounts to which email can be delivered. An account can accept email for
	# multiple domains, for multiple localparts, and deliver to multiple mailboxes.
	Accounts:
		x:

			# Default domain for addresses specified in Destinations. An address can specify a
			# domain override.
			Domain:

			# Free form description, e.g. full name or alternative contact info. (optional)
			Description:

			# Destinations, specified as (encoded) localpart for Domain, or a full address
			# including domain override.
			Destinations:
				x:

					# Mailbox to deliver to if none of Rulesets match. Default: Inbox. (optional)
					Mailbox:

					# Delivery rules based on message and SMTP transaction. You may want to match each
					# mailing list by SMTP MailFrom address, VerifiedDomain and/or List-ID header
					# (typically <listname.example.org> if the list address is listname@example.org),
					# delivering them to their own mailbox. (optional)
					Rulesets:
						-

							# Matches if this regular expression matches (a substring of) the SMTP MAIL FROM
							# address (not the message From-header). E.g. user@example.org. (optional)
							SMTPMailFromRegexp:

							# Matches if this domain matches an SPF- and/or DKIM-verified (sub)domain.
							# (optional)
							VerifiedDomain:

							# Matches if these header field/value regular expressions all match (substrings
							# of) the message headers. Header fields and valuees are converted to lower case
							# before matching. Whitespace is trimmed from the value before matching. A header
							# field can occur multiple times in a message, only one instance has to match. For
							# mailing lists, you could match on ^list-id$ with the value typically the mailing
							# list address in angled brackets with @ replaced with a dot, e.g.
							# <name\.lists\.example\.org>. (optional)
							HeadersRegexp:
								x:

							# Influence the spam filtering, this does not change whether this ruleset applies
							# to a message. If this domain matches an SPF- and/or DKIM-verified (sub)domain,
							# the message is accepted without further spam checks, such as a junk filter or
							# DMARC reject evaluation. DMARC rejects should not apply for mailing lists that
							# are not configured to rewrite the From-header of messages that don't have a
							# passing DKIM signature of the From-domain. Otherwise, by rejecting messages, you
							# may be automatically unsubscribed from the mailing list. The assumption is that
							# mailing lists do their own spam filtering/moderation. (optional)
							ListAllowDomain:

							# Mailbox to deliver to if this ruleset matches.
							Mailbox:

			# If configured, messages classified as weakly spam are rejected with instructions
			# to retry delivery, but this time with a signed token added to the subject.
			# During the next delivery attempt, the signed token will bypass the spam filter.
			# Messages with a clear spam signal, such as a known bad reputation, are
			# rejected/delayed without a signed token. (optional)
			SubjectPass:

				# How long unique values are accepted after generating, e.g. 12h.
				Period: 0s

			# Mail that looks like spam will be rejected, but a copy can be stored temporarily
			# in a mailbox, e.g. Rejects. If mail isn't coming in when you expect, you can
			# look there. The mail still isn't accepted, so the remote mail server may retry
			# (hopefully, if legitimate), or give up (hopefully, if indeed a spammer).
			# Messages are automatically removed from this mailbox, so do not set it to a
			# mailbox that has messages you want to keep. (optional)
			RejectsMailbox:

			# Automatically set $Junk and $NotJunk flags based on mailbox messages are
			# delivered/moved/copied to. Email clients typically have too limited
			# functionality to conveniently set these flags, especially $NonJunk, but they can
			# all move messages to a different mailbox, so this helps them. (optional)
			AutomaticJunkFlags:

				# If enabled, flags will be set automatically if they match a regular expression
				# below. When two of the three mailbox regular expressions are set, the remaining
				# one will match all unmatched messages. Messages are matched in the order
				# specified and the search stops on the first match. Mailboxes are lowercased
				# before matching.
				Enabled: false

				# Example: ^(junk|spam). (optional)
				JunkMailboxRegexp:

				# Example: ^(inbox|neutral|postmaster|dmarc|tlsrpt|rejects), and you may wish to
				# add trash depending on how you use it, or leave this empty. (optional)
				NeutralMailboxRegexp:

				# Example: .* or an empty string. (optional)
				NotJunkMailboxRegexp:

			# Content-based filtering, using the junk-status of individual messages to rank
			# words in such messages as spam or ham. It is recommended you always set the
			# applicable (non)-junk status on messages, and that you do not empty your Trash
			# because those messages contain valuable ham/spam training information.
			# (optional)
			JunkFilter:

				# Approximate spaminess score between 0 and 1 above which emails are rejected as
				# spam. Each delivery attempt adds a little noise to make it slightly harder for
				# spammers to identify words that strongly indicate non-spaminess and use it to
				# bypass the filter. E.g. 0.95.
				Threshold: 0.000000
				Params:

					# Track ham/spam ranking for single words. (optional)
					Onegrams: false

					# Track ham/spam ranking for each two consecutive words. (optional)
					Twograms: false

					# Track ham/spam ranking for each three consecutive words. (optional)
					Threegrams: false

					# Maximum power a word (combination) can have. If spaminess is 0.99, and max power
					# is 0.1, spaminess of the word will be set to 0.9. Similar for ham words.
					MaxPower: 0.000000

					# Number of most spammy/hammy words to use for calculating probability. E.g. 10.
					TopWords: 0

					# Ignore words that are this much away from 0.5 haminess/spaminess. E.g. 0.1,
					# causing word (combinations) of 0.4 to 0.6 to be ignored. (optional)
					IgnoreWords: 0.000000

					# Occurrences in word database until a word is considered rare and its influence
					# in calculating probability reduced. E.g. 1 or 2. (optional)
					RareWords: 0
*/
package config

// NOTE: DO NOT EDIT, this file is generated by ../gendoc.sh.
