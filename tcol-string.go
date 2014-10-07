// color project main.go
package tcol256

func FgBgString(str string, fg, bg int) string {
	return get_print_color(str, fg, bg)
}

func FgString(str string, fg int) string {
	return get_print_color(str, fg, NONE)
}

func BgString(str string, bg int) string {
	return get_print_color(str, NONE, bg)
}

func FgBg65536String(str string, fgr, fgg, fgb int, fbr, fbg, fbb int) string {
	return get_print_color(str, rgb65536(fgr, fgg, fgb), rgb256(fbr, fbg, fbb))
}

func Fg65536String(str string, fgr, fgg, fgb int) string {
	return get_print_color(str, rgb65536(fgr, fgg, fgb), NONE)
}

func Bg65536String(str string, fbr, fbg, fbb int) string {
	return get_print_color(str, NONE, rgb65536(fbr, fbg, fbb))
}

func FgBg256String(str string, fgr, fgg, fgb int, fbr, fbg, fbb int) string {
	return get_print_color(str, rgb256(fgr, fgg, fgb), rgb256(fbr, fbg, fbb))
}

func Fg256String(str string, fgr, fgg, fgb int) string {
	return get_print_color(str, rgb256(fgr, fgg, fgb), NONE)
}

func Bg256String(str string, fbr, fbg, fbb int) string {
	return get_print_color(str, NONE, rgb256(fbr, fbg, fbb))
}

func FgBg6String(str string, fgr, fgg, fgb int, fbr, fbg, fbb int) string {
	return get_print_color(str, rgb(fgr, fgg, fgb), rgb(fbr, fbg, fbb))
}

func Fg6String(str string, fgr, fgg, fgb int) string {
	return get_print_color(str, rgb(fgr, fgg, fgb), NONE)
}

func Bg6String(str string, fbr, fbg, fbb int) string {
	return get_print_color(str, NONE, rgb(fbr, fbg, fbb))
}

func get_set_color(fg, bg int) string {
	//Print escape codes to set the terminal color.
	//fg and bg are indices into the color palette for the foreground and
	//background colors.
	return _set_color(fg, bg)
}
func get_reset_color() string {
	//Reset terminal color to default.
	return _reset_color()
}

func get_print_color(str string, fg, bg int) string {
	return get_set_color(fg, bg) + str + get_reset_color()
}
