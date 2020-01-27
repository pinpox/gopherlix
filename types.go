package main

var itemTypes = map[string]string{

	//Canonical types from the RFC

	// 0   The item is a TextFile Entity.
	//     Client should use a TextFile Transaction.
	"TEXT": "0",

	// 1   The item is a Menu Entity.
	//     Client should use a Menu Transaction.
	"MENU": "1",

	// 2   The information applies to a CSO phone book entity.
	//     Client should talk CSO protocol.
	"CSO": "2",

	// 3   Signals an error condition.
	"ERROR": "3",

	// 4   Item is a Macintosh file encoded in BINHEX format
	"BINHEX": "4",

	// 5   Item is PC-DOS binary file of some sort.  Client gets to decide.
	"PCDOS": "5",

	// 6   Item is a uuencoded file.
	"UUENC": "6",

	// 7   The information applies to a Index Server.
	//     Client should use a FullText Search transaction.
	"FULLTEXT": "7",

	// 8   The information applies to a Telnet session.
	//     Connect to given host at given port. The name to login as at this
	//     host is in the selector string.
	"TELNET": "8",

	// 9   Item is a binary file.  Client must decide what to do with it.
	"BINARY": "9",

	// +   The information applies to a duplicated server.  The information
	//     contained within is a duplicate of the primary server.  The primary
	//     server is defined as the last DirEntity that is has a non-plus
	//     "Type" field.  The client should use the transaction as defined by
	//     the primary server Type field.
	"PRIMARY": "+",

	// g   Item is a GIF graphic file.
	"GIF": "g",

	// I   Item is some kind of image file.  Client gets to decide.
	"IMAGE": "I",

	// T   The information applies to a tn3270 based telnet session.
	//     Connect to given host at given port. The name to login as at this
	//     host is in the selector string.
	"TN3270": "T",

	// Also support non-cannonical types

	// d   Document
	"DOC": "d",

	// h   HTML-File
	"HTML": "h",

	// i   Informational Message
	"INFO": "i",

	// s   Sound file, mostly wav format
	"SOUND": "s",
}
