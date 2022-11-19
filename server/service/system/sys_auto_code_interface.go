package system

import (
	"github.com/edufriendchen/hertz-vue-admin/server/global"
	"github.com/edufriendchen/hertz-vue-admin/server/model/system/response"
)

type Database interface {
	GetDB(businessDB string) (data []response.Db, err error)
	GetTables(businessDB string, dbName string) (data []response.Table, err error)
	GetColumn(businessDB string, tableName string, dbName string) (data []response.Column, err error)
}

func (autoCodeService *AutoCodeService) Database(businessDB string) Database {

	if businessDB == "" {
		switch global.CONFIG.System.DbType {
		case "mysql":
			return AutoCodeMysql
		case "pgsql":
			return AutoCodePgsql
		default:
			return AutoCodeMysql
		}
	} else {
		for _, info := range global.CONFIG.DBList {
			if info.AliasName == businessDB {

				switch info.Type {
				case "mysql":
					return AutoCodeMysql
				case "pgsql":
					return AutoCodePgsql
				case "oracle":
					return AutoCodeOracle
				default:
					return AutoCodeMysql
				}
			}
		}
		return AutoCodeMysql
	}

}
