Bugfix: Some mods are missing a download URL

We found the cases where the provided download URL from the curseforge API were
missing, to work around this we are constructing the URL as a fallback now. This
is identicated by the fallback flag within the logs now.

https://github.com/webhippie/cursecli/issues/8
