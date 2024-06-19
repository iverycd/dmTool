package info

import (
	"fmt"
	"github.com/fatih/color"
)

var ver = "0.0.2"

func Info() {
	color.Red("DDDDDDDDDDDDD      BBBBBBBBBBBBBBBBB               AAA                  GGGGGGGGGGGGG     OOOOOOOOO     DDDDDDDDDDDDD        ")
	color.Red("D::::::::::::DDD   B::::::::::::::::B             A:::A              GGG::::::::::::G   OO:::::::::OO   D::::::::::::DDD     ")
	color.Red("D:::::::::::::::DD B::::::BBBBBB:::::B           A:::::A           GG:::::::::::::::G OO:::::::::::::OO D:::::::::::::::DD   ")
	color.Red("DDD:::::DDDDD:::::DBB:::::B     B:::::B         A:::::::A         G:::::GGGGGGGG::::GO:::::::OOO:::::::ODDD:::::DDDDD:::::D  ")
	color.Red("  D:::::D    D:::::D B::::B     B:::::B        A:::::::::A       G:::::G       GGGGGGO::::::O   O::::::O  D:::::D    D:::::D ")
	color.Red("  D:::::D     D:::::DB::::B     B:::::B       A:::::A:::::A     G:::::G              O:::::O     O:::::O  D:::::D     D:::::D")
	color.Red("  D:::::D     D:::::DB::::BBBBBB:::::B       A:::::A A:::::A    G:::::G              O:::::O     O:::::O  D:::::D     D:::::D")
	color.Red("  D:::::D     D:::::DB:::::::::::::BB       A:::::A   A:::::A   G:::::G    GGGGGGGGGGO:::::O     O:::::O  D:::::D     D:::::D")
	color.Red("  D:::::D     D:::::DB::::BBBBBB:::::B     A:::::A     A:::::A  G:::::G    G::::::::GO:::::O     O:::::O  D:::::D     D:::::D")
	color.Red("  D:::::D     D:::::DB::::B     B:::::B   A:::::AAAAAAAAA:::::A G:::::G    GGGGG::::GO:::::O     O:::::O  D:::::D     D:::::D")
	color.Red("  D:::::D     D:::::DB::::B     B:::::B  A:::::::::::::::::::::AG:::::G        G::::GO:::::O     O:::::O  D:::::D     D:::::D")
	color.Red("  D:::::D    D:::::D B::::B     B:::::B A:::::AAAAAAAAAAAAA:::::AG:::::G       G::::GO::::::O   O::::::O  D:::::D    D:::::D ")
	color.Red("DDD:::::DDDDD:::::DBB:::::BBBBBB::::::BA:::::A             A:::::AG:::::GGGGGGGG::::GO:::::::OOO:::::::ODDD:::::DDDDD:::::D  ")
	color.Red("D:::::::::::::::DD B:::::::::::::::::BA:::::A               A:::::AGG:::::::::::::::G OO:::::::::::::OO D:::::::::::::::DD   ")
	color.Red("D::::::::::::DDD   B::::::::::::::::BA:::::A                 A:::::A GGG::::::GGG:::G   OO:::::::::OO   D::::::::::::DDD     ")
	color.Red("DDDDDDDDDDDDD      BBBBBBBBBBBBBBBBBAAAAAAA                   AAAAAAA   GGGGGG   GGGG     OOOOOOOOO     DDDDDDDDDDDDD        ")
	colorStr := color.New()
	colorStr.Add(color.FgHiGreen)
	colorStr.Printf("VastBase export and import Tool\n")
	colorStr.Printf("Powered By: DBA Team Of Infrastructure Research Center \nRelease version v" + ver)
	fmt.Printf("\n")
}
