<div>
<p align="center">
<img height="300px" width="300px" id="logo" src="https://github.com/httpreserve/httpreserve/raw/main/src/images/httpreserve-logo.png" alt="httpreserve"/>
</p>
</div>

# linkstat

CLI implementation of httpreserve that can test links and retrieve Internet
Archive replacements. The tool can output the result of individual links, or
take a CSV list to output collected information in JSON, BoltDB, or CSV format.

## Usage
```bash
Usage:  linkstat [Optional -link] [Optional -label]
                 [Optional -list] [Optional -json]
                                  [Optional -bolt]
                                  [Optional -csv]
                 [Optional -version -v]

Output: [Json]
Output: [CSV]
Output: [BoltDB]
Output: [Version] 'exponentialDK-httpreserve/0.0.9 ...'

Usage of ./linkstat:
  -bolt
    	Output to static BoltDB.
  -csv
    	Output to CSV.
  -json
    	Output to JSON.
  -label string
    	Annotate single URL check response with label.
  -link string
    	Seek the status of a single URL: JSON
  -list string
    	Provide a list of URLs to test against in CSV format.
  -v	Return httpreserve version.
  -version
    	Return httpreserve version.
```

## Examples

#### Example combining [tikalinkextract][httpreserve-1]

Inspired by Harvard Innovation Labs to test the ability of
httpreserve-workbench at the time. This CLI version is a simplification of that
work but should still produce decent results. HTTPreserve
[Million Dollar Webpage Project][httpreserve-2]

[httpreserve-1]: https://github.com/httpreserve/tikalinkextract
[httpreserve-2]: https://github.com/httpreserve/million-dollar-webpage

#### CSV input

An input CSV `example.csv` might look as follows:
```csv
"BBC News", "http://www.bbc.co.uk/news"
"BBC Home", "http://www.bbc.co.uk/"
"BBC Radio", "http://www.bbc.co.uk/radio"
"Google", "http://www.google.com"
"exponentialdecay.co.uk", "http://www.exponentialdecay.co.uk"
"Internet Archive", "http://www.archive.org"
"perma.cc", "http://perma.cc"
"wikipedia.org", "http://wikipedia.org"
"The Million Dollar Homepage", "http://www.getpixel.net"
```

To output a CSV collecting all of the linkstat results, you can run a command
as follows:
```bash
$ ./linkstat -csv --list example.csv > output.csv
```

And the output looks as follows:
```
"id","filename","link","response code","response text","title","content-type","archived","internet archive response code","internet archive response text","wayback earliest date","internet archive earliest","wayback latest date","internet archive latest","internet archive save link","protocol error","protocol error","analysis version number","analysis version text","stats creation time"
"1651a00b16a12ba06fc6c6b049c7cf7c","BBC News","https://www.bbc.co.uk/news","200","OK","home - bbc news","text/html;charset=utf-8","true","302","Found","09 October 1997","http://web.archive.org/web/19971009011901/http://www.bbc.co.uk/news/","19 March 2019","http://web.archive.org/web/20190319173721/https://www.bbc.co.uk/news","http://web.archive.org/save/https://www.bbc.co.uk/news","","","0.0.9","exponentialDK-httpreserve/0.0.9","1.574649021s"
"57ab6349a47b53b982a939fb1da54fef","BBC Radio","https://www.bbc.co.uk/sounds","200","OK","bbc sounds - music. radio. podcasts","text/html; charset=utf-8","true","302","Found","19 March 2008","http://web.archive.org/web/20080319074038/http://www.bbc.co.uk/sounds","18 March 2019","http://web.archive.org/web/20190318211158/https://www.bbc.co.uk/sounds","http://web.archive.org/save/https://www.bbc.co.uk/sounds","","","0.0.9","exponentialDK-httpreserve/0.0.9","1.660729358s"
"c85da5e372ffe2200e46527b74537ba3","BBC Home","https://www.bbc.co.uk/","200","OK","bbc - home","text/html; charset=utf-8","true","302","Found","21 December 1996","http://web.archive.org/web/19961221203254/http://www0.bbc.co.uk/","19 March 2019","http://web.archive.org/web/20190319141018/https://www.bbc.co.uk/","http://web.archive.org/save/https://www.bbc.co.uk/","","","0.0.9","exponentialDK-httpreserve/0.0.9","1.95442772s"
"b3bd672c1014e07e87ef4a357a161528","exponentialdecay.co.uk","http://www.exponentialdecay.co.uk","206","Partial Content","ross spencer, digital preservation, archives, python developer, golang developer, uk, nz","text/html","true","302","Found","17 September 2008","http://web.archive.org/web/20080917054811/http://www.exponentialdecay.co.uk/","13 November 2018","http://web.archive.org/web/20181113021338/http://exponentialdecay.co.uk/","http://web.archive.org/save/http://www.exponentialdecay.co.uk","","","0.0.9","exponentialDK-httpreserve/0.0.9","425.368183ms"
```

