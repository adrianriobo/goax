# goax

Small framework to interact with UX element through OS accessibility native APIs in golang

## Security Threat

goax.exe can be detected as a threat by windows defender, as it may be identified as a [Trojan:Win32/Wacatac.B!ml](https://www.makeuseof.com/windows-wacatac-trojan/)  
this is a **false positive** from the code you can see that goax uses OS dll libraries to access OS functionality (accessibility, messaging,...).  
