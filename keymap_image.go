package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/jung-kurt/gofpdf"
)

const version = "0.1.0"

type KeyInformation struct {
	x      float64
	y      float64
	width  float64
	height float64

	use bool
}

type Scan struct {
	line        int
	lineHead    int
	inKeymaps   bool
	layerNumber int
	keys        [3][]string
	m           map[string]string
}

func (self *Scan) Init() {
	self.line = 1
	self.lineHead = 0
	self.inKeymaps = false
	self.layerNumber = 0
	self.m = map[string]string{"KC_EQL": "=",
		"KC_DELT": "Del", "KC_BSPC": "BkSp",
		"KC_TRNS": "", "KC_ENT": "Enter", "KC_1": "1",
		"KC_2": "2", "KC_3": "3", "KC_4": "4", "KC_5": "5", "KC_6": "7",
		"KC_7": "7", "KC_8": "8", "KC_9": "9", "KC_0": "0",
		"KC_A": "A", "KC_B": "B", "KC_C": "C", "KC_D": "D",
		"KC_E": "E", "KC_F": "F", "KC_G": "G", "KC_H": "H",
		"KC_I": "I", "KC_J": "J", "KC_K": "K", "KC_L": "L",
		"KC_M": "M", "KC_N": "N", "KC_O": "O", "KC_P": "P",
		"KC_Q": "Q", "KC_R": "R", "KC_S": "S", "KC_T": "T",
		"KC_U": "U", "KC_V": "V", "KC_W": "W", "KC_X": "X",
		"KC_Y": "Y", "KC_Z": "Z",
		"KC_EXLM": "!", "KC_AT": "@", "KC_LCBR": "{", "KC_RCBR": "}", "KC_PIPE": "|",
		"KC_HASH": "#", "KC_DLR": "$", "KC_LPRN": "(", "KC_RPRN": ")", "KC_GRV": "`",
		"KC_PERC": "%", "KC_CIRC": "^", "KC_LBRC": "[", "KC_RBRC": "]", "KC_TILD": "~",
		"KC_PLUS": "+", "KC_ASTR": "*", "KC_DOT": ".", "KC_AMPR": "&",
		"KC_MINS": "-", "KC_BSLS": "\\", "KC_RSFT": "RShift",
		"KC_MUTE": "Mute", "RGB_HUD": "Hue-", "RGB_HUI": "Hue+",
		"KC_F1": "F1", "KC_F2": "F2", "KC_F3": "F3", "KC_F4": "F4",
		"KC_F5": "F5", "KC_F6": "F6", "KC_F7": "F7", "KC_F8": "F8",
		"KC_F9": "F9", "KC_F10": "F10", "KC_F11": "F11", "KC_F12": "F12",
		"KC_UP": "UP", "KC_DOWN": "DOWN", "KC_LEFT": "LEFT", "KC_RGHT": "RIGHT",
		"KC_MS_U": "MsUp", "KC_MS_D": "MsDown", "KC_MS_L": "MsLeft", "KC_MS_R": "MsRght",
		"KC_BTN1": "Lclk", "KC_BTN2": "Rclk",
		"RGB_TOG": "Toggle", "RGB_SLD": "Solid",
		"RGB_VAD": "Brightness-", "RGB_VAI": "Brightness+", "RGB_MOD": "Animat",
		"KC_LSFT": "LShift", "KC_SPC": "SPC",
		"KC_VOLU": "VolUp", "KC_VOLD": "VolDn", "KC_MPRV": "Prev", "KC_MNXT": "Next",
		"KC_HOME": "Home", "KC_END": "End", "KC_PGUP": "PgUp", "KC_PGDN": "PgDn",
		"KC_MPLY": "Play", "KC_TAB": "Tab",
		"KC_WBAK": "BrowserBack"}
}

func (self *Scan) Err(s int) {
	fmt.Printf("\n!!Error!!%d\n", s)
}