#### An individual link

The command: `./linkstat -link https://github.com/ -label "GitHub"` will
output:
```json
{
   "FileName": "GitHub",
   "AnalysisVersionNumber": "0.0.15",
   "AnalysisVersionText": "exponentialDK-httpreserve/0.0.15",
   "SimpleRequestVersion": "httpreserve-simplerequest/0.0.4",
   "Link": "https://github.com/",
   "Title": "github: let’s build from here · github",
   "ContentType": "text/html; charset=utf-8",
   "ResponseCode": 200,
   "ResponseText": "OK",
   "SourceURL": "https://github.com/",
   "ScreenShot": "snapshots are not currently enabled",
   "InternetArchiveLinkEarliest": "http://web.archive.org/web/20080514210148/http://github.com/",
   "InternetArchiveEarliestDate": "2008-05-14 21:01:48 +0000 UTC",
   "InternetArchiveLinkLatest": "http://web.archive.org/web/20230829062855/https://github.com/",
   "InternetArchiveLatestDate": "2023-08-29 06:28:55 +0000 UTC",
   "InternetArchiveSaveLink": "http://web.archive.org/save/https://github.com/",
   "InternetArchiveResponseCode": 302,
   "InternetArchiveResponseText": "Found",
   "RobustLinkEarliest": "<a href='http://web.archive.org/web/20080514210148/http://github.com/' data-originalurl='https://github.com/' data-versiondate='2008-05-14'>HTTPreserve Robust Link - simply replace this text!!</a>",
   "RobustLinkLatest": "<a href='http://web.archive.org/web/20230829062855/https://github.com/' data-originalurl='https://github.com/' data-versiondate='2023-08-29'>HTTPreserve Robust Link - simply replace this text!!</a>",
   "PWID": "urn:pwid:archive.org:2023-08-29T06:28:55Z:page:https://github.com/",
   "Archived": true,
   "Error": false,
   "ErrorMessage": "",
   "StatsCreationTime": "7.070152149s"
}
```

## Archiving Weblinks

* [Find and Connect Project:][linkstat-1] Nicola Laurent on the impact of
broken links.
* [Binary Trees? Automatically Identifying the links between born digital records:][linkstat-2]
I write about hyperlinks as a public record in own right when submitted as part
of a documentary heritage.
* [HiberActive Pilot][linkstat-3] A scholarly publishing tool that extracts
URLs, returns both the original URL and a perma-link.
* [IIPC Awesome List][linkstat-4] A list of web-archiving links that invites
contributions from the community to keep it up-to-date.

[linkstat-1]: http://www.findandconnectwrblog.info/2016/11/broken-links-broken-trust/
[linkstat-2]: https://www.youtube.com/watch?v=Ked9GRmKlRw
[linkstat-3]: https://www.era.lib.ed.ac.uk/handle/1842/23366
[linkstat-4]: https://github.com/iipc/awesome-web-archiving

## License

GNU General Public License Version 3. [Full Text](LICENSE)
