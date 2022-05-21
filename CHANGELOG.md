# Changelog for 1.1.1

The following sections list the changes for 1.1.1.

## Summary

 * Fix #8: Some mods are missing a download URL

## Details

 * Bugfix #8: Some mods are missing a download URL

   We found the cases where the provided download URL from the curseforge API were missing, to work
   around this we are constructing the URL as a fallback now. This is identicated by the fallback
   flag within the logs now.

   https://github.com/webhippie/cursecli/issues/8


# Changelog for 1.1.0

The following sections list the changes for 1.1.0.

## Summary

 * Fix #6: Switch API as it is secured now

## Details

 * Bugfix #6: Switch API as it is secured now

   We switched to the official Curseforge API as the previously used API is not usable anymore.
   Part of that is also introducing a flag for the API key as we don't have any anonymous API
   available anymore.

   https://github.com/webhippie/cursecli/issues/6


# Changelog for 1.0.0

The following sections list the changes for 1.0.0.

## Summary

 * Chg #3: Initial release of basic version

## Details

 * Change #3: Initial release of basic version

   Just prepared an initial basic version which could be released to the public.

   https://github.com/webhippie/cursecli/issues/3


