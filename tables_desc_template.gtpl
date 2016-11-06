{{range .}}
## {{.TableName}} 表
{{with .Columns}}
|  字段  |   说明   |   类型   |  长度  |  是否可空  |  是否索引  |   备注   |
|:-----:|:-------:|:--------:|:------:|:--------:|:--------:|:--------:|
{{range .}}| {{.Name}} | {{.Comment}} | {{.DataType}} | {{if .CharacterMaximumLength.Valid}} {{.CharacterMaximumLength.Int64}} {{else}} - {{end}} | {{.IsNullable}} | {{.IsIndexed}} |  |
{{end}}
{{end}}
{{end}}