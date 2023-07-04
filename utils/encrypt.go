package utils

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"strings"

	"golang.org/x/crypto/argon2"
)

// HashPassword 为密码加密或者其他
func HashPassword(password string) (string, error) {
	// 定义参数
	salt := make([]byte, 16) // 随机生成盐值
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	// 将盐值和哈希值拼接为最终结果
	var builder strings.Builder
	builder.WriteString("$argon2id$v=19$m=65536,t=1,p=4$")
	builder.WriteString(base64.RawStdEncoding.EncodeToString(salt))
	builder.WriteString("$")
	builder.WriteString(base64.RawStdEncoding.EncodeToString(hash))
	//result := fmt.Sprintf("$argon2id$v=19$m=65536,t=1,p=4$%s$%s", base64.RawStdEncoding.EncodeToString(salt), base64.RawStdEncoding.EncodeToString(hash))
	return builder.String(), nil
}

// VerifyPassword 验证密码
func VerifyPassword(hashedPassword, password string) error {
	parts := strings.Split(hashedPassword, "$")
	if len(parts) != 6 {
		return errors.New("Invalid hashed password format")
	}
	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return err
	}
	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	originHash, _ := base64.RawStdEncoding.DecodeString(parts[5])
	if !bytes.Equal(hash, originHash) {
		return errors.New("Invalid password")
	}
	return nil
}
