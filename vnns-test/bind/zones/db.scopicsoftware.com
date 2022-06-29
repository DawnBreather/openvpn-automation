$TTL	604800
@	IN	SOA	vnns.scopicsoftware.com. admin.scopicsoftware.com. (
			 90		; Serial	; Serial	; Serial	; Serial	; Serial	; Serial	; Serial	; Serial	; Serial	; Serial	; Serial
			 604800		; Refresh
			86400		; Retry
			2419200		; Expire
			 604800 )	; Negative Cache TTL
;

	IN	NS	vnns.scopicsoftware.com.

; name servers - A records
vnns.scopicsoftware.com.				IN	A	10.253.255.122 ;10.252.186.254

; hosts
zoho-pmportal.scopicsoftware.com.			IN	A	10.253.253.210
jira.scopicsoftware.com.				IN	A	10.253.255.199
chp.scopicsoftawre.com.					IN	A	10.253.255.196
api-cpt.scopicsoftware.com.				IN	A	10.253.253.211
cpt.scopicsoftware.com.					IN	A	10.253.253.211
time.scopicsoftware.com.				IN	A	10.253.255.207 ;10.252.93.66
test-site.scopicsoftware.com.				IN	A	10.253.255.122 ;10.252.186.254
vpn-status.scopicsoftware.com.				IN	A	10.253.255.122 ;10.252.186.254
portal.scopicsoftware.com.				IN	A	10.253.253.238 ;10.252.28.148
prst.scopicsoftware.com.				IN	A	10.253.255.139 ;10.252.112.216
prn.scopicsoftware.com.					IN	A	10.253.255.202 ;10.252.205.107
new-rmt.scopicsoftware.com.				IN	A	10.253.253.212 ;10.252.185.156
rmt.scopicsoftware.com.					IN	A	10.253.253.212 ;10.252.185.156
invoicing-tool.scopicsoftware.com.			IN	A	10.253.255.135 ;10.252.245.132
escrow-alert-tool.scopicsoftware.com.			IN	A	10.253.255.128 ;10.252.254.176
bonus-tool.scopicsoftware.com.				IN	A	10.253.253.244 ;10.252.174.13
bonus.scopicdev.com					IN	A	10.253.253.244 ;10.252.174.13
bonus-dev.scopicdev.com					IN	A	10.253.253.244 ;10.252.174.13
login-location-tracking.scopicsoftware.com.		IN	A	10.253.255.126 ;10.252.207.4
recruiter-box.scopicsoftware.com.			IN	A	10.253.255.116 ;10.252.223.121
rama.scopicsoftware.com.				IN	A	10.253.255.121 ;10.252.43.86
texting.scopicsoftware.com.				IN	A	10.253.255.119
texting-app.scopicsoftware.com.				IN	A	10.253.255.119 ;10.252.1.207
ttet.scopicsoftware.com.				IN	A	10.253.255.118 ;10.252.208.48
interview-reminder.scopicsoftware.com.			IN	A	10.253.255.116 ;10.252.223.121
jira-timetask.scopicsoftware.com.			IN	A	10.253.255.113 ;10.252.147.43
job-posting-tool.scopicsoftware.com.			IN	A	10.253.255.111 ;10.252.72.113
acuity-scheduling.scopicsoftware.com.			IN	A	10.253.255.116 ;10.252.223.121
brm.scopicsoftware.com.					IN	A	10.253.255.111 ;10.252.72.113
hours-deficit-tool.scopicsoftware.com.			IN	A	10.253.255.196 ;10.252.50.130
vacation-tracking.scopicsoftware.com.			IN	A	10.253.255.196 ;10.252.50.130
teamwork2bamboohr.scopicsoftware.com.			IN	A	10.253.255.109 ;10.252.212.129
teamgant.scopicsoftware.com.				IN	A	10.253.255.118 ;10.252.208.48
beanstalk-activity-tracker.scopicsoftware.com.		IN	A	10.253.255.115 ;10.252.72.120
api-tt.scopicsoftware.com.				IN	A	10.253.255.207 ;10.252.93.66
recruiting-activity-report.scopicsoftware.com.		IN	A	10.253.255.108 ;10.252.184.193
cv-review-tool.scopicsoftware.com.			IN	A	10.253.255.107 ;10.252.50.167
tw2tt.scopicsoftware.com.				IN	A	10.253.255.111 ;10.252.72.113
otl.scopicsoftware.com.					IN	A	10.253.255.129 ;10.252.30.113
scopicsoftware.com.					IN	A	34.198.193.106
ewt.scopicsoftware.com.					IN	A	10.253.255.200 ;10.252.72.215
jpat.scopicsoftware.com.				IN	A	10.253.253.243 ;10.252.168.123
locationtracking.scopicsoftware.com.			IN	A	10.253.253.242 ;10.252.21.186
scopicapi.scopicsoftware.com.				IN	A	10.253.253.235 ;10.252.117.24
git-backups-verificator.scopicsoftware.com.		IN	A	10.253.253.213 ;10.252.187.233
sft.scopicsoftware.com.					IN	A	10.253.253.208 ;10.252.184.85
bs.scopicsoftware.com.					IN	A	10.253.253.239 ;10.252.248.75
scopicsoftware.com.		IN		MX		10 mx.zoho.com.
scopicsoftware.com.		IN		MX		20 mx2.zoho.com.
2109728.scopicsoftware.com.		IN		A		167.89.123.54
2109728.scopicsoftware.com.		IN		A		167.89.115.56
2109728.scopicsoftware.com.		IN		A		167.89.118.52
_d4d9b59e5cf28d49b18ab449b7ca0db5.scopicsoftware.com.		IN		CNAME		9D26DBB3A64A89CD67691A4EC8B64086.FDE2E38BA02AE1EB7931BDE24E2D36B6.59bb62f5545ec.comodoca.com.
s1._domainkey.scopicsoftware.com.		IN		CNAME		s1.domainkey.u2109728.wl229.sendgrid.net.
s2._domainkey.scopicsoftware.com.		IN		CNAME		s2.domainkey.u2109728.wl229.sendgrid.net.
adobeclientcontract.scopicsoftware.com.		IN		A		3.213.243.13
adobecontractexport.scopicsoftware.com.		IN		A		3.92.38.115
awx.scopicsoftware.com.		IN		A		3.214.208.222
backup-asb.scopicsoftware.com.		IN		A		34.206.145.136
car.scopicsoftware.com.		IN		A		35.174.57.164
careers.scopicsoftware.com.		IN		A		34.198.193.106
cb.scopicsoftware.com.		IN		A		52.72.134.241
chp-test.scopicsoftware.com.		IN		CNAME		payroll.scopicsoftware.com.
chp.scopicsoftware.com.		IN		A		104.236.119.146
chroma-dev.scopicsoftware.com.		IN		A		3.130.17.14
dev-dssp.scopicsoftware.com.		IN		A		3.87.85.56
dev-ewt.scopicsoftware.com.		IN		A		18.204.143.118
dev-jenkins-dssp.scopicsoftware.com.		IN		A		23.20.63.43
dev-server-jenkins-dssp.scopicsoftware.com.		IN		A		23.20.63.43
dssp.scopicsoftware.com.		IN		A		3.215.237.216
get.scopicsoftware.com.		IN		A		18.206.69.142
gst-test.scopicsoftware.com.		IN		A		52.5.118.142
gst.scopicsoftware.com.		IN		A		52.5.118.142
jenkins.scopicsoftware.com.		IN		A		18.204.118.140
_06d46fb28d3ff2317d574c137bf499ed.kreo.scopicsoftware.com.		IN		CNAME		_385b51578af7e5e81a5c2c2d5be40abd.tljzshvwok.acm-validations.aws.
_05634ee7992cc348b1c0b303de8cc63d.server.kreo.scopicsoftware.com.		IN		CNAME		_ea298fd09c2a895611d819bc277d75cc.tljzshvwok.acm-validations.aws.
link.scopicsoftware.com.		IN		A		167.89.115.56
link.scopicsoftware.com.		IN		A		167.89.118.52
link.scopicsoftware.com.		IN		A		167.89.123.54
mail.scopicsoftware.com.		IN		A		34.198.193.106
marketing.scopicsoftware.com.		IN		CNAME		u2109728.wl229.sendgrid.net.
monitor.scopicsoftware.com.		IN		A		35.168.130.65
payroll.scopicsoftware.com.		IN		A		34.227.122.161
rainloop.scopicsoftware.com.		IN		A		35.175.11.46
rb-filtering-tool.scopicsoftware.com.		IN		A		18.235.172.129
rrs.scopicsoftware.com.		IN		A		34.239.191.170
server-dev-dssp.scopicsoftware.com.		IN		A		3.87.85.56
server-dssp.scopicsoftware.com.		IN		A		3.215.237.216
skypebot.scopicsoftware.com.		IN		CNAME		prod.skypebot.scopicsoftware.com.
prod.skypebot.scopicsoftware.com.		IN		A		45.55.76.228
stage.scopicsoftware.com.		IN		A		35.174.57.164
sts2.scopicsoftware.com.		IN		A		34.238.35.31
test-rmt.scopicsoftware.com.		IN		A		3.93.127.197
testing-record.scopicsoftware.com.		IN		A		1.1.1.4
twa.scopicsoftware.com.		IN		A		3.219.101.149
twassistant.scopicsoftware.com.		IN		A		18.233.28.230
www.scopicsoftware.com.		IN		A		34.198.193.106
kreo.scopicsoftware.com.		IN		A		52.48.110.166
kreo.scopicsoftware.com.		IN		A		52.51.88.148
presence.kreo.scopicsoftware.com.		IN		A		52.48.110.166
kreo.scopicsoftware.com.		IN		A		52.51.88.148
server.kreo.scopicsoftware.com.		IN		A		52.48.110.166
kreo.scopicsoftware.com.		IN		A		52.51.88.148


