package main

/*

for example, when used with https://github.com/Yubico/yubikey-manager/tree/master/ykman

session:
	$(eval CODE := $(shell ykman oath code sfomuseum:aws | awk '{ print $$2 }'))
	bin/$(OS)/aws-mfa-session -profile $(USER) -code $(CODE) -duration PT8H

*/

import (
	"flag"
	_ "fmt"
	"github.com/aaronland/go-aws-tools/auth"
	"github.com/aaronland/go-aws-tools/cli"
	"github.com/aaronland/go-aws-tools/config"
	"github.com/whosonfirst/iso8601duration"
	"log"
	"strings"
	"time"
)

func main() {

	profile := flag.String("profile", "default", "A valid AWS credentials profile")
	code := flag.String("code", "", "A valid MFA code. If empty the application will block and prompt the user")
	session_profile := flag.String("session-profile", "session", "The name of the AWS credentials profile to update with session credentials")
	session_duration := flag.String("duration", "PT1H", "A valid ISO8601 duration string indicating how long the session should last (months are currently not supported)")

	flag.Parse()

	d, err := duration.FromString(*session_duration)

	if err != nil {
		log.Fatalf("Failed to parse session duration, %v", err)
	}

	ttl := d.ToDuration().Seconds()
	ttl_32 := int32(ttl)

	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatalf("Failed to create new config, %v", err)
	}

	aws_cfg, err := cfg.AWSConfigWithProfile(*profile)

	if err != nil {
		log.Fatalf("Failed to create AWS config, %v", err)
	}

	*code = strings.TrimSpace(*code)

	if *code == "" {

		*code = cli.Readline("Enter your MFA token code:")

		if *code == "" {
			log.Fatalf("Missing MFA code")
		}
	}

	creds, err := auth.GetCredentialsWithMFA(aws_cfg, *code, ttl_32)

	if err != nil {
		log.Fatalf("Failed to get credentials with MFA, %v", err)
	}

	err = cfg.SetSessionCredentialsWithProfile(*session_profile, creds)

	if err != nil {
		log.Fatalf("Failed to get credentials with session profile, %v", err)
	}

	now := time.Now()
	then := now.Add(d.ToDuration())

	log.Printf("Updated session credentials for '%s' profile, expires %s (%s)\n", *session_profile, then.Format(time.Stamp), *creds.Expiration)
}
