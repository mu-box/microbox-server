package util

import (
	"strconv"

	"github.com/mu-box/golang-lvs"
	"github.com/mu-box/microbox-server/config"
)

// make sure the router is being forwarded
func init() {
	lvs.DefaultIpvs.Save()
	AddForward("80", config.IP, config.Ports["router"])
	AddForward("443", config.IP, config.Ports["router"])
}

// add a server into the lvs system
func AddForward(fromPort, toIp, toPort string) error {
	fromInt, err := strconv.Atoi(fromPort)
	if err != nil {
		config.Log.Error("error: %s\n", err.Error())
		return err
	}
	err = lvs.DefaultIpvs.AddService(lvs.Service{Host: config.IP, Port: fromInt, Type: "tcp", Persistence: 300})
	if err != nil {
		config.Log.Error("error: %s\n", err.Error())
		return err
	}
	// look up the service instead of using the one we created
	// this keeps the communication with the lvs package simple
	service := lvs.DefaultIpvs.FindService("tcp", config.IP, fromInt)

	toInt, _ := strconv.Atoi(toPort)
	server := lvs.Server{Host: toIp, Port: toInt, Weight: 1, Forwarder: "m"}
	err = service.AddServer(server)
	if err != nil {
		config.Log.Error("error: %s\n", err.Error())
		return err
	}
	return nil
}

func RemoveForward(ip string) error {
	for _, service := range lvs.DefaultIpvs.Services {
		for _, server := range service.Servers {
			if server.Host == ip {
				err := lvs.DefaultIpvs.RemoveService(service.Type, service.Host, service.Port)
				if err != nil {
					config.Log.Error("error: %s\n", err.Error())
					return err
				}
			}
		}
	}
	return nil

	// vips, err := lvs.ListVips()
	// if err != nil {
	// 	return err
	// }

	// errorString := ""

	// for _, vip := range vips {
	// 	for _, server := range vip.Servers {
	// 		if server.Host == ip {
	// 			err := lvs.DeleteVip(fmt.Sprintf("%s:%d", vip.Host, vip.Port))
	// 			if err != nil {
	// 				errorString = fmt.Sprintf("%s%v\n", errorString, err.Error())
	// 			}
	// 			break
	// 		}
	// 	}
	// }

	// if errorString != "" {
	// 	return fmt.Errorf(errorString)
	// }
	// return nil
}

func ListVips() ([]lvs.Service, error) {
	return lvs.DefaultIpvs.Services, nil
}
