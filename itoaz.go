// (c) 2022 Ben C. Forsberg.
// See LICENSE.txt for license information.

// Module itoaz implements a completely ordinary integer-to-string
// conversion, but adds the option to not use zeroes. An example of
// when this might be useful is converting an integer to the name of
// a column in a spreadsheet. The number after 'Z' is 'AA', but how
// can that be possible? The 'A' in the left column stands for 1 (x26),
// but the 'A' in the right stands for 0!
//
// I don't imagine this is a much-needed library. But the problem was
// bugging me, and I wanted to work out a solution for myself.
package itoaz

// reverse converts the digits from left-to-right, or "English" order
// to right-to-left, "Arabic" order as we're used to seeing numbers
// represented.
func reverse(orig []rune) []rune {
	l := len(orig)
	last := l - 1
	mid := l / 2
	for i := 0; i < mid; i++ {
		orig[i], orig[last-i] = orig[last-i], orig[i]
	}
	return orig
}

// Itoaz converts a number to a string using the supplied radix and alphabet.
// The nozero option controls how the alphabet is interpreted. If nozero is false,
// the runes in the alphabet are used for the digits from 0 to radix - 1. If nozero
// is true, the values represent 1 to radix.
// 
// This function will panic if the length of alphabet is not equal to radix.
func Itoaz(num, radix int, alphabet []rune, nozero bool) (az string) {
	if len(alphabet) != radix {
		panic(`alphabet size must match radix`)
	}
	digits := make([]rune, 0)
	var rem int
	var nozeroOffset int
	if nozero {
		nozeroOffset = -1
	}
	defer func() {
		if !nozero && len(digits) == 0 {
			az = `0`
		} else {
			az = string(reverse(digits))
		}
	}()
	for {
		if num == 0 {
			return
		}
		num, rem = num/radix, num%radix
		if nozero && rem == 0 {
			rem = radix
			num--
		}
		digits = append(digits, alphabet[rem+nozeroOffset])
	}
}

// Format is a struct that allows you to reuse the extra arguments to Itoaz
// for multiple numbers.
type Format struct {
	Alphabet []rune
	Radix    int
	NoZero   bool
}

// Itoa uses the parameters of Format to convert an integer into a string.
func (f Format) Itoa(num int) string {
	return Itoaz(num, f.Radix, f.Alphabet, f.NoZero)
}

// Base10 is your normal, run of the mill, Arabic number representation.
var Base10 = Format{
  Alphabet: []rune(`0123456789`),
  Radix: 10,
  NoZero: false,
}

// Hexadecimal is your normal, run-of-the-mill base16 number representation. 
var Hexadecimal = Format{
  Alphabet: []rune(`0123456789ABCDEF`),
  Radix: 16,
  NoZero: false,
}

// Column should be how your favorite spreadsheet software names its columns
// This is really the whole reason for this library.
var Column = Format{
  Alphabet: []rune(`ABCDEFGHIJKLMNOPQRSTUVWXYZ`),
  Radix: 26,
  NoZero: true,
}
