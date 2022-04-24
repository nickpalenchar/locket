package clicommands

import (
	"fmt"
	"locket/aws"
	"locket/cli"
	"locket/configloader"
	"locket/metadata"
	"locket/unix/openssl"
	"locket/unix/tar"
	"os"
	"os/user"
	"strings"
	"time"
)

/*
commandBackup takes all directories in the .locket.conf.yaml and uploads them to s3.
Each directory is uploaded as its own tar archive (gziped) and then encrypted in base64
*/
func commandBackup() int {
	conf := configloader.Config()

	prefix := isoDateString(time.Now().UTC())

	for _, dir := range conf.Directories {
		fmt.Printf("Backing up dir %s\n", expandPath(dir))

		encryptAndUploadToS3(
			expandPath(dir),
			conf.Auth.Aws.Bucket,
			conf.Auth.Aws.Profile,
			prefix,
		)
		cli.Print("Done 🔐")
	}
	return 0
}

func encryptAndUploadToS3(dir, bucket, profile, prefix string) {
	archive := tar.Create(dir)
	encrypted := openssl.Enc(archive, "thisisatester2888kd89od80228de<3@")

	now := time.Now().UTC()

	aws.UploadToS3(encrypted, bucket, profile, prefix+"/"+normalizeFilepath(dir), map[string]string{
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
