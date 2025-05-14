// scalingSuffix enables printing of scaled numbers and appropriate suffixes.
// this is done by wrapping a number in a fmt.Formatter where the scaling can be chosen directly using the formatters verb or, more generally, automatically determined for the value of the number.
// the scaling verb used is the same as the suffix and is an SI standard suffix.
// types:
// SI :- scaling in powers of 1000.
// IEC :- scaling in powers of 1024.
// flags:
// '+'; the number, after scaling, is separated (in blocks of three digits) using a ','
// '-'; only the suffix is returned. (useful when scaling is automatic)
//
// other types here, implementing fmt.Stringer, just perform comma separation without scaling.
// Uint :- unsigned
// Int :- signed 
// Float :- float (no separation below decimal point)

package scalingSuffix

