package banner

import (
	"echo-demo-print-routes/config"
	"fmt"
	"github.com/labstack/echo/v4"
	"os"
	"strconv"
	"strings"
)

const banner = `
    ____                          
   / __ \  ___    ____ ___   ____ 
  / / / / / _ \  / __ '__ \ / __ \
 / /_/ / /  __/ / / / / / // /_/ /
/_____/  \___/ /_/ /_/ /_/ \____/ 

                 ..                             
                 '-.                            
                  --'                           
                  ..-                           
                 '-:'   '                       
             '  -/:-  -:'                       
           '/: ////' :/:                        
           -o+/+++/-'++/.                       
         ..-s+++++++/+++/-                      
         //:sooooooooooo++:                     
        :ooooooooooooooooo+:                    
      ''ssssssssssssssssssso .    %s              
      :/osssssssssssssssssso+/    %s              
      -yyyyyyyyyyyyyyyyyyyyyy-    %s              
       oyyyyyyyyyyyyyyyyyyyyo     %s              
        +yyyyyyyyyyyyyyyyyyo'     %s              
         .+yyyyyyyyyyyyyys:                     
            .:/+ooossys/.                       
               '....'
`

// show banner only for dev environment.
func ShowBanner(e *echo.Echo, addr string) {

	if config.GetEnvironment() == "prod" {
		return
	}
	e.HideBanner = true
	e.HidePort = true

	routes := e.Routes()

	var logo string
	logo += "%s"
	logo += " ┌─────────────────────────────────────────────────────┐\n"
	logo += " │ %s │\n"
	logo += " │ %s │\n"
	logo += " │                                                     │\n"
	logo += " │ Handlers %s  Processes %s │\n"
	logo += " │ Prefork .%s  PID ....%s │\n"
	logo += " └─────────────────────────────────────────────────────┘"
	logo += "%s"

	const (
		cBlack   = "\u001b[90m"
		cRed     = "\u001b[91m"
		cCyan    = "\u001b[96m"
		cGreen   = "\u001b[92m"
		cYellow  = "\u001b[93m"
		cBlue    = "\u001b[94m"
		cMagenta = "\u001b[95m"
		cWhite   = "\u001b[97m"
		cReset   = "\u001b[0m"
	)
	var (
		borderColor   = cBlack
		endPointColor = cCyan
		methodColor   = cMagenta
	)

	align := func(s string, width int, direction string, color string) string {
		pad := width - len(s)
		str := ""
		for i := 0; i < pad; i++ {
			str += " "
		}
		if direction == "left" {
			v := s + str
			str = fmt.Sprintf(" %s%s%s", color, v, cBlack)
		} else {
			str += fmt.Sprintf("%s%s%s", color, s, cBlack)
		}
		return str
	}

	host, port := parseAddr(addr)
	if host == "" || host == "0.0.0.0" {
		host = "127.0.0.1"
	}
	addr = "http://" + host + ":" + port

	fmt.Printf(cBlack+banner,
		fmt.Sprintf(align("running on", 15, "left", endPointColor)+"%s", align(addr, 26, "", endPointColor)),
		fmt.Sprintf(align("Framework", 15, "left", endPointColor)+"%s", align("Echo v"+echo.Version, 26, "", endPointColor)),
		fmt.Sprintf(align("Database", 15, "left", endPointColor)+"%s", align("Mongo v1.4.6", 26, "", endPointColor)),
		fmt.Sprintf(align("Http Client", 15, "left", endPointColor)+"%s", align("go-resty v2.4.0", 26, "", endPointColor)),
		fmt.Sprintf(align("PID", 15, "left", endPointColor)+"%s", align(strconv.Itoa(os.Getpid()), 26, "", endPointColor)),
	)

	mRoutes := make(map[string][]*echo.Route)
	// group the routes by name to print them in order.
	for _, route := range routes {
		initialPath := strings.Split(route.Path, "/")[1]
		mRoutes[initialPath] = append(mRoutes[initialPath], route)
	}
	var grp string

	for _, arr := range mRoutes {
		v := " ┌─────────────────────────────────────────────────────┐\n"
		v += " |  Path:                                      Method: |\n"
		grp += fmt.Sprintf("%s%s%s", borderColor, v, cBlack)

		str := ""
		for _, route := range arr {

			vle := fmt.Sprintf("%s %s", align(route.Path, 42, "left", cBlue), align(route.Method, 7, "right", methodColor))
			str += fmt.Sprintf(
				fmt.Sprintf("%s%s%s", borderColor, " |", cBlack)+
					"%s "+
					fmt.Sprintf("%s%s%s", borderColor, "|\n", cBlack),
				align(vle, 60, "left", ""))
		}
		grp += str

		v2 := " └─────────────────────────────────────────────────────┘\n"
		grp += fmt.Sprintf("%s%s%s", borderColor, v2, cBlack)
	}
	fmt.Println(grp)

}

func parseAddr(raw string) (host, port string) {
	if i := strings.LastIndex(raw, ":"); i != -1 {
		return raw[:i], raw[i+1:]
	}
	return raw, ""
}
