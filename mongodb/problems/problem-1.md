Problem: `querySrv ESERVFAIL _mongodb._tcp.cluster0.abcd0.mongodb.net`

Causes: `querySrv ESERVFAIL` is a DNS error. This means that your local machine is not able to get a response from your DNS resolver for the SRV record `_mongodb._tcp.cluster0.abcd0.mongodb.net`.

From your local machine, test SRV lookup from a command line, possibly one of these:

```bash
nslookup -type=SRV _mongodb._tcp.cluster0.abcd0.mongodb.net
host -t SRV _mongodb._tcp.cluster0.abcd0.mongodb.net
```

If that fails, your DNS provider is the problem.

Solution: Go to the Atlas UI and get the pre-3.6 connection string. It will start with mongodb:// and not mongodb+srv://.

Reference: https://stackoverflow.com/questions/68875026/error-querysrv-eservfail-mongodb-tcp-cluster0-abcd0-mongodb-net
