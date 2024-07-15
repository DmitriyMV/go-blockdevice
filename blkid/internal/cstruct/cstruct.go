// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

//go:build exclude

package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	uintRe      = regexp.MustCompile(`(uint16_t|uint32_t|uint64_t|uint8_t|unsigned char)\s+(\w+);(?:\s+/\*\s+(.+?)\s+\*/)?`)
	charRe      = regexp.MustCompile(`(char|unsigned char|uint8_t)\s+(\w+)\[(\d+)\];(?:\s+/\*\s+(.+?)\s+\*/)?`)
	uintSliceRe = regexp.MustCompile(`(uint32_t|uint16_t)\s+(\w+)\[(\d+)\];(?:\s+/\*\s+(.+?)\s+\*/)?`)
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("cstruct: ")

	var (
		pkgName    string
		structName string
		input      string
		endianness string
	)

	flag.StringVar(&pkgName, "pkg", "", "package name")
	flag.StringVar(&structName, "struct", "", "struct name")
	flag.StringVar(&input, "input", "", "input file")
	flag.StringVar(&endianness, "endianness", "NativeEndian", "field endianness")
	flag.Parse()

	var g Generator

	// Print the header and package clause.

	g.Printf("// This Source Code Form is subject to the terms of the Mozilla Public\n")
	g.Printf("// License, v. 2.0. If a copy of the MPL was not distributed with this\n")
	g.Printf("// file, You can obtain one at http://mozilla.org/MPL/2.0/.\n\n")

	g.Printf("// Code generated by \"cstruct %s\"; DO NOT EDIT.\n", strings.Join(os.Args[1:], " "))
	g.Printf("\n")
	g.Printf("package %s\n", pkgName)
	g.Printf("import \"encoding/binary\"\n\n")
	g.Printf("var _ = binary.%s\n\n", endianness)

	in, err := os.Open(input)
	if err != nil {
		log.Fatalf("opening input file: %s", err)
	}

	defer in.Close()

	scanner := bufio.NewScanner(in)

	sizeConst := strings.ToUpper(structName) + "_SIZE"

	g.Printf("// %s is a byte slice representing the %s C header.\n", structName, input)
	g.Printf("type %s []byte;\n", structName)

	var size int

	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case uintRe.MatchString(line):
			matches := uintRe.FindStringSubmatch(line)

			typ, name, comment := matches[1], matches[2], matches[3]

			if comment == "" {
				comment = name
			}

			switch typ {
			case "uint64_t":
				g.Printf("// Get_%s returns %s.\n", name, comment)
				g.Printf("func (s %s) Get_%s() uint64 {\nreturn binary.%s.Uint64(s[%d:%d])\n}\n", structName, name, endianness, size, size+8)
				g.Printf("// Put_%s sets %s.\n", name, comment)
				g.Printf("func (s %s) Put_%s(v uint64) {\nbinary.%s.PutUint64(s[%d:%d], v)\n}\n", structName, name, endianness, size, size+8)

				size += 8
			case "uint32_t":
				g.Printf("// Get_%s returns %s.\n", name, comment)
				g.Printf("func (s %s) Get_%s() uint32 {\nreturn binary.%s.Uint32(s[%d:%d])\n}\n", structName, name, endianness, size, size+4)
				g.Printf("// Put_%s sets %s.\n", name, comment)
				g.Printf("func (s %s) Put_%s(v uint32) {\nbinary.%s.PutUint32(s[%d:%d], v)\n}\n", structName, name, endianness, size, size+4)

				size += 4
			case "uint16_t":
				g.Printf("// Get_%s returns %s.\n", name, comment)
				g.Printf("func (s %s) Get_%s() uint16 {\nreturn binary.%s.Uint16(s[%d:%d])\n}\n", structName, name, endianness, size, size+2)
				g.Printf("// Put_%s sets %s.\n", name, comment)
				g.Printf("func (s %s) Put_%s(v uint16) {\nbinary.%s.PutUint16(s[%d:%d], v)\n}\n", structName, name, endianness, size, size+2)

				size += 2
			case "uint8_t", "unsigned char":
				g.Printf("// Get_%s returns %s.\n", name, comment)
				g.Printf("func (s %s) Get_%s() byte {\nreturn s[%d]\n}\n", structName, name, size)
				g.Printf("// Put_%s sets %s.\n", name, comment)
				g.Printf("func (s %s) Put_%s(v byte) {\ns[%d] = v\n}\n", structName, name, size)

				size++
			}
		case charRe.MatchString(line):
			matches := charRe.FindStringSubmatch(line)

			name, fieldSizeStr, comment := matches[2], matches[3], matches[4]

			if comment == "" {
				comment = name
			}

			fieldSize, err := strconv.Atoi(fieldSizeStr)
			if err != nil {
				log.Fatalf("parsing field size: %s", err)
			}

			g.Printf("// Get_%s returns %s.\n", name, comment)
			g.Printf("func (s %s) Get_%s() []byte {\nreturn s[%d:%d]\n}\n", structName, name, size, size+fieldSize)
			g.Printf("// Put_%s sets %s.\n", name, comment)
			g.Printf("func (s %s) Put_%s(v []byte) {\ncopy(s[%d:%d], v)\n}\n", structName, name, size, size+fieldSize)

			size += fieldSize
		case uintSliceRe.MatchString(line):
			matches := uintSliceRe.FindStringSubmatch(line)

			typ, fieldSizeStr := matches[1], matches[3]

			fieldSize, err := strconv.Atoi(fieldSizeStr)
			if err != nil {
				log.Fatalf("parsing field size: %s", err)
			}

			switch typ {
			case "uint32_t":
				size += fieldSize * 4
			case "uint16_t":
				size += fieldSize * 4
			}
		default:
			log.Printf("skipping line: %s", line)
		}
	}

	g.Printf("// %s is the size of the %s struct.\n", sizeConst, structName)
	g.Printf("const %s = %d\n", sizeConst, size)

	src := g.format()

	if err := os.WriteFile(strings.ToLower(structName)+".go", src, 0o644); err != nil {
		log.Fatalf("writing output: %s", err)
	}
}

type Generator struct {
	buf bytes.Buffer
}

func (g *Generator) Printf(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format, args...)
}

// format returns the gofmt-ed contents of the Generator's buffer.
func (g *Generator) format() []byte {
	src, err := format.Source(g.buf.Bytes())
	if err != nil {
		// Should never happen, but can arise when developing this code.
		// The user can compile the output to see the error.
		log.Printf("warning: internal error: invalid Go generated: %s", err)
		log.Printf("warning: compile the package to analyze the error")
		return g.buf.Bytes()
	}

	return src
}
