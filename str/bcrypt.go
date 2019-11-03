/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 *      _ _
 *     | (_)_ __   __ _
 *  _  | | | '_ \ / _` |
 * | |_| | | | | | (_| |
 *  \___/|_|_| |_|\__, |                        https://github.com/yingzhuo/jing
 *                |___/
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
package str

import (
	"golang.org/x/crypto/bcrypt"
)

func BCrypt(s string) string {
	if h, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost); err != nil {
		panic(err)
	} else {
		return string(h)
	}
}

func CheckBCrypt(raw, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(raw), []byte(hashed))
	return err != nil
}
