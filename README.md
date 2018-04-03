# URLIZER

simple tool to drop a list of things into a url.

ex:

```shell
urlizer "a,b,c,d,e,f,j" "https://super.cool-site.com/things/{}"
```

will result in 
```https://super.cool-site.com/things/a  
https://super.cool-site.com/things/b  
https://super.cool-site.com/things/c  
https://super.cool-site.com/things/d  
https://super.cool-site.com/things/e  
https://super.cool-site.com/things/f  
https://super.cool-site.com/things/j  
```