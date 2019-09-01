/*
Copyright 2017 WALLIX

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	"fmt"
	os "os"
	"path/filepath"
	// syscall "syscall"

	// "bufio"
	"regexp"

	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"

	// "io/ioutil"

	// "errors"
	"strconv"
	"strings"

	// "os/signal"

	// "golang.org/x/crypto/ssh/terminal"

	"github.com/howeyc/gopass"

	awsservices "github.com/wallix/awless/aws/services"
	"github.com/wallix/awless/global"
	"github.com/wallix/awless/database"
)

var (
	AwlessHome = filepath.Join(os.Getenv("HOME"), ".acentera")
	DBPath     = filepath.Join(AwlessHome, database.Filename)
	Dir        = filepath.Join(AwlessHome, "aws")
	KeysDir    = filepath.Join(AwlessHome, "keys")

	emailRe            = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	AwlessFirstInstall bool
)

func init() {
	os.Setenv("__AWLESS_HOME", AwlessHome)
	os.Setenv("__AWLESS_CACHE", filepath.Join(AwlessHome, "cache"))
	os.Setenv("__AWLESS_KEYS_DIR", KeysDir)
}

func InitAwlessEnv() error {
	_, err := os.Stat(DBPath)

	AwlessFirstInstall = os.IsNotExist(err)
	os.Setenv("__AWLESS_FIRST_INSTALL", strconv.FormatBool(AwlessFirstInstall))

	os.MkdirAll(KeysDir, 0700)

	var username string
	var pass string
	var enc []byte

	if AwlessFirstInstall {
		// fmt.Fprintln(os.Stderr, AWLESS_ASCII_LOGO)
		// fmt.Fprintln(os.Stderr, "Welcome! Resolving environment data...")

		if err = InitConfig(resolveRequiredConfigFromEnv()); err != nil {
			return err
		}

		err = database.Execute(func(db *database.DB) error {
			return db.SetStringValue("current.version", Version)
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "cannot store current version in db: %s\n", err)
		}
		/*
			err = database.Execute(func(db *database.DB) error {
				return db.SetStringValue("_enc", string(enc))
			})
			if err != nil {
				fmt.Fprintf(os.Stderr, "cannot store password in db: %s\n", err)
			}
			err = database.Execute(func(db *database.DB) error {
				return db.SetStringValue("user.username", username)
			})
			if err != nil {
				fmt.Fprintf(os.Stderr, "cannot store username in db: %s\n", err)
			}
		*/
	}

	if err = LoadConfig(); err != nil {
		return err
	}

	enc = []byte(GetPassword())
	username = GetUsername()
        pass = string(Decrypt(enc, global.ENC_PWD))

	if (AskUserPassword(&username, &pass)) {
		// Ok got new user/pass?
	}
        if (username == "") {
		fmt.Printf("Error: Usernamei s no  valid email address or is empty\n")
	}
        if (pass == "") {
		// ?? No passwords? exit with errors?
		fmt.Printf("Error: No password for user: [%s].\n", username)
	}
	return nil
}

func resolveRequiredConfigFromEnv() map[string]string {
	region := awsservices.ResolveRegionFromEnv()

	resolved := make(map[string]string)
	if region != "" {
		resolved[RegionConfigKey] = region
	} else {
		// Force default to us-east-1
		resolved[RegionConfigKey] = "us-east-1"
	}

	return resolved
}

func promptUntilNonEmpty(question string, v *string) {
	ask := func(v *string) bool {
		fmt.Print(question)
		_, err := fmt.Scanln(v)
		if err == nil && strings.TrimSpace(*v) != "" {
			if emailRe.MatchString(*v) {
				return false
			} else {
				fmt.Printf("Error: %s. Retry please...\n", "Invalid email address")
				return true
			}
		}
		if err != nil {
			fmt.Printf("Error: %s. Retry please...\n", err)
			fmt.Printf("You must enter a valid email address\n")
		}
		return true
	}
	for ask(v) {
	}
	return
}
func promptUntilNonEmptySecure(question string, v *string) {
	ask := func(v *string) bool {
		fmt.Print(question)
		passwd, err := gopass.GetPasswd()
		if err == nil && strings.TrimSpace(string(passwd)) != "" {
			*v = string(passwd)
			return false
		}
		if err != nil {
			if (strings.Contains(fmt.Sprintf("%s", err),"interrup")) {
				fmt.Printf("Error: %s. ...\n", err)
				return false
			}
			fmt.Printf("Error: %s. Retry please...\n", err)
		}
		return true
	}
	for ask(v) {
	}
	return
}

// Unused for now also not valid for cross compilation
/*
func getPassword(prompt string) string {
	// Get the initial state of the terminal.
	fmt.Println(prompt)
	initialTermState, e1 := terminal.GetState(syscall.Stdin)
	if e1 != nil {
		panic(e1)
	}

	// Restore it in the event of an interrupt.
	// CITATION: Konstantin Shaposhnikov - https://groups.google.com/forum/#!topic/golang-nuts/kTVAbtee9UA
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		<-c
		_ = terminal.Restore(syscall.Stdin, initialTermState)
		os.Exit(1)
	}()

	// Now get the password.
	fmt.Print(prompt)
	p, err := terminal.ReadPassword(syscall.Stdin)
	fmt.Println("")
	if err != nil {
		panic(err)
	}

	// Stop looking for ^C on the channel.
	signal.Stop(c)

	// Return the password as a string.
	return string(p)
}
*/

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func Decrypt(data []byte, passphrase string) []byte {
	if (len(data) <= 0) {
		return nil
        }
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
                return []byte(string(""))
		// panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
                return []byte(string(""))
		// panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
                return []byte(string(""))
		// panic(err.Error())
	}
	return plaintext
}

func AskUserPassword(username *string, pass *string) bool {
	fmt.Printf("\nPlease enter you credentials.\n")

	prompted := false
	if *username == "" {
	    *username = os.Getenv("ACENTERA_USERNAME")
	    if *username == "" {
		prompted = true
		promptUntilNonEmpty("\nUsername: ", username)
		Set("user.username", *username)
            }
	}
	if *pass == "" {
	  *pass = os.Getenv("ACENTERA_PASSWORD")
	  if *pass == "" {
		prompted = true
		promptUntilNonEmptySecure("Password: ", pass)
		enc := encrypt([]byte(*pass), global.ENC_PWD)
		Set("_enc", string(enc))
	  }
        }
	if (prompted) {
		Set("_token", string(""))
	}
	return prompted
}
