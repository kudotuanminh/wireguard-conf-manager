package main

import (
	"database/sql"
	"os"
)

func genClientConf(db *sql.DB, profileName string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err == nil {
		var filePath = homeDir + "/" + profileName + ".conf"
		f, err := os.Create(filePath)
		if err == nil {
			profile, err := getProfile(db, profileName)
			if err == nil {
				f.WriteString("[Interface]\n")
				f.WriteString("PrivateKey = " + profile.PrivateKey + "\n")
				f.WriteString("Address = " + profile.ClientIP + "\n\n")
				f.WriteString("[Peer]\n")
				f.WriteString("PublicKey = " + os.Getenv("SERVER_PUBKEY") + "\n")
				f.WriteString("AllowedIPs = " + profile.AllowedIPs + "\n")
				f.WriteString("Endpoint = " + os.Getenv("SERVER_IP") + ":" + os.Getenv("SERVER_PORT") + "\n")
				f.WriteString("PersistentKeepalive = 25\n")
				defer f.Close()
				return filePath, err
			}
			return "", err
		}
		return "", err
	}
	return "", err
}

func genServerConf(db *sql.DB) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err == nil {
		var filePath = homeDir + "/server.conf"
		// var filePath = "/etc/wireguard/wg0.conf"
		f, err := os.Create(filePath)
		if err == nil {
			profiles, err := getProfiles(db)
			if err == nil {
				f.WriteString("[Interface]\n")
				f.WriteString("PrivateKey = " + os.Getenv("SERVER_PRIVKEY") + "\n")
				f.WriteString("Address = 192.168.1.1/32\n")
				f.WriteString("ListenPort = " + os.Getenv("SERVER_PORT") + "\n\n")
				for _, profile := range profiles {
					f.WriteString("// " + profile.ProfileName + "\n")
					f.WriteString("[Peer]\n")
					f.WriteString("PublicKey = " + profile.PublicKey + "\n")
					f.WriteString("AllowedIPs = " + profile.AllowedIPs + "\n\n")
				}
				defer f.Close()
				return filePath, err
			}
			return "", err
		}
		return "", err
	}
	return "", err
}
