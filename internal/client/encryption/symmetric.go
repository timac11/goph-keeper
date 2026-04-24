package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/binary"
	"io"
	"os"
)

const chunkSize = 4 * 1024 * 1024

func generateKey() ([]byte, error) {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}
	return key, nil
}

func generateNonce(counter uint64) []byte {
	nonce := make([]byte, 12)
	binary.BigEndian.PutUint64(nonce[4:], counter)
	return nonce
}

func EncryptFile(in, out *os.File) (*[]byte, error) {
	key, err := generateKey()
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, chunkSize)
	var counter uint64

	for {
		n, err := in.Read(buf)

		if n > 0 {
			nonce := generateNonce(counter)

			ciphertext := gcm.Seal(nil, nonce, buf[:n], nil)

			if err := binary.Write(out, binary.BigEndian, uint32(len(ciphertext))); err != nil {
				return nil, err
			}

			if _, err := out.Write(ciphertext); err != nil {
				return nil, err
			}

			counter++
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
	}

	return &key, nil
}

func DecryptFile(in, out *os.File, key []byte) error {
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	var counter uint64

	for {
		var chunkLen uint32

		err := binary.Read(in, binary.BigEndian, &chunkLen)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		ciphertext := make([]byte, chunkLen)

		if _, err := io.ReadFull(in, ciphertext); err != nil {
			return err
		}

		nonce := generateNonce(counter)

		plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
		if err != nil {
			return err
		}

		if _, err := out.Write(plaintext); err != nil {
			return err
		}

		counter++
	}

	return nil
}
