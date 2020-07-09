package runtime

import (
	"errors"
	"io"
	"os"
	"path"
	"strings"
)

type fileWriter struct {
	width    float32
	height   float32
	fileType string
	filePath string
	evalCh   <-chan evalCtx
}

func (w fileWriter) Start() error {
	for ctx := range w.evalCh {
		var (
			ft string = w.fileType
		)
		fp, err := os.Create(w.filePath)
		if err != nil {
			return err
		}
		if ft == "" {
			// use the extension to guess the file type
			ft = strings.Replace(path.Ext(w.filePath), ".", "", 1)
			if ft == "" {
				return errors.New("could not determine file type")
			}
		}
		reader, err := toReader(ctx, w.width, w.height, ft)
		if err != nil {
			return err
		}
		io.Copy(fp, reader)
	}
	return nil
}
