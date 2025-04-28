package main

import (
	"fmt"
	"log"
	"os"
	"github.com/urfave/cli/v2"
	"time"
	"net"
)

 
//checking the health of a website
func main(){
	app := &cli.App{
		Name:"Healthchecker",
		Usage:"A tiny tool that checks whether a website is running or not",
		Flags:[]cli.Flag{//All i see is there are two flags initialise
			&cli.StringFlag{
			Name:"domain",//one name domain
			Aliases:[]string{"d"},
			Usage:"Domain name to check",
			Required:true,
			},
			&cli.StringFlag{
				Name:"port",//another name port
				Aliases:[]string{"p"},
				Usage:"Port number to check.",
				Required:true,
			},
		},
		Action:func(c *cli.Context) error {//wrote some action
			port:= c.String("port")
			if "port"  == ""{
				port="80"
			}
			status := Check(c.String("domain"),port)
			fmt.Println(status)
			return nil 
		},
	}
	err:=app.Run(os.Args)
	if err!=nil{
		log.Fatal(err)
	}
}

func Check(destination string,port string ) string{
	adress := destination + ":"+ port
	timeout := time.Duration(5 * time.Second)

	conn,err := net.DialTimeout("tcp",adress,timeout)
	var status string

	if err!=nil {
		status = fmt.Sprintf("[DOWN] %v is unreachable, \n %v",destination,err)
	}else{
		status = fmt.Sprintf("[UP] %v is reachable,\n From:%v\n To:%v",destination,conn.LocalAddr(),conn.RemoteAddr())
	}

	return status
}
