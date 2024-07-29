package signatures

type Signature []byte

var (
	JPEG  = Signature{0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 0x4A, 0x46}
	JPEG2 = Signature{0xFF, 0xD8, 0xFF, 0xE1}
	PNG   = Signature{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
	RAF   = Signature{0x46, 0x55, 0x4A, 0x49, 0x46, 0x49, 0x4C, 0x4D}
)

var acceptedSignatures = []Signature{JPEG, JPEG2, PNG, RAF}

func IsSignature(data []byte, signature Signature) bool {
	for i, b := range signature {
		if data[i] != b {
			return false
		}
	}

	return true
}

func IsAcceptedSignature(data []byte) bool {
	for _, sig := range acceptedSignatures {
		if IsSignature(data, sig) {
			return true
		}
	}
	return false
}

func IsJPEG(sig Signature) bool {
	return IsSignature(sig, JPEG) || IsSignature(sig, JPEG2)
}

func IsPNG(sig Signature) bool {
	return IsSignature(sig, PNG)
}

func IsRAF(sig Signature) bool {
	return IsSignature(sig, RAF)
}

func IsImage(sig Signature) bool {
	return IsJPEG(sig) || IsPNG(sig)
}

func IsRaw(sig Signature) bool {
	return IsRAF(sig)
}
