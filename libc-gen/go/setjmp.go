package libc

func _throw_longjmp() { panic("emscripten_longjmp") }

func _catch_longjmp(sp sptr_t) {
	if r := recover(); r == "emscripten_longjmp" {
		*m.X__stack_pointer() = sp
	} else if r != nil {
		panic(r)
	}
}
