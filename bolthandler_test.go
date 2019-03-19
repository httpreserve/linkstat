package main

/* Example JSON:
{
   "FileName": "bbc news",
   "AnalysisVersionNumber": "0.0.0",
   "AnalysisVersionText": "exponentialDK-httpreserve/0.0.0",
   "Link": "http://www.bbc.co.uk/news",
   "ResponseCode": 200,
   "ResponseText": "OK",
   "ScreenShot": "",
   "InternetArchiveLinkLatest": "http://web.archive.org/web/20170328040059/http://www.bbc.co.uk/news/",
   "InternetArchiveLinkEarliest": "http://web.archive.org/web/19971009011901/http://www.bbc.co.uk/news/",
   "InternetArchiveSaveLink": "http://web.archive.org/save/http://www.bbc.co.uk/news",
   "InternetArchiveResponseCode": 200,
   "InternetArchiveResponseText": "OK",
   "Archived": true,
   "ProtocolError": false,
   "ProtocolErrorMessage": ""
},
*/

/*
filename:bbc home
id:891609239375c54fe326a2e23a8c5397

filename:bbc radio
id:1d15698856a2487bade7d8994d21d30c

filename:tna
id:3357924f215d974f627690dd6382076c

id:43d4a499caa590e912c8a059f7ab8323
filename:bbc news */

/*
//for now, for testing...
var linkmap = map[string]string{
   "http://www.bbc.co.uk/news":           "bbc news",
   "http://www.bbc.co.uk/":               "bbc home",
   "http://www.bbc.co.uk/radio":          "bbc radio",
   "http://www.nationalarchives.gov.uk/": "tna",
}
*/
