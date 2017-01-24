package recordutil

import "golang.org/x/text/transform"

// NewlineNormalizer is a transform.Transformer to transform CR/LF to single LF.
type NewlineNormalizer struct {
	prev byte
}

// Transform is a func of transform.Transformer.
func (nn *NewlineNormalizer) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	for nSrc < len(src) && nDst < len(dst) {
		b := src[nSrc]
		nSrc++
		switch b {
		case '\r':
			dst[nDst] = '\n'
		case '\n':
			if nn.prev == '\r' {
				nn.prev = b
				continue
			}
			dst[nDst] = b
		default:
			dst[nDst] = b
		}
		nDst++
		nn.prev = b
	}
	if nSrc < len(src) {
		err = transform.ErrShortDst
	}
	return
}

// Reset is a func of transform.Transformer.
func (nn *NewlineNormalizer) Reset() {
	nn.prev = 0
}

// NewNewlineNormalizer returns a new transform.Transformer of NewlineNormalizer.
func NewNewlineNormalizer() transform.Transformer {
	return &NewlineNormalizer{}
}
