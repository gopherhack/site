package lib

import "gopkg.in/kothar/go-backblaze.v0"


b2, _ := backblaze.NewB2(backblaze.Credentials{
	AccountID:      accountID,
	ApplicationKey: applicationKey,
})

//Account ID: 4eb526a102db
//Application Key: 001813fcb7c5f9535da6703227ff6f0b046af7b918
