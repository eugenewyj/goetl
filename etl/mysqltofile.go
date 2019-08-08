package etl

type MysqlToFile struct {
	mysql    string // mysql命令路径
	user     string // 用户名
	password string // 密码
	host     string // 主机
	port     int    // 端口
	database string // 数据库
	encoding string // 编码
	sql      string // 查询SQL
	file     string // 本地路径
}

// execute 执行将Mysql数据导出到本地文件
func (*MysqlToFile) Execute() {

}

// assemblyCommand 组装命令行内容
func (*MysqlToFile) assemblyCommand() (string) {
	return ""
}
