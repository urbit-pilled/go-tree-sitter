package hcl

//#include "parser.h"
//TSLanguage *tree_sitter_hcl();
import "C"
import (
	"unsafe"

	sitter "github.com/urbit-pilled/go-tree-sitter"
)

func GetLanguage() *sitter.Language {
	ptr := unsafe.Pointer(C.tree_sitter_hcl())
	return sitter.NewLanguage(ptr)
}
