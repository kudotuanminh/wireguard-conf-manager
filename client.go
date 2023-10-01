package main

import (
	"database/sql"
	"fmt"

	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

type Client struct {
	ProfileID   int    `json:"ProfileID"`
	ProfileName string `json:"ProfileName"`
	ClientIP    string `json:"ClientIP"`
	PrivateKey  string `json:"PrivateKey"`
	PublicKey   string `json:"PublicKey"`
	AllowedIPs  string `json:"AllowedIPs"`
}

func getProfiles(db *sql.DB) ([]Client, error) {
	rows, err := db.Query("SELECT * FROM wg_conf_clients")
	if err == nil {
		defer rows.Close()
		var profiles []Client
		for rows.Next() {
			var profile Client
			err = rows.Scan(&profile.ProfileID, &profile.ProfileName, &profile.ClientIP, &profile.PrivateKey, &profile.PublicKey, &profile.AllowedIPs)
			profiles = append(profiles, profile)
		}
		return profiles, err
	}
	return nil, err
}

func getProfile(db *sql.DB, profileName string) (Client, error) {
	var profile Client
	err := db.QueryRow("SELECT * FROM wg_conf_clients WHERE Profile_Name = ?", profileName).Scan(&profile.ProfileID, &profile.ProfileName, &profile.ClientIP, &profile.PrivateKey, &profile.PublicKey, &profile.AllowedIPs)
	return profile, err
}

func newProfile(db *sql.DB, profileName string, clientIP string, allowedIPs string) (Client, error) {
	privkey, pubkey, err := genKey()
	if err == nil {
		_, err := db.Exec("INSERT INTO wg_conf_clients (Profile_Name, Client_IP, Private_Key, Public_Key, Allowed_IPs) VALUES (?, ?, ?, ?, ?)", profileName, clientIP, privkey, pubkey, allowedIPs)
		profile, _ := getProfile(db, profileName)
		return profile, err
	}
	return Client{}, err
}

func genKey() (string, string, error) {
	privKey, err := wgtypes.GeneratePrivateKey()
	pubKey := privKey.PublicKey()
	// fmt.Println("New Private Key: ", privKey.String(), "\nNew Public Key: ", pubKey.String())
	return privKey.String(), pubKey.String(), err
}

func deleteProfile(db *sql.DB, profileName string) error {
	_, err := getProfile(db, profileName)
	if err == nil {
		_, err := db.Exec("DELETE FROM wg_conf_clients WHERE Profile_Name = ?", profileName)
		if err == nil {
			fmt.Println("Deleted profile: ", profileName)
		}
		return err
	}
	return err
}
