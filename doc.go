/*
Package log is a generic logging interface.
Logger is a leveled, structured, and nest-able logging interface. It is
intended for embedding (and not importing) into your projects.

Philosophy

Logging interfaces are too complex as laid out in 
https://dave.cheney.net/2015/11/05/lets-talk-about-logging. To that end
this package opts for a very simple interface that only supports Info
and Debug levels as well as an ability to include logger context for use
in sub-loggers by using With. Three methods, that's it. To facilitate
structured logging we accept variadic parameters of any type with the
understanding (but not requirement) that the receiver of those
parameters will likely be implementing a key-value logging style with
string keys.

The Library Dependency Problem

Just moving to simpler logger interfaces doesn't solve everything.
Dependency hell and opinionated logging libraries can make third party
library use annoying.

Consider a Go project, call it Foo. Foo has a package called Bar that
you would like to use on its own in your project. All of Foo and thus
Bar, too, use a third party logging project called AllTheLogs. But your
project already has a different logging project. So in order to use
Bar unmodified and log to your existing logger you would either have to
use AllTheLogs logger or adapt its interfaces to plug into your existing
logger. Not ideal in either case.

While it's true that we're just inventing another abstraction here the
thought is that the super simple interface makes this the process of
adapting log frameworks much much simpler.

Embed, Not Import

Unlike other interface-only logging projects like
https://github.com/go-log/log or https://github.com/go-logr/logr it is
intended that you embed this package (or the parts that you need) into
your project instead of depending on (importing) it. Though you could
import it if you like, the idea is to get away from the above Library
Dependency Problem. In this way a third party user of your package does
not need to commit to your logging project of choice.

By embedding this logging interface we can come close to achieving our
main goals here: a tiny limited scope (but featureful) interface, and 
no third party dependencies for our shareable packages.

Actually Logging

As this package only includes an interface (and a no-op logger) how does
one actually log logs? You don't. Not with this package. A logger
implementation that fits the simple Logger interface needs to be 
created. Far more likely though is an implementation of another logger
project is adapted to this interface.
*/
package log
