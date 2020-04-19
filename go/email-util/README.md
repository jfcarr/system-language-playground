Reports number of unread email messages for one or more accounts, through IMAP.

Copy the email-util-template.json file to email-util.json, then update the email-util.json file with information for the accounts you want to check.  Each account entry looks like this:

```json
{
	"ServerName": "imap_server_name",
	"ServerPort": 993,
	"UserName": "user_name_or_email_address",
	"UserPassword": "user_password",
	"ReportAll": false
}
```

You can add as many account entries as you like.

Field | Description
---------|----------
ServerName | IMAP server address, e.g., **outlook.office365.com**
ServerPort | IMAP server port, e.g., **993**
UserName | User name or email address (depending on what your email provider requires), e.g., **johndoe@outlook.com**
UserPassword | The password for your email account
ReportAll | Do you want the unread email account to be reported even if it's zero?  **true** or **false**

After you build **email-util.go**, place the **email-util** executable and the **email-util.json** file together anywhere you'd like, then run email-util whenever you want to check unread mail counts.
