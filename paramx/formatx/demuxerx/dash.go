package demuxerx

// 20.6 dash  DASH
type DASH struct {
	CencDecryptionKey bool `json:"cenc_decryption_key" flag:"-cenc_decryption_key"`
	//16-byte key, in hex, to decrypt files encrypted using ISO Common Encryption (CENC/AES-128 CTR; ISO/IEC 23001-7).
}
