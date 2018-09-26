# ReverseProxy

 
This is simple HTTP reverse proxy with some advanced features tipicla of a WAF (web application firewall) like rate limiting, URL/Method firewall and urlknocking
(a secret URL which allows access to a specific path, like the admin zone).
It also logs all HTTP requests for trafic analysis.
 
This software si desigend to be placed in front of a blog engine to log all accesses and secure the management part of it.
When the proxy starts it blocks by default the specified path (ie/admin) but it allows access if a specific url is requested,
its string is randomly generated at the startup and shown in the standard output.
in the config file you can specify the lenght of this random string to make it less likely to be guessed,
the charset used includes letters, numbers and some special caracters.
After having finished the management work, you can block again the access to this path by calling another random generated URL.

This is the output during startup

```

Connecting to http://172.18.0.2:2368/...
Open sesame path will be: gucx_qdbmyw26uuz!
Close sesame path will be: cpsuxfmcnoq1zz:un
Listening to 0.0.0.0:80

```