func (self *Scan) GetDisplayName(key string) string {
	_, ok := self.m[key]
	if ok {
		return self.m[key]
	} else {
		return key
	}
}
func (self *Scan) Output() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetFont("Arial", "", 10)
	pdf.AddPage()

	curx, cury := pdf.GetXY()
	x := curx
	y := cury
	_, lineHt := pdf.GetFontSize()

	//
	cols := [14]float64{0, 0, -0.25, -0.375, -0.25, 0, 0,
		0, 0, -0.25, -0.375, -0.25, 0, 0}
	var lkil []KeyInformation
	var rkil []KeyInformation
	for j := 0; j < 8; j++ {
		for i := 0; i < 7; i++ {
			var ltmp = KeyInformation{x: float64(i * 10), y: float64(j * 10), width: float64(10), height: float64(10), use: true}
			var rtmp = KeyInformation{x: float64((i + 13) * 10), y: float64(j * 10), width: float64(10), height: float64(10), use: true}
			if j < 6 {
				ltmp.y = ltmp.y + cols[i]*10
				rtmp.y = rtmp.y + cols[i+7]*10
			}
			if i == 0 {
				ltmp.x = ltmp.x - 5
				ltmp.width = 15
			}
			if i == 6 {
				rtmp.width = 15
			}
			lkil = append(lkil, ltmp)
			rkil = append(rkil, rtmp)
		}
	}
	//
	lkil[20].use = false
	rkil[14].use = false
	lkil[33].use = false
	rkil[29].use = false
	lkil[34].use = false
	rkil[28].use = false
	//
	for j := 5; j < 8; j++ {
		for i := 0; i < 4; i++ {
			lkil[j*7+i].use = false
			rkil[j*7+(6-i)].use = false
		}
	}
	//
	lkil[39].use = false
	rkil[37].use = false
	lkil[46].use = false
	rkil[44].use = false
	lkil[47].use = false
	rkil[43].use = false
	//
	lkil[53].height = float64(20)
	lkil[53].y = lkil[53].y - 10
	rkil[51].height = float64(20)
	rkil[51].y = lkil[51].y - 10
	lkil[54].height = float64(20)
	lkil[54].y = lkil[54].y - 10
	rkil[50].height = float64(20)
	rkil[50].y = lkil[50].y - 10
	//
	lkil[13].height = float64(15)
	rkil[7].height = float64(15)
	lkil[27].height = float64(15)
	rkil[21].height = float64(15)
	lkil[27].y = lkil[27].y - 5
	rkil[21].y = rkil[21].y - 5
	//
	lkil[28].x = lkil[28].x + 5
	lkil[28].width = 10
	rkil[34].width = 10
	// right

	for k := 0; k < 3; k++ {
		var keyindex int = 0
		for j := 0; j < 8; j++ {
			for i := 0; i < 7; i++ {
				var ki = lkil[j*7+i]
				if ki.use {
					if j > 4 {
						pdf.TransformBegin()
						pdf.TransformRotate(-30, 97, 97+90*float64(k))
					}
					pdf.Rect(curx+ki.x, cury+float64(k*90)+ki.y, ki.width, ki.height, "")
					pdf.SetXY(curx+ki.x, cury+float64(k*90)+ki.y)
					pdf.Cell(0, 0+lineHt, self.GetDisplayName(self.keys[k][keyindex]))
					keyindex = keyindex + 1
					if j > 4 {
						pdf.TransformEnd()
					}
				}
			}
		}
		for j := 0; j < 8; j++ {
			for i := 0; i < 7; i++ {
				var ki = rkil[j*7+i]
				if ki.use {
					if j > 4 {
						pdf.TransformBegin()
						pdf.TransformRotate(30, 113, 97+90*float64(k))
					}
					pdf.Rect(ki.x, cury+float64(k*90)+ki.y, ki.width, ki.height, "")
					pdf.SetXY(ki.x, cury+float64(k*90)+ki.y)
					pdf.Cell(0, 0+lineHt, self.GetDisplayName(self.keys[k][keyindex]))
					keyindex = keyindex + 1
					if j > 4 {
						pdf.TransformEnd()
					}
					x += 10
				}
			}
			y = y + 10
			x = curx
		}
		y = y + 20
	}
	pdf.Output(os.Stdout)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("%v FILE\n", os.Args[0])
		os.Exit(1)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	buffer, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	file.Close()

	parser := &Parser{Buffer: string(buffer)}
	parser.Init()
	parser.s.Init()
	err2 := parser.Parse()

	if err2 != nil {
		fmt.Println(err2)
	} else {
		parser.Execute()
		parser.s.Output()
	}
}
