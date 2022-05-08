package clicommands

import (
	"bytes"
	"fmt"
	"locket/aws"
	"locket/cli"
	"locket/configloader"
	"locket/constants"
	"locket/metadata"
	"locket/password"
	"locket/unix/openssl"
	"locket/unix/tar"
	"log"
	"os"
	"os/user"
	"strings"
	"time"
)

var s3Client aws.S3Client
var conf configloader.Configopts

func init() {
	conf = *configloader.Config()
	s3Client = aws.NewS3Client(conf.Auth.Aws.Profile, conf.Auth.Aws.Bucket)
}

/*
commandBackup takes all directories in the .locket.conf.yaml and uploads them to s3.
Each directory is uploaded as its own tar archive (gziped) and then encrypted in base64
*/
func commandBackup() int {
	conf := configloader.Config()

	prefix := isoDateString(time.Now().UTC())

	pw := password.GetPassword(conf.PasswordType)
	if pw == "" {
		log.Fatalf("Invalid option for \"passwordType\" in config.")
	}

	for _, dir := range conf.Directories {
		fmt.Printf("Backing up dir %s\n", expandPath(dir))

		encryptAndUploadToS3(
			expandPath(dir),
			conf.Auth.Aws.Bucket,
			conf.Auth.Aws.Profile,
			prefix,
			pw,
		)
		cli.Print("Done 🔐")
	}
	return 0
}

func encryptAndUploadToS3(dir, bucket, profile, prefix string, pw string) {

	archive := tar.Create(dir)
	encrypted := openssl.Enc(archive, pw)

	now := time.Now().UTC()

	checkfile := bytes.NewBufferString(constants.Constants.VERIFIER_FILE_CONTENTS)
	encCheckFile := openssl.Enc(checkfile, pw)
	s3Client.Upload(encCheckFile, prefix+"/"+constants.Constants.VERIFIER_FILE, map[string]string{})

	s3Client.Upload(encrypted, prefix+"/"+normalizeFilepath(dir), map[string]string{
		"Created":          isoDateString(now),
		"OriginalFilepath": dir,
		"locket-version":   metadata.ApiVersion(),
	})

}

func isoDateString(time time.Time) string {
	return fmt.Sprintf(
		"%v-%v-%vT%v:%v:%v.%vZ",
		time.Year(),
		time.Month(),
		time.Day(),
		time.Hour(),
		time.Minute(),
		time.Second(),
		time.UnixMilli(),
	)
}

func expandPath(p string) string {
	usr, _ := user.Current()
	return strings.Replace(p, "~", usr.HomeDir, 1)
}

/*
nomalizeFilepath replaces all seperators in a file path with underscores
so that its interpreted as a single filename, rather than a path.
*/
func normalizeFilepath(p string) string {
	toReplace := []string{string(os.PathSeparator), ":", "."}

	for _, char := range toReplace {
		p = strings.ReplaceAll(p, char, "_")
	}

	return p
}
