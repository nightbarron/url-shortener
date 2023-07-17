package services

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"url-shortener/configs"
	"url-shortener/helpers"
)

var NginxVhostPath = "/etc/nginx/conf.d/urlShorten.conf"

func ApplyToNginx(config configs.GlobalConfig, shortUrlHash, longUrl string) error {
	// Open the file
	_, err := os.Stat(NginxVhostPath)
	if os.IsNotExist(err) {
		log.Info("File " + NginxVhostPath + " does not exist, creating new one")
		// If the file does not exist, create it from urlShort.tml
		template, _ := helpers.ReadFile("templates/urlShort.tml")
		replacer := strings.NewReplacer("{{ domain }}", strings.Split(config.ShortUrlDomain, "/")[2])
		defaultVhost := replacer.Replace(template)
		// Create the file
		_, _ = os.Create(NginxVhostPath)
		log.Info("File " + NginxVhostPath + " created")

		// Write the new content back to the file
		err = ioutil.WriteFile(NginxVhostPath, []byte(defaultVhost), 755)
		if err != nil {
			return err
		}

	}

	file, err := os.Open(NginxVhostPath)
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// Read the file
	dataVhost, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	locationStrTml, _ := helpers.ReadFile("templates/locationStr.tml")
	replacer := strings.NewReplacer("{{ shortUrl }}", shortUrlHash, "{{ longUrl }}", longUrl)
	locationStr := replacer.Replace(locationStrTml)

	// Add location to vhost
	replacer = strings.NewReplacer("# {{ location }}", locationStr)

	newDataVhost := replacer.Replace(string(dataVhost))

	// Write the new content back to the file
	err = ioutil.WriteFile(NginxVhostPath, []byte(newDataVhost), 755)
	if err != nil {
		return err
	}

	// Get docker container name nginx
	cmd := exec.Command("docker", "ps", "--format", "{{.Names}}", "--filter", "name=nginx")
	out, err := cmd.Output()
	if err != nil {
		return err
	}
	containerName := strings.Split(string(out), "\n")[0]

	// Reload the nginx
	cmd = exec.Command("docker", "exec", containerName, "nginx", "-s", "reload")
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
