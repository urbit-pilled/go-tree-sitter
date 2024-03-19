package groovy 

//#include "parser.h"
//TSLanguage *tree_sitter_groovy();
import "C"
import (
	"unsafe"

	sitter "github.com/urbit-pilled/go-tree-sitter"
)

func GetLanguage() *sitter.Language {
	ptr := unsafe.Pointer(C.tree_sitter_groovy())
	return sitter.NewLanguage(ptr)
}
