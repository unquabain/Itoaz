// Copyright 2022 Ben C. Forsberg.
// See LICENSE.txt for license information.

package itoaz

import (
  "testing"
  "fmt"
)

type Example struct {
	Number   int
	Expected string
	Format
}

func (e Example) Test(t *testing.T) {
	actual := e.Itoa(e.Number)
	if actual != e.Expected {
		t.Fatalf(`expected %d to render as %q, got %q`, e.Number, e.Expected, actual)
	}
}

func TestItoaz(t *testing.T) {
	normalBase10 := Format{
		Radix:    10,
		Alphabet: []rune(`0123456789`),
	}

	nozeroBase10 := Format{
		Radix:    10,
		Alphabet: []rune(`123456789X`),
		NoZero:   true,
	}
	excelFormat := Format{
		Radix:    26,
		Alphabet: []rune(`ABCDEFGHIJKLMNOPQRSTUVWXYZ`),
		NoZero:   true,
	}
	examples := []Example{
		{
			Number:   0,
			Expected: `0`,
			Format:   normalBase10,
		},

		{
			Number:   1,
			Expected: `1`,
			Format:   normalBase10,
		},

		{
			Number:   9,
			Expected: `9`,
			Format:   normalBase10,
		},

		{
			Number:   10,
			Expected: `10`,
			Format:   normalBase10,
		},

		{
			Number:   11,
			Expected: `11`,
			Format:   normalBase10,
		},

		{
			Number:   99,
			Expected: `99`,
			Format:   normalBase10,
		},

		{
			Number:   100,
			Expected: `100`,
			Format:   normalBase10,
		},

		{
			Number:   101,
			Expected: `101`,
			Format:   normalBase10,
		},

		{
			Number:   1000,
			Expected: `1000`,
			Format:   normalBase10,
		},

		{
			Number:   1001,
			Expected: `1001`,
			Format:   normalBase10,
		},

		{
			Number:   1099,
			Expected: `1099`,
			Format:   normalBase10,
		},

		{
			Number:   1100,
			Expected: `1100`,
			Format:   normalBase10,
		},

		{
			Number:   1101,
			Expected: `1101`,
			Format:   normalBase10,
		},
		// NON-ZERO

		{
			Number:   0,
			Expected: ``,
			Format:   nozeroBase10,
		},

		{
			Number:   1,
			Expected: `1`,
			Format:   nozeroBase10,
		},

		{
			Number:   9,
			Expected: `9`,
			Format:   nozeroBase10,
		},

		{
			Number:   10,
			Expected: `X`,
			Format:   nozeroBase10,
		},

		{
			Number:   11,
			Expected: `11`,
			Format:   nozeroBase10,
		},

		{
			Number:   99,
			Expected: `99`,
			Format:   nozeroBase10,
		},

		{
			Number:   100,
			Expected: `9X`,
			Format:   nozeroBase10,
		},

		{
			Number:   101,
			Expected: `X1`,
			Format:   nozeroBase10,
		},

		{
			Number:   1000,
			Expected: `99X`,
			Format:   nozeroBase10,
		},

		{
			Number:   1001,
			Expected: `9X1`,
			Format:   nozeroBase10,
		},

		{
			Number:   1099,
			Expected: `X99`,
			Format:   nozeroBase10,
		},

		{
			Number:   1100,
			Expected: `X9X`,
			Format:   nozeroBase10,
		},

		{
			Number:   1101,
			Expected: `XX1`,
			Format:   nozeroBase10,
		},

		// Excel Format
		{
			Number:   1,
			Expected: `A`,
			Format:   excelFormat,
		},

		{
			Number:   2,
			Expected: `B`,
			Format:   excelFormat,
		},

		{
			Number:   26,
			Expected: `Z`,
			Format:   excelFormat,
		},

		{
			Number:   27,
			Expected: `AA`,
			Format:   excelFormat,
		},

		{
			Number:   26 * 26,
			Expected: `YZ`,
			Format:   excelFormat,
		},

		{
			Number:   26*26 + 26,
			Expected: `ZZ`,
			Format:   excelFormat,
		},

		{
			Number:   26*26 + 27,
			Expected: `AAA`,
			Format:   excelFormat,
		},

		{
			Number:   26*26 + 26*2,
			Expected: `AAZ`,
			Format:   excelFormat,
		},

		{
			Number:   26*26*2 + 26*4 + 13,
			Expected: `BDM`,
			Format:   excelFormat,
		},
	}
	for _, example := range examples {
		example.Test(t)
	}

}

func ExampleItoaz() {
  // Print numbers in Base10, but don't use zero.
  // The digit 'X' will stand for 10. The result
  // should be as if you took a normal, base10
  // number, but for every '0', you replaced it
  // with an 'X' and borrowed one from the column
  // to the left. 302 becomes 2X2.
  alphabet := []rune(`123456789X`)
  radix := 10
  nozero := true

  fmt.Println(Itoaz(1,   radix, alphabet, nozero)) // 1
  fmt.Println(Itoaz(10,  radix, alphabet, nozero)) // X
  fmt.Println(Itoaz(11,  radix, alphabet, nozero)) // 11
  fmt.Println(Itoaz(20,  radix, alphabet, nozero)) // 1X
  fmt.Println(Itoaz(101, radix, alphabet, nozero)) // X1
  fmt.Println(Itoaz(110, radix, alphabet, nozero)) // XX
  fmt.Println(Itoaz(111, radix, alphabet, nozero)) // 111
  fmt.Println(Itoaz(302, radix, alphabet, nozero)) // 2X2
  fmt.Println(Itoaz(345, radix, alphabet, nozero)) // 345
  // Output:
  // 1
  // X
  // 11
  // 1X
  // X1
  // XX
  // 111
  // 2X2
  // 345
}

func ExampleFormat_Itoa() {
  // This shows the raison d'etre for this module: spreadsheet column
  // formatting.

  fmt.Println(Column.Itoa(1))  // A
  fmt.Println(Column.Itoa(26)) // Z
  fmt.Println(Column.Itoa(27)) // AA

  // Output:
  // A
  // Z
  // AA
}
