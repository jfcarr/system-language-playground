Reports number of unread email messages for one or more accounts, through IMAP.

Update the email-util.config file with entries for the accounts you want to check.  Entries look like this:

imap_server_name|imap_server_port|user_name_or_email_address|password|report_zero_new_messages_true_or_false


Field | Description
---------|----------
imap_server_name | IMAP server address, e.g., **outlook.office365.com**
imap_server_port | IMAP server port, e.g., **993**
user_name_or_email_address | User name or email address (depending on what your email provider requires), e.g., **johndoe@outlook.com**
password | The password for your email account
report_zero_new_messages_true_or_false | Do you want the unread email account to be reported even if it's zero?  **true** or **false**

(Entries are pipe-separated)

After you build **email-util.go**, place the **email-util** executable and the **email-util.config** file together anywhere you'd like, then run email-util whenever you want to check unread mail counts.
