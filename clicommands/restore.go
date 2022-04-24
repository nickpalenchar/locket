package clicommands

import (
	"bytes"
	"fmt"
	"io"
	"locket/aws"
	"locket/cli"
	"locket/configloader"
	"locket/stringutil"
	"locket/unix/openssl"
	"locket/unix/tar"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/s3/types"

	"golang.org/x/exp/slices"
)

func commandRestore() int {
	cfg := configloader.Config()
	fullObjs := aws.ListObjects(cfg.Auth.Aws.Bucket, cfg.Auth.Aws.Profile)

	if len(fullObjs) == 0 {
		fmt.Print("There are no backups available.\n")
		return 0
	}

	s3dirs := parseTopLevelObjects(fullObjs)

	cli.Print("Select from the following backups:")
	for k, v := range s3dirs {
		cli.Print(fmt.Sprintf("    %v) %v", k+1, v))
	}
	selected := cli.Prompt("enter a number > ")
	selectedInt, err := strconv.ParseInt(selected, 10, 64)
	if err != nil {
		log.Fatal("Wrong input. Rerun and try again")
	}
	selectedDir := s3dirs[selectedInt-1]
	fmt.Printf("you selected %s\n", selectedDir)

	cli.Print("Enter the location of where you'd want the restore to go:")
	cli.Print("  . for the current working directory")
	cli.Print("  / to restore to the *original* filepaths")

	targetDir := cli.Prompt("location > ")
	targetDir = strings.TrimSpace(targetDir)
	reportRestorePlan(targetDir)

	restoreAllWithPrefix(cfg.Auth.Aws.Bucket, selectedDir, targetDir, fullObjs, cfg)

	cli.Print("restore complete.")

	return 0
}

func reportRestorePlan(downloadLocation string) {
	if downloadLocation == "/" {
		cli.Print("\n***** DANGEROUS OPERATON AHEAD -- PLEASE READ *****")
		cli.Print("You chose `/`, which restores to the *original location.")
		cli.Print("This will overwrite any data in newer files since the backup, if any!")
		cli.Print("\n Type `yes` to continue, otherwise the operation will abort")
		answer := cli.Prompt("yes to proceed > ")
		if strings.ToLower(answer) != "yes" {
			cli.Print("Did not say yes. Exiting.")
			os.Exit(0)
		}
	} else {
		cli.Print("locket will restore to the following location: ")
		backupPath, _ := filepath.Abs(downloadLocation)
		cli.Print(fmt.Sprintf("    %s", backupPath))
		cli.Print("Type `y` to continue.")
		for {
			answer := strings.ToLower(cli.Prompt("y/n > "))
			if answer == "y" {
				return
			}
			if answer == "n" {
				cli.Print("Exiting")
				os.Exit(0)
			}
			cli.Print("Type `y` or `n`")
		}
	}
}

func restoreAllWithPrefix(bucket, prefix, targetDir string, fullObjs []types.Object, cfg *configloader.Configopts) (result []*bytes.Buffer) {
	pw := "thisisatester2888kd89od80228de<3@"
	for _, obj := range fullObjs {
		if strings.HasPrefix(*obj.Key, prefix) {
			data := aws.DownloadFromS3(bucket, cfg.Auth.Aws.Profile, *obj.Key)
			decrypted := decryptDataWithPassword(data, pw)
			err := tar.Extract(decrypted, targetDir)
			if err != nil {
				cli.Print(fmt.Sprintf("[Error]: %s", err))
			}
			result = append(result, decrypted)
		}
	}
	//	data := aws.DownloadFromS3(cfg.Auth.Aws.Bucket, cfg.Auth.Aws.Profile, fmt.Sprintf("%s/*", selectedDir))
	return result
}

/* decryptDataWithPassword decrypts base64 data using a provided password. */
func decryptDataWithPassword(data io.Reader, pw string) *bytes.Buffer {
	return openssl.Dec(data, pw)
}

func parseTopLevelObjects(objs []types.Object) []string {

	var result []string

	for _, obj := range objs {
		topLevel := stringutil.PathPos(*obj.Key, 0)
		if !slices.Contains(result, topLevel) {
			result = append(result, topLevel)
		}
	}
	fmt.Printf("result %v", result)
	return result

}
