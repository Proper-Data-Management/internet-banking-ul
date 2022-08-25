package squirrel

func SubQuery(sb SelectBuilder) Sqlizer {
	sql, params, _ := sb.ToSql()
	return Expr("("+sql+")", params...)
}
