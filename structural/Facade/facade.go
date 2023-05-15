package Facade

import "fmt"

// 门面模式/外观模式
// 将所有内部的东西都封装在门面里面，外面对里面的具体实现无感知

type Screen struct{}

func NewScreen() *Screen { return new(Screen) }
func (s *Screen) setup() {
	fmt.Println("安装屏幕")
}

type Keyboard struct{}

func NewKeyboard() *Keyboard { return new(Keyboard) }
func (k *Keyboard) setup() {
	fmt.Println("安装键盘")
}

type Speaker struct{}

func NewSpeaker() *Speaker { return new(Speaker) }
func (s *Speaker) setup() {
	fmt.Println("安装扩音器")
}

// 整个手机的门面
type PhoneFacade struct {
	screen   *Screen
	keyboard *Keyboard
	speaker  *Speaker
}

func NewPhoneFacade() *PhoneFacade {
	return &PhoneFacade{
		NewScreen(),
		NewKeyboard(),
		NewSpeaker(),
	}
}

// 门面暴露一个方法，在里面设置所有的内部部件。外部的人不需要知道里面是怎么操作的，只调用一下该方法就可以
func (pf *PhoneFacade) setup() {
	pf.screen.setup()
	pf.keyboard.setup()
	pf.speaker.setup()
}
