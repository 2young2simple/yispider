package common

import (
	"fmt"

	"strings"

	"github.com/saintfish/chardet"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var (
	charsetDetector  = chardet.NewTextDetector()
	charsetDetectors = map[string]encoding.Encoding{
		"Big5":         traditionalchinese.Big5,
		"EUC-JP":       japanese.EUCJP,
		"EUC-KR":       korean.EUCKR,
		"GB-18030":     simplifiedchinese.GB18030,
		"ISO-2022-JP":  japanese.ISO2022JP,
		"ISO-8859-5":   charmap.ISO8859_5,
		"ISO-8859-6":   charmap.ISO8859_6,
		"ISO-8859-7":   charmap.ISO8859_7,
		"ISO-8859-8":   charmap.ISO8859_8,
		"ISO-8859-8-I": charmap.ISO8859_8I,
		"KOI8-R":       charmap.KOI8R,
		"Shift_JIS":    japanese.ShiftJIS,
		"UTF-16BE":     unicode.UTF16(unicode.BigEndian, unicode.UseBOM),
		"UTF-16LE":     unicode.UTF16(unicode.LittleEndian, unicode.UseBOM),
		"windows-1251": charmap.Windows1251,
		"windows-1252": charmap.Windows1252,
		"windows-1253": charmap.Windows1253,
		"windows-1254": charmap.Windows1254,
		"windows-1255": charmap.Windows1255,
		"windows-1256": charmap.Windows1256,
	}
)

func ToUtf8(html []byte) ([]byte, error) {
	r, err := charsetDetector.DetectBest(html)
	if err != nil {
		return nil, err
	}
	if strings.ToLower(r.Charset) == strings.ToLower("UTF-8") || strings.ToLower(r.Charset) == strings.ToLower("ISO-8859-1") || strings.ToLower(r.Charset) == strings.ToLower("Big5") {
		return html, nil
	}

	t, ok := charsetDetectors[r.Charset]
	if !ok {
		return nil, fmt.Errorf(
			"could not find charset decoder for `%s`",
			r.Charset)
	}

	html, _, err = transform.Bytes(t.NewDecoder(), html)
	return html, err
}
