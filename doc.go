/*
enables formatting of numbers, scaled and suffixed, for more compact and readable numerical i/o.

this is done by wrapping a number in a generic type implementing fmt.Scanner+fmt.Stringer, then using that inplace of the number for i/o.

scaling is achieved though manipulation of the string returned by fmt.Print, on the embedded value, not by calcuation.

[SI Suffix]: https://en.wikipedia.org/wiki/Metric_prefix.
*/
package scalingSuffix

