<h1 align="center">TimeCalc</h1>
<p align="center">A command-line calculator with support for performing operations on time</p>
<p align="center">Written exclusively in <a href="https://golang.org">Go</a></p>

- [Overview](#overview)
- [Installation](#installation)
	- [From Source](#from-source)
- [Usage](#usage)

## <span id="overview">Overview</span>

Sometimes just being able to add numbers isn't enough. Often, you need to perform some calculation that involves a period of time. Doing so with a traditional calculator requires constantly transposing the time to and from a natural number that the calculator can handle. This can become extremely time consuming and even unreliable.

TimeCalc aims to solve this problem by providing a complete calculator interface exactly as you would expect, capable of addition, subtraction, multiplication, division, and even remainder operations, but with support for time. Time is formatted exactly as you might expect, right down to the millisecond – `hour:minute:second.millisecond` – and can be mixed in anywhere with any other calculations. Time takes precedence over a general number, ensuring that whenever a time is involved in an equation the output is always in time format.

## <span id="installation">Installation</span>

Download the latest [release](https://github.com/octacian/timecalc/releases/latest) for your operating system (be sure to download a binary file, not the source code). Open a command line where you saved the binary and run:

```
$ ./timecalc
```

TimeCalc can be installed globally by moving its executable to `/usr/bin/` and locally by adding the executable to your `PATH`:

```
$ export PATH=$PATH:/path/to/timecalc
```

After re-opening the command line you can now simply run:

```
$ timecalc
```

### <span id="from-source">From Source</span>

Make sure you have a working Go environment. Go version 1.10 has been tested. [See the install instructions for Go.](https://golang.org/doc/install)

To install TimeCalc to your `GOPATH`, simply run:

```
$ go get github.com/octacian/timecalc
```

Make sure your `PATH` includes the `$GOPATH/bin` directory so your commands can be easily used:

```
export PATH=$PATH:$GOPATH/bin
```

## <span id="usage">Usage</span>

TimeCalc utilizes a natural calculator syntax supporting negative and positive numbers and decimals, groups, and of course, time periods (__warning__: order of operations outside of basic groups is not yet complete). For example:

```
>>> 10 + (((7 * 81) - 42) / 20)
36.25
>> 10 + ((7 * 81) - 42) / 20
26.75 # Demonstrates lack of order of operations
>>> .5 - 1.7
-1.2
>>> ::30.5 * 3
00:01:31.5
>>> :30 % (8 * 1000 * 60)
00:06
>>> : + 1
00:00:00.001 # Ones place represents milliseconds, see below
```

Several shorthand structures are demonstrated above:
- `.5` –> `0.5`
- `:` –> `00:00:00.0`
- `1:` –> `1:00:00.0`
- `:10` –> `00:10:00.0`
- `::05` –> `00:00:05.0`
- `::.8` –> `00:00:00.8` (800 milliseconds)
- `:10:.2` –> `00:10:00.2` (10 minutes and 200 milliseconds)
- etc...

Times are converted to and from numbers as a simple method of applying operations by way of transposing all fields to milliseconds. The max precision of a time is milliseconds, and as a result, when a number is converted to a time, the __ones place represents a single millisecond__.
