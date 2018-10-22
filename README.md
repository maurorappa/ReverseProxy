# ReverseProxy

 
This is simple HTTP reverse proxy with some advanced features tipical of a WAF (web application firewall) like rate limiting, URL/Method firewall and urlknocking. It also logs all HTTP requests for trafic analysis.
The most interesting features are:

* Urlknocking (what?) 
 
It is an implementation in HTTP of a concept called PortKnocking (https://en.wikipedia.org/wiki/Port_knocking).  
This software is desigend to be placed in front of a blog engine to log all accesses and secure the management part of it.
When the proxy starts it blocks by default the specified path (ie/admin) but it allows access if a specific url is requested,
its string is randomly generated at the startup,shown in the standard output and conveniently sent to your mailbox!
in the config file you can specify the lenght of this random string to make it less likely to be guessed,
the charset used includes letters, numbers and some special caracters.
The IP of this request is saved and every request to the restriced area will be restricted just to this IP.
After having finished the management work, you can block again the access to this path by calling another random generated URL.
The next implementation will make them valid for one time and create a new one

This is the output during startup:

```

Connecting to http://172.18.0.2:2368/...
Open sesame path will be: gucx_qdbmyw26uuz!
Close sesame path will be: cpsuxfmcnoq1zz:un
Listening to 0.0.0.0:80

```

* Block undesired file extensions

My blog does not contains a single page written in PHP, but all web worms scan for Wordpress resources which happens to be php pages. With some clever filtering, this proxy block those requests and redirect somewhere else, this is to avoid to clutter the access log and make the backend waste resources.

* Rate limiting

The HTTP handler implements a rate limiter (token bucket type) with burst capacity. The config file contains the values for a personal blog the suggested 20 concurrent connections is a good value considering an HTTP request in Go get executed in microseconds. The burst limit is hardcoded to be three times the limit.
More information about its implementation can be found here: https://godoc.org/golang.org/x/time/rate

* Method filter

You can also blacklist the method not implemented in your reverseproxy to protect it from maliciuos activities, for example block all calls using method DELETE or PUT .

