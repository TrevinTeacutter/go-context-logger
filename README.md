# go-context-logger

## Goals

### Log levels have one switch, normal or verbose

While [Dave Cheney's article about logging](https://dave.cheney.net/2015/11/05/lets-talk-about-logging) inspired a lot of logging APIs to try and simplify, they often left V levels in. IMO V levels are actually worse than typical log levels because they are even more arbitrary and may differ worse across package/API boundaries. They have no consistent meaning that communities try to uphold and the user/operator experience is not made better. The typical behavior is if you are looking for some log that may tell you what is going on, you bump V levels up a few times and if you don't find what you are looking for you set it as high as possible. Noisy logs are always a problem but what if you had a way to tune out noisy logs...

### Code specific controls

To better control and tune out logs, you could build something into your log collection pipeline to drop them, but this could lead to a lot of complication, especially in environments where you may own your code but not the log collection pipeline. So what do you do? You could use log levels like discussed above with all the pain points, or you could add configuration to change log levels of packages, files, functions, etc. Stuff that should already be in the log message to help inform you of issues.

Since this is still meant to be a structured logger first, you can use matchers against the log itself to filter out logs even leveraging the message if need be.

### Append functionality with labels

### Support for different outputs

### Lazy evaluation
