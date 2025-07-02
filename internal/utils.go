package internal
import (
	"fmt"
	"path/filepath"
	"github.com/Ashank007/docai/reader"
)

func getReaderByExtension(path string) (reader.Reader, error) {
	switch filepath.Ext(path) {
	case ".pdf":
		return reader.NewPDFReader(), nil
	case ".txt":
		return reader.NewTextReader(), nil
	case ".docx":
		return reader.NewDocxReader(), nil
	default:
		return nil, fmt.Errorf("unsupported extension: %s", path)
	}
}
