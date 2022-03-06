package pdf

import (
	"html/template"

	"github.com/MalukiMuthusi/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
)

func GeneratePDF() (*string, error) {
	t, err := template.New("statements").Parse(TPL)
	if err != nil {
		logger.Info("GeneratePDF", zap.Error(err))
		return nil, err
	}

	var b buffer.Buffer
	err = t.Execute(&b, "Maluki Muthusi")
	return nil, nil
}
