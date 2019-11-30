package util

import (
	"fmt"
	"github.com/golang/protobuf/ptypes"
	timeReturn "github.com/golang/protobuf/ptypes/timestamp"
	"time"
)

/** DataBRtoProtoBuffDate recebe a data em string no formato brasileiro e converte para Timestamp para o protobuff
* formato da date "DIA-MES-ANO" ---
* Exemple: "20-12-2019"
* Return (*timeReturn.Timestamp, string)
 */
func DataBRtoProtoBuffDate(data string) (*timeReturn.Timestamp, error) {
	date, err := time.Parse("02-01-2006", data)
	if err != nil {
		return nil, fmt.Errorf("Erro na conversão de Data para time :  %s", err)
	}
	t := date.In(time.UTC)
	timestamp, err := ptypes.TimestampProto(t)
	if err != nil {
		return nil, fmt.Errorf("Erro na conversão de Data para Timestamp :  %s", err)
	}
	return timestamp, nil
}
