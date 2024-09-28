package recorddm

type Record struct {
	Fields map[string]string
}

func NewRecord(fields map[string]string) *Record {
	return &Record{
		Fields: fields,
	}
}

func (r *Record) GetField(fieldName string) (string, bool) {
	value, exists := r.Fields[fieldName]
	return value, exists
}
